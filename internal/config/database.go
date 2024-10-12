package config

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go-fiber-project-template/internal/model/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

type DBConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
}

func (c *DBConfig) Read(viper *viper.Viper) {
	c.Host = viper.GetString("DB_HOST")
	c.User = viper.GetString("DB_USERNAME")
	c.Password = viper.GetString("DB_PASSWORD")
	c.DBName = viper.GetString("DB_NAME")
	c.Port = viper.GetString("DB_PORT")
}

func NewDB(viper *viper.Viper, log *logrus.Logger) *gorm.DB {
	db := &DBConfig{}
	db.Read(viper)

	idleConnection := 5
	maxConnection := 20
	maxLifeTimeConnection := 60

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Asia%%2FJakarta",
		db.User,
		db.Password,
		db.Host,
		db.Port,
		db.DBName)

	dbOpen, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:      logger.Default.LogMode(logger.Info),
		PrepareStmt: true,
	})

	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	err = dbOpen.AutoMigrate(&entities.Category{}, &entities.Product{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	connection, err := dbOpen.DB()
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	connection.SetMaxIdleConns(idleConnection)
	connection.SetMaxOpenConns(maxConnection)
	connection.SetConnMaxLifetime(time.Second * time.Duration(maxLifeTimeConnection))

	return dbOpen
}
