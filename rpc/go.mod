module github.com/nerdynz/skeleton/rpc

go 1.18

replace github.com/nerdynz/security => ../../../nerdynz/security

replace github.com/nerdynz/datastore => ../../../nerdynz/datastore

require (
	github.com/nerdynz/datastore v0.0.0-00010101000000-000000000000
	github.com/nerdynz/rcache v0.0.0-20200404024229-09aee2ea3078
	github.com/nerdynz/security v0.0.0-20220524061829-ddebeb4081bc
	github.com/nerdynz/trove v0.0.0-20200425063959-61f6ab2f6311
	github.com/sirupsen/logrus v1.8.1
	github.com/twitchtv/twirp v8.1.2+incompatible
	github.com/urfave/negroni v1.0.0
	golang.org/x/exp v0.0.0-20220722155223-a9213eeb770e
	google.golang.org/protobuf v1.28.0
)

require (
	github.com/FZambia/go-sentinel v0.0.0-20171204085413-76bd05e8e22f // indirect
	github.com/cenkalti/backoff v2.2.1+incompatible // indirect
	github.com/garyburd/redigo v1.6.0 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/jmoiron/sqlx v1.2.0 // indirect
	github.com/joho/godotenv v1.3.0 // indirect
	github.com/lib/pq v1.4.0 // indirect
	github.com/mgutz/str v1.2.0 // indirect
	github.com/nats-io/nats.go v1.12.3 // indirect
	github.com/nats-io/nkeys v0.3.0 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/nerdynz/dat v1.2.0 // indirect
	github.com/oklog/ulid v1.3.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmylund/go-cache v2.1.0+incompatible // indirect
	github.com/satori/go.uuid v1.2.0 // indirect
	github.com/shomali11/xredis v0.0.0-20190608143638-0b54a6bbf40b // indirect
	golang.org/x/crypto v0.0.0-20210616213533-5ff15b29337e // indirect
	golang.org/x/sys v0.0.0-20211019181941-9d821ace8654 // indirect
	gopkg.in/mattes/migrate.v1 v1.3.2 // indirect
	gopkg.in/olahol/melody.v1 v1.0.0-20170518105555-d52139073376 // indirect
)
