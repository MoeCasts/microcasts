module github.com/moecasts/microcasts

go 1.15

// 替换成本地包
replace github.com/moecasts/microcasts => ./

replace github.com/moecasts/microcasts/novels => ./novels

require (
	github.com/jackc/pgx/v4 v4.13.0 // indirect
	github.com/kr/pretty v0.2.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/spf13/viper v1.8.1
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gorm.io/driver/postgres v1.1.0
	gorm.io/gorm v1.21.12
)
