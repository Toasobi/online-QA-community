package respository

import "gorm.io/gorm"

type UserBasic struct {
	gorm.Model        //通用的
	Identity   string `gorm:"column:identity;type:varchar(36);"json:"identity"` //用户表的唯一标识
	Name       string `gorm:"column:name;type:varchar(100);"json:"name"`        //用户名
	Password   string `gorm:"column:password;type:varchar(32);"json:"password"` //密码
	Phone      string `gorm:"column:phone;type:varchar(20);"json:"phone"`       //手机号
	Mail       string `gorm:"column:mail;type:varchar(100);"json:"mail"`        //邮箱
	IsAdmin    int    `gorm:"column:is_admin;type:int(11);"json:"is_admin"`     //是否是管理员 0不是 1是
	Attention  int    `gorm:"column:attention;type:int(11);"json:"attention"`   //关注数
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}

//创建用户并存入数据库
func CreateUser(username, identity, password, phone, mail string, isAdmin int) (*UserBasic, error) {
	user := &UserBasic{
		Identity: identity,
		Name:     username,
		Password: password,
		Phone:    phone,
		Mail:     mail,
		IsAdmin:  isAdmin,
	}
	if err := DB.Create(&user).Error; err != nil {
		return &UserBasic{}, err
	}
	return user, nil
}
