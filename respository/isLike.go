package respository

import (
	"gorm.io/gorm"
)

type IsLike struct {
	gorm.Model
	Identity string `gorm:"column:identity;type:varchar(36);"json:"identity"` //评论表的唯一标识
	Name     string `gorm:"column:name;type:varchar(100);"json:"name"`        //点赞的用户名
	Like     int    `gorm:"column:like;type:int(11);"json:"like"`             //点赞标志位
}

func (table *IsLike) TableName() string {
	return "is_like"
}

func AnswerDisLike(identity, name string) error {
	var answer *Answer
	var like int
	if err := DB.Where("identity = ? AND name = ?", identity, name).Delete(new(IsLike)).Error; err != nil {
		return err
	}
	DB.First(&answer)
	like = answer.Like

	if err := DB.Model(new(Answer)).Where("identity = ?", identity).Update("like", like-1).Error; err != nil {
		return err
	}
	return nil

}

func AnswerLike(identity, name string) error {
	var islike *IsLike
	var answer *Answer
	var like int
	islike = &IsLike{
		Identity: identity,
		Name:     name,
		Like:     1,
	}
	if err := DB.Create(&islike).Error; err != nil {
		return err
	}
	DB.First(&answer)
	like = answer.Like

	if err := DB.Model(new(Answer)).Where("identity = ?", identity).Update("like", like+1).Error; err != nil {
		return err
	}
	return nil
}
