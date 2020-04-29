package option-pattern

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
	"xorm.io/xorm/log"
	"xorm.io/xorm/names"
)

func NewDataBase(
	driver string,
	host string,
	port uint,
	dbname string,
	user string,
	password string,
	options ...Option,
) (*xorm.Engine, error) {
	engine, err := xorm.NewEngine(driver, fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", user, password, host, port, dbname))
	if err != nil {
		return nil, err
	}

	// set value
	for _, option := range options {
		option.Apply(engine)
	}

	return engine, nil
}

// An Option configures a xorm.Engine
type Option interface {
	Apply(*xorm.Engine)
}

// OptionFunc is a function that configures a xorm.Engine
type OptionFunc func(*xorm.Engine)

// Apply is a function that set value to xorm.Engine
func (f OptionFunc) Apply(engine *xorm.Engine) {
	f(engine)
}

func SetConnMaxLifetime(maxlifetime time.Duration) Option {
	return OptionFunc(func(engine *xorm.Engine) {
		engine.SetConnMaxLifetime(maxlifetime)
	})
}

func SetMaxIdleConns(maxIdleConns int) Option {
	return OptionFunc(func(engine *xorm.Engine) {
		engine.SetMaxIdleConns(maxIdleConns)
	})
}

func SetMaxOpenConns(maxOpenConns int) Option {
	return OptionFunc(func(engine *xorm.Engine) {
		engine.SetMaxOpenConns(maxOpenConns)
	})
}

func SetLogLevel(logLevel log.LogLevel) Option {
	return OptionFunc(func(engine *xorm.Engine) {
		engine.SetLogLevel(logLevel)
	})
}

func SetLogger(logger interface{}) Option {
	return OptionFunc(func(engine *xorm.Engine) {
		engine.SetLogger(logger)
	})
}

func SetShowSQL(showSQL ...bool) Option {
	return OptionFunc(func(engine *xorm.Engine) {
		engine.ShowSQL(showSQL...)
	})
}

func SetMapper(mapper names.Mapper) Option {
	return OptionFunc(func(engine *xorm.Engine) {
		engine.SetMapper(mapper)
	})
}
