package orm

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)
type User struct {
	gorm.Model
	Name string
	Sex bool
	Age int
}
func InitDatabase()  {
	db, err := gorm.Open("mysql", "root:mysqladmin@/test_jdbc?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&User{})
	db.Where("name = ?", "jean").Find(&[]User{}).Updates(map[string]interface{}{
		"Name": "rui",
		"Sex": false,
	})
	db.Where("name = ?", "shun").Find(&[]User{}).Unscoped().Delete(&[]User{})
}
