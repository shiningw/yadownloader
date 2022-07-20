package db

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

var (
	mainTableName = "downloads"
)

type DbTable[T IRow] struct {
	table string
	row   T
	DB    *sql.DB
}

func NewDbTable[T IRow](db *sql.DB, table string, row T) DbTable[T] {
	return DbTable[T]{table, row, db}
}
func (d *DbTable[T]) Query(query string) (*sql.Rows, error) {
	return d.DB.Query(query)
}

func (d *DbTable[T]) Exec(query string) (sql.Result, error) {
	return d.DB.Exec(query)
}
func (d *DbTable[T]) ExecuteQuery(query string, args []any) (sql.Result, error) {
	stmt, err := d.DB.Prepare(query)
	if err == nil {
		defer stmt.Close()
		res, err := stmt.Exec(args...)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
	return nil, err
}
func (d *DbTable[T]) Insert() (sql.Result, error) {
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES(%s)", d.table, strings.Join(d.row.FieldNames(), ","), strings.Repeat("?,", d.row.NumField()-1)+"?")
	vals := make([]any, d.row.NumField())
	for i := 0; i < d.row.NumField(); i++ {
		vals[i] = reflect.ValueOf(d.row).Field(i).Interface()
	}
	return d.ExecuteQuery(query, vals)
}

func (d *DbTable[T]) InsertTx(tx *sql.Tx) (sql.Result, error) {
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES(%s)", d.table, strings.Join(d.row.FieldNames(), ","), strings.Repeat("?,", d.row.NumField()-1)+"?")
	vals := make([]any, d.row.NumField())
	for i := 0; i < d.row.NumField(); i++ {
		vals[i] = reflect.ValueOf(d.row).Field(i).Interface()
	}
	res, err := tx.Exec(query, vals...)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (d *DbTable[T]) UpdateQuery(query string, args []any) (sql.Result, error) {
	return d.ExecuteQuery(query, args)
}

func (d *DbTable[T]) DeleteQuery(query string, args []any) (sql.Result, error) {
	return d.ExecuteQuery(query, args)
}

func (d *DbTable[T]) Select(query string) []T {
	results := make([]T, 0)
	st := reflect.ValueOf(&d.row).Elem()
	if st.Kind() != reflect.Struct {
		return nil
	}
	//columns := row.FieldNames()
	//query := fmt.Sprintf("SELECT %s FROM %s", strings.Join(columns, ","), table)
	rows, err := d.Query(query)
	if err != nil {
		log.Println(err)
		return nil
	}
	dest := GetDest(st)
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(dest...)
		if err != nil {
			log.Println(err)
			continue
		}
		if len(dest) != st.NumField() {
			log.Println("dest != st.NumField()")
			continue
		}

		for i := 0; i < st.NumField(); i++ {
			if st.Field(i).CanSet() {
				switch v := dest[i].(type) {
				case *int64:
					st.Field(i).SetInt(*v)
				case *string:
					st.Field(i).SetString(*v)

				}
			}
		}
		results = append(results, d.row)
	}
	return results
}

func (d *DbTable[T]) SelectAll() []T {

	query := fmt.Sprintf("SELECT %s from %s", strings.Join(d.row.FieldNames(), ","), d.table)
	return d.Select(query)
}

func (d *DbTable[T]) Count(query string) int {
	var rows *sql.Rows
	var err error
	if query == "" {
		rows, err = d.Query("SELECT count(*) from " + d.table)
	} else {
		rows, err = d.Query(query)
	}
	if err != nil {
		return 0
	}
	defer rows.Close()
	if err != nil {
		return 0
	}
	c := 0
	if rows.Next() {
		e := rows.Scan(&c)
		if e != nil {
			log.Println(e)
		}
	}
	return c
}

func UseTable[T IRow](table string, row T) DbTable[T] {
	d := getDatabase()
	return NewDbTable(d.DB(), table, row)
}

func GetDest(v reflect.Value) []interface{} {
	if v.Kind() != reflect.Struct {
		return nil
	}
	vals := make([]interface{}, 0)
	for i := 0; i < v.NumField(); i++ {
		vv := v.Field(i)
		switch vv.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			vals = append(vals, new(int64))
		case reflect.String:
			vals = append(vals, new(string))
		case reflect.Bool:
			vals = append(vals, new(bool))
		case reflect.Float32, reflect.Float64:
			vals = append(vals, new(float64))
		}
	}
	return vals
}

func GetFieldValues(r IRow) []string {

	names := r.FieldNames()
	var values []string
	for _, name := range names {
		value := FieldValue(r, name)
		if value.Kind() == reflect.String {
			values = append(values, "'"+value.String()+"'")
		} else {
			values = append(values, fmt.Sprintf("%d", value.Int()))
		}
	}
	return values
}
func StructNames(s interface{}) []string {
	names := []string{}
	t := reflect.TypeOf(s)
	if t.Kind() == reflect.Struct {
		for i := 0; i < t.NumField(); i++ {
			names = append(names, t.Field(i).Name)
		}
	}
	return names
}

func FieldValue(st any, name string) reflect.Value {
	v := reflect.ValueOf(st)
	value := reflect.Indirect(v)
	if value.Kind() != reflect.Struct {
		panic("not struct")
	}
	field := value.FieldByName(name)
	if !field.IsValid() {
		panic("not found:" + name)
	}
	return field
}

func saveDownloads(row IRow, table string, tx *sql.Tx) (sql.Result, error) {
	t := UseTable(table, row)
	res, e := t.InsertTx(tx)
	return res, e
}
func Update(sql string, args []any) {
	db := UseTable(mainTableName, YtdRow{})
	db.UpdateQuery(sql, args)
}

func Save(row DownloadsRow, typ string) error {
	dbHelper := getDatabase()
	dbHelper.Begin()
	_, e := saveDownloads(row, mainTableName, dbHelper.TX())
	if e != nil {
		dbHelper.TX().Rollback()
	} else {
		var extra IRow
		if strings.ToLower(typ) == "aria2" {
			extra = Aria2Row{row.Gid, "{}"}
		} else {
			extra = YtdRow{Gid: row.Gid}
		}
		if _, e = saveDownloads(extra, strings.ToLower(typ), dbHelper.TX()); e == nil {
			dbHelper.TX().Commit()
		} else {
			log.Println(e)
			dbHelper.TX().Rollback()
		}
	}
	return e
}

func GetCount(typ DownloadType) int {
	db := UseTable(mainTableName, DownloadsRow{})
	query := fmt.Sprintf("select count(*) from %s where type = %d", mainTableName, typ)
	return db.Count(query)
}

func GetYTDDownloads() []YtdQueue {
	row := YtdQueue{}
	db := UseTable(mainTableName, row)
	cols := row.FieldNames()
	cols[0] = "y.gid"
	fields := strings.Join(cols, ",")
	query := fmt.Sprintf("select %s from %s as dl inner join ytd as y on dl.gid = y.gid where type = %d", fields, mainTableName, YTD)
	return db.Select(query)
}

//delete entry from downloads and ytd
func DeleteByGid(gid string) (sql.Result, error) {
	query := fmt.Sprintf("delete from %s where gid = ?", mainTableName)
	args := []any{gid}
	db := UseTable(mainTableName, DownloadsRow{})
	return db.DeleteQuery(query, args)
}
