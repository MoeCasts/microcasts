module github.com/moecasts/microcasts/novels

go 1.15

replace github.com/moecasts/microcasts => ../

replace github.com/moecasts/microcasts/novels => ./

require (
	github.com/go-kit/kit v0.11.0
	github.com/golang/protobuf v1.5.2
	github.com/google/uuid v1.3.0 // indirect
	github.com/infobloxopen/atlas-app-toolkit v1.0.0
	github.com/infobloxopen/protoc-gen-gorm v0.21.0
	github.com/jackc/pgx/v4 v4.13.0 // indirect
	github.com/jinzhu/gorm v1.9.16
	github.com/lightstep/lightstep-tracer-go v0.25.0
	github.com/moecasts/microcasts v0.0.0-00010101000000-000000000000
	github.com/oklog/oklog v0.3.2
	github.com/oklog/run v1.1.0 // indirect
	github.com/opentracing/basictracer-go v1.1.0 // indirect
	github.com/opentracing/opentracing-go v1.2.0
	github.com/openzipkin-contrib/zipkin-go-opentracing v0.4.5
	github.com/openzipkin/zipkin-go v0.2.5
	github.com/prometheus/client_golang v1.11.0
	github.com/satori/go.uuid v1.2.0
	github.com/stretchr/testify v1.7.0
	golang.org/x/net v0.0.0-20210726213435-c6fcb2dbf985
	google.golang.org/genproto v0.0.0-20210617175327-b9e0b3197ced
	google.golang.org/grpc v1.39.0
	google.golang.org/protobuf v1.27.1
	gorm.io/driver/postgres v1.1.0 // indirect
	gorm.io/gorm v1.21.12
	sourcegraph.com/sourcegraph/appdash v0.0.0-20190731080439-ebfcffb1b5c0
)
