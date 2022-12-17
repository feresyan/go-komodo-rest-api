package db_connection

import (
	"fmt"
	"gokomodo-assesment/app/domain/entity"
	"gokomodo-assesment/app/helper/faker"
	"gokomodo-assesment/internal/config"
	"gokomodo-assesment/internal/config/env"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

var (
	DbConnection *gorm.DB
)

var table = []interface{}{
	&entity.Buyer{},
	&entity.Seller{},
	&entity.Product{},
	&entity.Order{},
}

func init() {
	var err error
	cfg := config.GetConfig()
	DbConnection, err = Conn(cfg.EnvSystem.GokomodoDB)
	if err != nil {
		panic("Cannot Connect to Database")
	}

	// Migrate all entity
	MigrateSchema(DbConnection, table)

	// Faker generator
	faker.GenerateBuyerFaker(DbConnection)
	faker.GenerateSellerFaker(DbConnection)
}

func Conn(cfg env.Database) (*gorm.DB, error) {
	fmt.Printf("Connect to Database %v\n", cfg.Dbname)
	user := cfg.Username
	password := cfg.Password
	host := cfg.Host
	port := cfg.Port
	dbname := cfg.Dbname

	dsn := user + ":" + password + "@(" + host + ":" + port + ")/" + dbname + "?charset=utf8&parseTime=True&loc=Local"
	// dbConn, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	dbConn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, err
	}

	conn, err := dbConn.DB()

	if err != nil {
		return nil, err
	}

	conn.SetMaxIdleConns(7)
	conn.SetMaxOpenConns(10)
	conn.SetConnMaxLifetime(1 * time.Hour)

	if err != nil {
		return nil, err
	}

	return dbConn, nil
}

func MigrateSchema(db *gorm.DB, tables []interface{}) {
	err := db.AutoMigrate(tables...)
	if err != nil {
		log.Fatalln("Error While Migrateing Schema", err.Error())
	}
}
