package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"yuyuid.id/config"
)

func ConnectGORM(conf config.Database) *gorm.DB {

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // output ke stdout
		logger.Config{
			SlowThreshold: time.Millisecond * 200, // threshold query lambat
			LogLevel:      logger.Info,            // tampilkan info
			Colorful:      true,                   // kalau console support warna
		},
	)

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable Timezone=%s",
		conf.Host,
		conf.Port,
		conf.User,
		conf.Pass,
		conf.Name,
		conf.Tz,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
		panic("failed to connect to database")
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err.Error())
	}
	sqlDB.SetMaxIdleConns(10) // Atur jumlah koneksi idle
	sqlDB.SetMaxOpenConns(25) // Atur jumlah koneksi terbuka
	sqlDB.SetConnMaxLifetime(time.Hour)
	return db
}
