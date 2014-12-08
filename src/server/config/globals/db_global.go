package global

import (
	"flag"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

type DatabaseConnection struct {
	Connection gorm.DB
}

var DB *DatabaseConnection = nil

func InitDB() {
	DB = new(DatabaseConnection)
	DB.DBConnect()
}

// Private Methods

func (db_struct *DatabaseConnection) DBConnect() {
	var db gorm.DB
	var err error

	var db_type = flag.String("db_type", "sqlite3", "DB Type, ex: sqlite3/mysql/postgres")
	var db_name = flag.String("db_name", "./tmp/db.db", "DB Name or Path")
	var db_user = flag.String("db_user", "", "DB User, only for MySQL/PostGres")
	var db_pass = flag.String("db_pass", "", "DB Pass, only for MySQL/PostGres")
	var db_host = flag.String("db_host", "", "DB Host, only for MySQL/PostGres")

	flag.Parse()

	if *db_type == "mysql" {
		db, err = gorm.Open("mysql", *db_user+":"+*db_pass+"@"+*db_host+"/"+*db_name+"?charset=utf8&parseTime=True")
	} else if *db_type == "postgres" {
		db, err = gorm.Open("postgres", "host= "+*db_host+" user="+*db_user+" pass="+*db_pass+" dbname="+*db_name+" sslmode=enable")
	} else {
		db, err = gorm.Open("sqlite3", "./tmp/db.db")
	}

	// Then you could invoke `*sql.DB`'s functions with it
	if err == nil {
		db.DB().Ping()
		db.DB().SetMaxIdleConns(10)
		db.DB().SetMaxOpenConns(100)

		db_struct.Connection = db
	}
}

// Public Methods

func (db_struct *DatabaseConnection) Debug() {
	db_struct.Connection.LogMode(true)
}
