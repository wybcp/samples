package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/beego/samples/shorturl/models"
)

var (
	urlcache cache.Cache
)

func init() {
	urlcache, _ = cache.NewCache("memory", `{"interval":0}`)
}

// ShortResult 结构
type ShortResult struct {
	URLShort string
	URLLong  string
}
// ShortController 短URL控制器
type ShortController struct {
	beego.Controller
}

// Get we can simulate easier in the browser
func (S *ShortController) Get() {
	var result ShortResult
	longurl := S.Input().Get("longurl")
	beego.Info(longurl)
	result.URLLong = longurl
	urlmd5 := models.GetMD5(longurl)
	beego.Info(urlmd5)
	if urlcache.IsExist(urlmd5) {
		result.URLShort = urlcache.Get(urlmd5).(string)
	} else {
		result.URLShort = models.Generate()
		err := urlcache.Put(urlmd5, result.URLShort, 0)
		if err != nil {
			beego.Info(err)
		}
		err = urlcache.Put(result.URLShort, longurl, 0)
		if err != nil {
			beego.Info(err)
		}
	}
	S.Data["json"] = result
	S.ServeJSON()
}
