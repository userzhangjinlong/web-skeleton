package web

import (
	"encoding/json"
	"linkr-frame/app/model/linkrit"
	"linkr-frame/global"
)

type IndexLogic struct {
}

//这里是逻辑层
func (i *IndexLogic) Index() string {
	//example model
	//异步无问题
	//go func() {
	//	var testModel linkrit.Test
	//	testModel.Username = "12334"
	//	testModel.Password = "erwerewr"
	//	global.LinkItDB.Create(&testModel)
	//}()

	//同步测试
	var testModel linkrit.Test
	err := global.LinkItDB.Model(&testModel).Where("id = ?", 2442).Find(&testModel).Error
	if err != nil {
		return "查询错误"
	}
	val, err := json.Marshal(testModel)

	//example redis
	test := global.Redis.Get("/config/string/mysql").String()

	return "this is index!" + test + string(val)
}
