package respository

import "gorm.io/gorm"

type IsAttention struct {
	gorm.Model
	Identity  string `gorm:"column:identity;type:varchar(36);"json:"identity"` //用户表的唯一标识
	Name      string `gorm:"column:name;type:varchar(100);"json:"name"`        //点赞的用户名
	Attention int    `gorm:"column:attention;type:int(11);"json:"attention"`   //关注标志位
}

func (table *IsAttention) TableName() string {
	return "is_attention"
}

func DisAttention(identity, name string) error {
	var user *UserBasic
	var att int
	if err := DB.Where("identity = ? AND name = ?", identity, name).Delete(new(IsAttention)).Error; err != nil {
		return err
	}
	DB.First(&user)
	att = user.Attention

	if err := DB.Model(new(UserBasic)).Where("identity = ?", identity).Update("attention", att-1).Error; err != nil {
		return err
	}
	return nil
}

func Attention(identity, name string) error {
	var isatt *IsAttention
	var user *UserBasic
	var att int
	isatt = &IsAttention{
		Identity:  identity,
		Name:      name,
		Attention: 1,
	}
	if err := DB.Create(&isatt).Error; err != nil {
		return err
	}
	DB.First(&user)
	att = user.Attention

	if err := DB.Model(new(UserBasic)).Where("identity = ?", identity).Update("attention", att+1).Error; err != nil {
		return err
	}
	return nil
}
