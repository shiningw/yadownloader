module github.com/shiningw/yadownloader

go 1.18

//replace github.com/shiningw/aria2go => ../aria2go/

require (
	github.com/gorilla/mux v1.8.0
	github.com/mattn/go-sqlite3 v1.14.14
	github.com/shiningw/aria2go v0.0.0-20220720110802-4833b353857f
)

require github.com/filebrowser/filebrowser/v2 v2.0.0

require (
	github.com/golang-jwt/jwt/v4 v4.4.2 // indirect
	github.com/maruel/natural v1.0.0 // indirect
	github.com/spf13/afero v1.8.2 // indirect
	golang.org/x/crypto v0.0.0-20220427172511-eb4f295cb31f // indirect
	golang.org/x/text v0.3.7 // indirect
)

//replace github.com/filebrowser/filebrowser/v2 => ../fb
