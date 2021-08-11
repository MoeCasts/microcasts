package microcasts

import (
	"github.com/moecasts/microcasts/config"
	"github.com/moecasts/microcasts/orm"
	"gorm.io/gorm"
)

func Orm(name ...string) *gorm.DB {
	BootDefaultGorm()

	return orm.DefaultPool.Get(name...)
}

func BootDefaultGorm(opts ...orm.GormConfigOption) {
	if len(orm.DefaultPool) > 0 {
		return
	}

	// opts = append(opts, nil)
	//
	// copy(opts[1:], opts)

	if len(orm.DefaultPool) == 0 {
		for name, opt := range config.WithOrmPostgresConfig(config.DefaultConfig()) {
			db := orm.NewPostgresOrm(opts, []orm.PostgresConfigOption{opt})
			orm.DefaultPool.Add(name, db)
		}
	}
}
