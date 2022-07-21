module github.com/shiningw/yadownloader

go 1.18

require (
	github.com/gorilla/mux v1.8.0
	github.com/mattn/go-sqlite3 v1.14.14
)

require (
	github.com/filebrowser/filebrowser/v2 v2.22.4
	github.com/golang-jwt/jwt/v4 v4.4.2
)

require (
	github.com/maruel/natural v1.1.0 // indirect
	github.com/spf13/afero v1.9.2 // indirect
	golang.org/x/crypto v0.0.0-20220622213112-05595931fe9d // indirect
	golang.org/x/text v0.3.7 // indirect
)

replace github.com/filebrowser/filebrowser/v2 => github.com/shiningw/filebrowser/v2 v2.0.0
