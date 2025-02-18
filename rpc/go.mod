module github.com/nerdynz/skeleton/rpc

go 1.22.4

toolchain go1.22.5

replace github.com/nerdynz/security => ../../../nerdynz/security

replace github.com/nerdynz/datastore => ../../../nerdynz/datastore

require (
	github.com/go-zoo/bone v1.3.0
	github.com/jinzhu/copier v0.4.0
	github.com/lmittmann/tint v1.0.5
	github.com/nerdynz/datastore v0.0.0-20210404043820-fca6c2b865be
	github.com/nerdynz/fileupload v1.0.1
	github.com/nerdynz/flow v1.0.0
	github.com/nerdynz/rcache v0.0.0-20200404024229-09aee2ea3078
	github.com/nerdynz/router v1.0.0
	github.com/nerdynz/security v0.0.0-20231122090124-02bb87723116
	github.com/nerdynz/trove v0.0.0-20200425063959-61f6ab2f6311
	github.com/r3labs/sse v0.0.0-20210224172625-26fe804710bc
	github.com/sirupsen/logrus v1.8.1
	github.com/twitchtv/twirp v8.1.2+incompatible
	github.com/unrolled/render v1.7.0
	github.com/urfave/negroni v1.0.0
	golang.org/x/exp v0.0.0-20220722155223-a9213eeb770e
	golang.org/x/net v0.14.0
	google.golang.org/protobuf v1.28.0
	rogchap.com/v8go v0.9.0
)

require (
	github.com/FZambia/go-sentinel v0.0.0-20171204085413-76bd05e8e22f // indirect
	github.com/cenkalti/backoff v2.2.1+incompatible // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/garyburd/redigo v1.6.3 // indirect
	github.com/georgysavva/scany/v2 v2.1.3 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.6.0 // indirect
	github.com/jackc/puddle/v2 v2.2.1 // indirect
	github.com/jmoiron/sqlx v1.3.5 // indirect
	github.com/joho/godotenv v1.3.0 // indirect
	github.com/kennygrant/sanitize v1.2.4 // indirect
	github.com/lib/pq v1.10.9 // indirect
	github.com/mgutz/str v1.2.0 // indirect
	github.com/nerdynz/dat v1.4.0 // indirect
	github.com/oklog/ulid v1.3.1 // indirect
	github.com/oklog/ulid/v2 v2.1.0 // indirect
	github.com/pmylund/go-cache v2.1.0+incompatible // indirect
	github.com/shomali11/xredis v0.0.0-20190608143638-0b54a6bbf40b // indirect
	github.com/simukti/sqldb-logger v0.0.0-20230108155151-646c1a075551 // indirect
	golang.org/x/crypto v0.17.0 // indirect
	golang.org/x/sync v0.1.0 // indirect
	golang.org/x/sys v0.15.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	gopkg.in/cenkalti/backoff.v1 v1.1.0 // indirect
)
