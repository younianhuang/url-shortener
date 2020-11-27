package controllers

import (
	beego "github.com/astaxie/beego"
)

type RedirectController struct {
	beego.Controller
}

func (this *RedirectController) Get() {
	shorturl := this.Ctx.Input.Param(":shorturl")

	var result ShortResult
	result.UrlShort = shorturl
	if urlcache.IsExist(shorturl) {
		result.UrlLong = urlcache.Get(shorturl).(string)
		this.Ctx.Redirect(302, result.UrlLong)
	} else {
		result.UrlLong = ""
		this.Data["json"] = result
		this.ServeJSON()
	}

	/*
		this.Ctx.Output.Body([]byte(shortURL))

		shorturl := this.Input().Get("shorturl")
		result.UrlShort = shorturl
		if urlcache.IsExist(shorturl) {
			result.UrlLong = urlcache.Get(shorturl).(string)
		} else {
			result.UrlLong = ""
		}
	*/
}
