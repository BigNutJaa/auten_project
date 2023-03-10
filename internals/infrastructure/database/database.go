package database

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/BigNutJaa/users/internals/config"
	"github.com/BigNutJaa/users/internals/entity"

	"gorm.io/gorm/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB is struct of this file.
type DB struct {
	Connection *gorm.DB
	sql        *sql.DB
	env        config.Configuration
}

func (db *DB) IsErrorRecordNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}

// Close Connection DB
func (db *DB) Close() {
	if err := db.sql.Close(); err != nil {
		log.Errorf("Error closing db connection %s", err)
	} else {
		log.Info("DB connection closed")
	}
}

func (db *DB) MigrateDB() {
	log.Info("Start migrate db READ")

	if !db.Connection.Migrator().HasTable(entity.Users{}.TableName()) {
		err := db.Connection.AutoMigrate(&entity.Users{})

		log.Println("Error :", err)
	}
	if !db.Connection.Migrator().HasTable(entity.Token{}.TableName()) {
		err := db.Connection.AutoMigrate(&entity.Token{})

		log.Println("Error :", err)
	}
	if !db.Connection.Migrator().HasTable(entity.Products{}.TableName()) {
		err := db.Connection.AutoMigrate(&entity.Products{})

		log.Println("Error :", err)
	}
}

// NewServerBase is start connection database.
func NewServerBase(env config.Configuration) *DB {
	log.Info("start New serverBase")

	//dsn := fmt.Sprintf(
	//	os.Getenv("MYSQL_DNS"),
	//	env.DbHost,
	//	env.DbPort,
	//	env.DbUser,
	//	env.DbName,
	//	env.DbPassword,
	//)

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		env.DbHost,
		env.DbPort,
		env.DbUser,
		env.DbName,
		env.DbPassword,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Panic(err)
	}

	if env.Env != "production" {
		db.Logger = logger.Default.LogMode(logger.Silent)
	}

	sqlDB.SetConnMaxLifetime(time.Minute * 5)
	sqlDB.SetMaxOpenConns(7)
	sqlDB.SetMaxIdleConns(5)

	return &DB{
		Connection: db,
		sql:        sqlDB,
		env:        env,
	}
}
