package zgorm

import (
	"errors"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type MysqlConf struct {
	DataSource      string
	MaxIdleConns    int           `json:",default=64"`                               // 空闲中的最大连接数
	MaxOpenConns    int           `json:",default=64"`                               // 打开到数据库的最大连接数
	ConnMaxLifeTime time.Duration `json:",default=1m"`                               // sets the maximum amount of time a connection may be reused
	LogMode         string        `json:",default=dev,options=dev|test|prod|silent"` // 是否开启Gorm全局日志
	LogZap          bool          // 是否通过zap写入日志文件
	SlowThreshold   int64         `json:",default=1000"`
}

func (m *MysqlConf) GetGormLogMode() logger.LogLevel {
	return overwriteGormLogMode(m.LogMode)
}

func (m *MysqlConf) GetSlowThreshold() time.Duration {
	return time.Duration(m.SlowThreshold) * time.Millisecond
}
func (m *MysqlConf) GetColorful() bool {
	return true
}

func ConnectMysql(m MysqlConf) (*gorm.DB, error) {
	if m.DataSource == "" {
		return nil, errors.New("database dsn is empty")
	}
	mysqlCfg := mysql.Config{
		DSN: m.DataSource,
	}
	newLogger := newDefaultGormLogger(&m)
	db, err := gorm.Open(mysql.New(mysqlCfg), &gorm.Config{
		//Logger: logger.Default.LogMode(logger.Info),
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
