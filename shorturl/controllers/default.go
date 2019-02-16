package controllers

import (
	"github.com/astaxie/beego"
)

// MainController 主控制器
type MainController struct {
	beego.Controller
}

// Get 获得
func (M *MainController) Get() {
	M.Ctx.Output.Body([]byte("shorturl"))
}
