module github.com/shiningw/yadownloader

go 1.18

//replace github.com/shiningw/aria2go => ../aria2go/

require (
	github.com/gorilla/mux v1.8.0
	github.com/mattn/go-sqlite3 v1.14.14
	github.com/shiningw/aria2go v0.0.0-20220720110802-4833b353857f
)

require github.com/filebrowser/filebrowser/v2 v2.22.4

require (
	github.com/GeertJohan/go.rice v1.0.2 // indirect
	github.com/asdine/storm v2.1.2+incompatible // indirect
	github.com/daaku/go.zipexe v1.0.1 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/dsnet/compress v0.0.1 // indirect
	github.com/flynn/go-shlex v0.0.0-20150515145356-3f9db97f8568 // indirect
	github.com/fsnotify/fsnotify v1.5.4 // indirect
	github.com/golang-jwt/jwt/v4 v4.4.2 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/hacdias/fileutils v1.0.0 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/kardianos/osext v0.0.0-20190222173326-2bc1f35cddc0 // indirect
	github.com/magiconair/properties v1.8.6 // indirect
	github.com/maruel/natural v1.1.0 // indirect
	github.com/mholt/archiver v3.1.1+incompatible // indirect
	github.com/mholt/caddy v1.0.0 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/nwaples/rardecode v1.1.3 // indirect
	github.com/pelletier/go-toml v1.9.5 // indirect
	github.com/pelletier/go-toml/v2 v2.0.2 // indirect
	github.com/pierrec/lz4 v2.6.1+incompatible // indirect
	github.com/spf13/afero v1.9.2 // indirect
	github.com/spf13/cast v1.5.0 // indirect
	github.com/spf13/cobra v1.5.0 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.12.0 // indirect
	github.com/subosito/gotenv v1.4.0 // indirect
	github.com/ulikunitz/xz v0.5.10 // indirect
	github.com/xi2/xz v0.0.0-20171230120015-48954b6210f8 // indirect
	go.etcd.io/bbolt v1.3.6 // indirect
	golang.org/x/crypto v0.0.0-20220622213112-05595931fe9d // indirect
	golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect
	golang.org/x/text v0.3.7 // indirect
	gopkg.in/ini.v1 v1.66.6 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/filebrowser/filebrowser/v2 => github.com/shiningw/filebrowser/v2 v2.0.0
