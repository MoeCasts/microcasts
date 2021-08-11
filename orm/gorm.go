package orm

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type PostgresConfigOption func(*postgres.Config)

type GormConfigOption func(*gorm.Config)

func NewPostgresOrm(gormOpts []GormConfigOption, postgresOpts []PostgresConfigOption) *gorm.DB {
	postgresConfig := &postgres.Config{}

	for _, postgresOpt := range postgresOpts {
		postgresOpt(postgresConfig)
	}

	gormConfig := &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	}

	for _, gormOpt := range gormOpts {
		gormOpt(gormConfig)
	}

	db, err := gorm.Open(postgres.New(*postgresConfig), gormConfig)

	if err != nil {
		panic(fmt.Errorf("db 实例化失败 %v", err))
	}

	return db
}

func WithTablePrefix(prefix string) GormConfigOption {
	return func(config *gorm.Config) {
		config.NamingStrategy = schema.NamingStrategy{
			TablePrefix: prefix,
		}
	}
}

type Pool map[string]*gorm.DB

const DefaultDBName = "default"

var DefaultPool Pool = make(map[string]*gorm.DB)

func (p Pool) Get(name ...string) *gorm.DB {
	if p == nil {
		return nil
	}

	return p[DefaultDBName]
}

func (p Pool) Add(name string, db *gorm.DB) {
	if p == nil {
		p = make(Pool)
	}

	if len(name) == 0 {
		name = DefaultDBName
	}

	p[name] = db
}

func WithDSN(config PostgresDSNConfig) PostgresConfigOption {
	return func(c *postgres.Config) {
		c.DSN = fmt.Sprintf(
			"host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=%v",
			config.Host,
			config.User,
			config.Password,
			config.Database,
			config.Port,
			config.SslMode,
			config.TimeZone,
		)
	}
}

type PostgresDSNConfig struct {
	Name     string
	User     string
	Password string
	Host     string
	Port     int
	Database string
	SslMode  string
	TimeZone string
}
