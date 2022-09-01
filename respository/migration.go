package respository

import (
	"errors"
	"fmt"
	"log"
)

func migration() {
	//自动迁移
	err := DB.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(
			&UserBasic{},
			&QuestionBasic{},
			&Answer{},
			&Comment{},
			&IsLike{},
			IsAttention{},
		)
	if err != nil {
		log.Fatal(errors.New("register table fail"))
	}
	fmt.Println("register table success")
}
