package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ToDo struct {
	ID   uint64 `gorm:"primaryKey"`
	Code string `sql:"NOT NULL"`
	Type bool   `sql:"DEFAULT:false"`
}

var DB *gorm.DB

func new_db(db *gorm.DB) {
	DB = db
}
func InitsilizeDb() {
	DB := init_Db()
	new_db(DB)
}
func init_Db() *gorm.DB {
	dsn := "host=localhost user=postgres password=2002 dbname=test port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	DbConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error con to db")
	}
	fmt.Println("Successfullt connected to database")
	// Migrate the schema
	DbConn.AutoMigrate(&ToDo{})
	return DbConn
}

func Get_Values() []ToDo {
	db := DB
	var todos []ToDo
	db.Find(&todos)
	return todos
}
