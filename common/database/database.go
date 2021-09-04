package database

import (
	"apus-sample/common/constant"
	"apus-sample/common/utils"
	"apus-sample/internal/appconf"
	"context"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"sync"
)

type Database struct {
	dialector gorm.Dialector
	schemas   sync.Map
}

func New() (*Database, error) {
	conf := appconf.Database()
	s := &Database{}
	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d",
		conf.Host, conf.Username, conf.Password, conf.DBName, conf.Port)
	s.dialector = postgres.Open(dns)
	_, err := s.connectAndStore(constant.DefaultSchema)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (s *Database) MustGet(ctx context.Context, schemaName string) *gorm.DB {
	conn, err := s.Get(ctx, schemaName)
	utils.PanicWhenError(err)
	return conn
}

func (s *Database) Get(ctx context.Context, schemaName string) (conn *gorm.DB, err error) {
	schemaConn, found := s.schemas.Load(schemaName)
	if !found {
		dbConn, err := s.connectAndStore(schemaName)
		if err != nil {
			return nil, err
		}
		return dbConn.WithContext(ctx), nil
	} else {
		return schemaConn.(*gorm.DB).WithContext(ctx), nil
	}
}

func (s *Database) connectAndStore(schemaName string) (*gorm.DB, error) {
	var ormConf = &gorm.Config{NamingStrategy: schema.NamingStrategy{
		TablePrefix:   schemaName + ".",
		SingularTable: true,
	}}
	if utils.GetApplicationEnv() == constant.EnvDev {
		ormConf.Logger = ormLogger
	}
	dbConn, err := gorm.Open(s.dialector, ormConf)
	if err != nil {
		return nil, err
	}
	s.schemas.Store(schemaName, dbConn)
	return dbConn, err
}
