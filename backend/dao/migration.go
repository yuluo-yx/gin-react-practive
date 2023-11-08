package dao

import (
	"backend/model"
	"fmt"
)

func Migration() {
	err := _db.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate(
		&model.User{},
		&model.Message{},
	)
	if err != nil {
		fmt.Println("err: ", err)
	}

	return

}
