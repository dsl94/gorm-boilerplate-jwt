package database

import "github.com/jinzhu/gorm"

func Connect() *gorm.DB {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=testGorm3 password=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}

	return db
}
