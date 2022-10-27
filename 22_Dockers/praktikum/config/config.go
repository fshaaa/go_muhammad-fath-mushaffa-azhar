package config

import (
	"database/sql"
	"docker-problem/models"
	"fmt"
	m "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func InitDB() {
	config := map[string]string{
		"DB_Username": "root",
		"DB_Password": "019283",
		"DB_Port":     "3306",
		"DB_Host":     os.Getenv("DB_ADDRESS"),
		"DB_Name":     "crud_go",
	}
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config["DB_Username"],
		config["DB_Password"],
		config["DB_Host"],
		config["DB_Port"],
		config["DB_Name"],
	)
	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{
		AllowGlobalUpdate: true,
	})
	if err != nil {
		panic(err.Error())
	}
	InitMigration()
}

func InitMigration() {
	DB.AutoMigrate(&models.User{}, &models.Book{})
}

func InitDatabaseSql() *sql.DB {
	cfg := m.Config{
		User:   "root",
		Passwd: "root",
		Net:    "tcp",
		Addr:   "192.168.1.71:3306",
		DBName: "orm_aja",
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		panic(err)
	}

	return db
}

//func InitDBTest() {
//	config := map[string]string{
//		"DB_Username": "root",
//		"DB_Password": "019283",
//		"DB_Port":     "3306",
//		"DB_Host":     "localhost",
//		"DB_Name":     "crud_go_test",
//	}
//
//	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
//		config["DB_Username"],
//		config["DB_Password"],
//		config["DB_Host"],
//		config["DB_Port"],
//		config["DB_Name"],
//	)
//
//	var err error
//	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
//	if err != nil {
//		panic(err.Error())
//	}
//	InitMigrateTest()
//}
//
//func InitMigrateTest() {
//	DB.Migrator().DropTable(&models.User{})
//	DB.AutoMigrate(&models.User{})
//}
