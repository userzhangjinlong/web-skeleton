package service

import "fmt"

type Example struct {
}

//通用service层代码
func (e *Example) TestService() {
	fmt.Println("this is service code!")
}
