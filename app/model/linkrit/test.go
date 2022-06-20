package linkrit

import "linkr-frame/app/model"

type Test struct {
	// id
	Id int64 `gorm:"column:id" json:"id"`

	//测试名称  
	Username string `gorm:"column:username" json:"username"`

	//测试密码  
	Password string `gorm:"column:password" json:"password"`

	//继承父类model
	model.Model
}

func (model *Test) TableName() string {
	return "test"
}
