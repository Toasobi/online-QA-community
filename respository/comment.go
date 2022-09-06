package respository

import (
	"online-QA-community/helper"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	QuestionIdentity string `gorm:"column:question_identity;type:varchar(36);"json:"question_identity"` //问题表的唯一标识
	AnswerIdentity   string `gorm:"column:answer_identity;type:varchar(36);"json:"answer_identity"`     //文章唯一标识
	Identity         string `gorm:"column:identity;type:varchar(36);"json:"identity"`                   //评论表的唯一标识
	Content          string `gorm:"column:content;type:text;"json:"content"`                            //评论正文
	Name             string `gorm:"column:name;type:varchar(100);"json:"name"`                          //回答文章的用户名
	Like             int    `gorm:"column:like;type:int(11);"json:"like"`                               //点赞数
}

func (table *Comment) TableName() string {
	return "comment"
}

func CreateComment(content, qid, aid, name string) (*Comment, error) {
	var comment *Comment
	comment = &Comment{
		QuestionIdentity: qid,
		AnswerIdentity:   aid,
		Identity:         helper.GenerateUUID(),
		Content:          content,
		Name:             name,
	}
	if err := DB.Create(&comment).Error; err != nil {
		return &Comment{}, err
	}
	return comment, nil
}

func GetCommentList(name string) *gorm.DB {
	tx := DB.Model(new(Comment)).Where("name = ?", name)
	return tx
}
