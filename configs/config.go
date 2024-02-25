package configs

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root@tcp(localhost:3306)/ecomm?charset=utf8&parseTime=True&loc=Local") // Jika Anda ingin menggunakan database MySQL, silakan sesuaikan dengan konfigurasi database Anda sesuai username, password, host, dan database yang digunakan, contohnya username: root, password: (kosong), host: localhost, dan database: ecomm
	if err != nil {
		return nil, err
	}
	DB = db
	return DB, nil
}