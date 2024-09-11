package libs

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost)/mst_vip?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
