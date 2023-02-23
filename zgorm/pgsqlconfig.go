package zgorm

import (
	"errors"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type PgSqlConf struct {
	DataSource      string
	MaxIdleConns    int           `json:",default=10"`                               // 空闲中的最大连接数
	MaxOpenConns    int           `json:",default=10"`                               // 打开到数据库的最大连接数
	ConnMaxLifeTime time.Duration `json:",default=1m"`                               //sets the maximum amount of time a connection may be reused
	LogMode         string        `json:",default=dev,options=dev|test|prod|silent"` // 是否开启Gorm全局日志
	LogZap          bool          // 是否通过zap写入日志文件
	SlowThreshold   int64         `json:",default=1000"`
}

func (m *PgSqlConf) GetGormLogMode() logger.LogLevel {
	return overwriteGormLogMode(m.LogMode)
}

func (m *PgSqlConf) GetSlowThreshold() time.Duration {
	return time.Duration(m.SlowThreshold) * time.Millisecond
}
func (m *PgSqlConf) GetColorful() bool {
	return true
}

func ConnectPgSql(m PgSqlConf) (*gorm.DB, error) {
	if m.DataSource == "" {
		return nil, errors.New("database dsn is empty")
	}
	newLogger := newDefaultGormLogger(&m)
	pgsqlCfg := postgres.Config{
		DSN:                  m.DataSource,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}
	db, err := gorm.Open(postgres.New(pgsqlCfg), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, err
	} else {
		sqldb, _ := db.DB()
		sqldb.SetMaxIdleConns(m.MaxIdleConns)
		sqldb.SetMaxOpenConns(m.MaxOpenConns)
		sqldb.SetConnMaxLifetime(m.ConnMaxLifeTime)
		return db, nil
	}
}
