package respository

import (
	"online-QA-community/helper"

	"gorm.io/gorm"
)

type QuestionBasic struct {
	gorm.Model        //通用的
	Identity   string `gorm:"column:identity;type:varchar(36);"json:"identity"` //问题表的唯一标识
	Name       string `gorm:"column:name;type:varchar(100);"json:"name"`        //发起提问的用户名
	Title      string `gorm:"column:title;type:varchar(255);"json:"title"`      //文章标题
	Content    string `gorm:"column:content;type:text;"json:"content"`          //问题正文
}

func (table *QuestionBasic) TableName() string {
	return "qusetion_basic"
}

//发起问题并存入数据库
func CreateQA(title, content, name string) (*QuestionBasic, error) {
	var qa *QuestionBasic
	qa = &QuestionBasic{
		Identity: helper.GenerateUUID(),
		Name:     name,
		Title:    title,
		Content:  content,
	}
	if err := DB.Create(&qa).Error; err != nil {
		return &QuestionBasic{}, err
	}
	return qa, nil
}

//查询列表
func GetProblemList(keyword, name string) *gorm.DB {
	//模糊查询
	tx := DB.Model(new(QuestionBasic)).Where("name = ? AND title like ? ", name, "%"+keyword+"%").Or("name = ? AND content like ?", name, "%"+keyword+"%")
	return tx
}
