package db

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/shiningw/yadownloader/config"
)

type Dbhelper struct {
	dsn string
	dbh *sql.DB
	tx  *sql.Tx
}

var (
	dbHelper *Dbhelper
)
var once sync.Once

func getDatabase() *Dbhelper {
	var dsn string = config.GetConfig().Dsn
	once.Do(func() {
		dbHelper = NewDbhelper(dsn, "sqlite3")
	})
	return dbHelper
}

func NewDbhelper(dsn, driver string) *Dbhelper {
	var err error
	d := &Dbhelper{}
	//simplify constructing dsn
	if driver == "sqlite3" {
		if strings.HasPrefix(dsn, "file:") {
			d.dsn = dsn
		} else {
			d.dsn = fmt.Sprintf("file:%s?_foreign_keys=true", dsn)
		}
	} else {
		//d.dsn = "root:@tcp(localhost:3306)/test?charset=utf8"
		//d.dsn = "root:password@unix(/var/run/mysqld/mysqld.sock)/test?charset=utf8"
		d.dsn = dsn
	}
	//d.dsn = fmt.Sprintf("file:%s?_foreign_keys=true", file)
	d.dbh, err = sql.Open(driver, d.dsn)
	if err != nil {
		log.Panic(err)
	}
	d.Create(nil)
	return d
}

func (d *Dbhelper) Close() {
	d.dbh.Close()
}

func (d *Dbhelper) DB() *sql.DB {
	return d.dbh
}

func (d *Dbhelper) TX() *sql.Tx {
	return d.tx
}

func (d *Dbhelper) Begin() *sql.Tx {
	tx, e := d.dbh.Begin()
	if e != nil {
		log.Fatal(e)
	}
	d.tx = tx
	return tx
}

func (d *Dbhelper) Rollback() {
	d.tx.Rollback()
}

func (d *Dbhelper) Commit() {
	d.tx.Commit()
}

func (d *Dbhelper) Create(sqlStmt []string) {
	if sqlStmt == nil {
		sqlStmt = defaultSchema()
	}
	for _, table := range sqlStmt {
		_, e := d.dbh.Exec(table)
		if e != nil {
			log.Fatal(e)
		}
	}
}

func defaultSchema() []string {
	sqlStmt := []string{`CREATE TABLE if not exists downloads(
		id integer PRIMARY KEY AUTOINCREMENT,
		filename text default 'unknown',
		uid text not NULL,
		gid text NOT NULL,
		url text NOT NULL,
		type integer default 1,
		timestamp integer default 0,
		status integer default 1,
		UNIQUE(gid)
	  );`, `CREATE TABLE if not exists aria2 (
		id integer PRIMARY KEY AUTOINCREMENT,
		gid integer references downloads(gid) on delete cascade,
		data text default '{}');
		`, `CREATE TABLE if not exists ytd (
			id integer PRIMARY KEY AUTOINCREMENT,
			gid integer references downloads(gid) on delete cascade,
			speed text default '0',
			progress text default '0%',
			filesize integer default '0',
		    data text default '{}');
		`}
	return sqlStmt
}
