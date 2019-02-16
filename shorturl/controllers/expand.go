package controllers

import (
	"github.com/astaxie/beego"
)

// ExpandController 结构
type ExpandController struct {
	beego.Controller
}

// Get 获取
func (E *ExpandController) Get() {
	var result ShortResult
	shorturl := E.Input().Get("shorturl")
	result.URLShort = shorturl
	if urlcache.IsExist(shorturl) {
		result.URLLong = urlcache.Get(shorturl).(string)
	} else {
		result.URLLong = ""
	}
	E.Data["json"] = result
	E.ServeJSON()
}
