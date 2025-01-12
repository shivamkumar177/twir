module github.com/satont/twir/apps/watched

go 1.20

replace github.com/satont/twir/libs/config => ../../libs/config

replace github.com/satont/twir/libs/gomodels => ../../libs/gomodels

replace github.com/satont/twir/libs/twitch => ../../libs/twitch

require (
	github.com/gofrs/uuid v4.0.0+incompatible
	github.com/nicklaw5/helix/v2 v2.22.2
	github.com/satont/twir/libs/config v0.0.0-20221125194658-5cb70dbdbf2a
	github.com/satont/twir/libs/gomodels v0.0.0-20221125194658-5cb70dbdbf2a
	github.com/satont/twir/libs/grpc v0.0.0-00010101000000-000000000000
	github.com/satont/twir/libs/twitch v0.0.0-20221125194658-5cb70dbdbf2a
	go.uber.org/zap v1.24.0
	google.golang.org/grpc v1.55.0
	google.golang.org/protobuf v1.30.0
	gorm.io/driver/postgres v1.5.2
	gorm.io/gorm v1.25.1
)

require (
	github.com/benbjohnson/clock v1.3.0 // indirect
	github.com/golang-jwt/jwt/v4 v4.5.0 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/guregu/null v4.0.0+incompatible // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.4.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/kelseyhightower/envconfig v1.4.0 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/lib/pq v1.10.9 // indirect
	github.com/samber/lo v1.38.1 // indirect
	github.com/satori/go.uuid v1.2.0 // indirect
	go.uber.org/atomic v1.11.0 // indirect
	go.uber.org/goleak v1.1.12 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/crypto v0.10.0 // indirect
	golang.org/x/exp v0.0.0-20230713183714-613f0c0eb8a1 // indirect
	golang.org/x/net v0.11.0 // indirect
	golang.org/x/sys v0.10.0 // indirect
	golang.org/x/text v0.11.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230530153820-e85fd2cbaebc // indirect
)

replace github.com/satont/twir/libs/grpc => ../../libs/grpc
