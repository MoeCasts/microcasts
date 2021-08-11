package config

import (
	"fmt"

	"github.com/moecasts/microcasts/orm"
)

func WithOrmPostgresConfig(config Config) map[string]orm.PostgresConfigOption {
	configs := make([]orm.PostgresDSNConfig, 0)
	err := config.UnmarshalKey("postgres", &configs)
	if err != nil {
		panic(fmt.Errorf("postgres 配置错误 %v", err))
	}

	opts := make(map[string]orm.PostgresConfigOption)

	for _, config := range configs {
		opts[config.Name] = orm.WithDSN(config)
	}
	return opts
}
