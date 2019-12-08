package helpers

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func DBConnect() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@/go-blog")

	ValidateError(err)

	fmt.Println("Db Connected.....")

	return db
}
