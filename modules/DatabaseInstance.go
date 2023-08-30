package modules

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func GetDatabaseInstance() *gorm.DB {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
		panic(err)
		return nil
	}
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASS")
	HOST := os.Getenv("DB_HOST")
	PORT := os.Getenv("DB_PORT")
	DBNAME := os.Getenv("DB_NAME")

	createDBDsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, PORT)

	database, err := gorm.Open(mysql.Open(createDBDsn), &gorm.Config{})

	_ = database.Exec("CREATE DATABASE IF NOT EXISTS " + DBNAME + ";")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	//Connect and migrate database tables

	if err := db.AutoMigrate(&User{}, &Website{}, &Manga{}, &Chapter{}, &WebsitePattern{}); err != nil {
		fmt.Printf("Could not Migrate data\n%v\n", err)
		panic(err)
		return nil
	}

	if err != nil {
		panic(err.Error())
	}

	return db
}
