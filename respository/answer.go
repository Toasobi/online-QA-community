package respository

import (
	"online-QA-community/helper"

	"gorm.io/gorm"
)

type Answer struct {
	gorm.Model
	QuestionIdentity string `gorm:"column:question_identity;type:varchar(36);"json:"question_identity";` //问题表的唯一标识
	Identity         string `gorm:"column:identity;type:varchar(36);"json:"identity"`                    //回答表的唯一标识
	Content          string `gorm:"column:content;type:text;"json:"content"`                             //回答正文
	Name             string `gorm:"column:name;type:varchar(100);"json:"name"`                           //回答提问的用户名
	Like             int    `gorm:"column:like;type:int(11);"json:"like"`                                //点赞数
}

func (table *Answer) TableName() string {
	return "answer"
}

//创建回答并存入数据库
func CreateAnswer(QaId, content, name string) (*Answer, error) {
	var answer *Answer
	answer = &Answer{
		QuestionIdentity: QaId,
		Identity:         helper.GenerateUUID(),
		Content:          content,
		Name:             name,
	}
	if err := DB.Create(&answer).Error; err != nil {
		return &Answer{}, err
	}
	return answer, nil
}

func GetAnswerList(name string) *gorm.DB {
	tx := DB.Model(new(Answer)).Where("name = ?", name)
	return tx
}
