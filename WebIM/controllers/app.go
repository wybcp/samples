// Copyright 2013 Beego Samples authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package controllers

import (
	"strings"

	"github.com/astaxie/beego"
	"github.com/beego/i18n"
)

var langTypes []string // Languages that are supported.

func init() {
	// Initialize language type list.
	langTypes = strings.Split(beego.AppConfig.String("lang_types"), "|")

	// Load locale files according to language types.
	for _, lang := range langTypes {
		beego.Trace("Loading language: " + lang)
		if err := i18n.SetMessage(lang, "conf/"+"locale_"+lang+".ini"); err != nil {
			beego.Error("Fail to set message file:", err)
			return
		}
	}
}

// baseController represents base router for all other app routers.
// It implemented some methods for the same implementation;
// thus, it will be embedded into other routers.
type baseController struct {
	beego.Controller // Embed struct that has stub implementation of the interface.
	i18n.Locale      // For i18n usage when process data and render template.
}

// Prepare implemented Prepare() method for baseController.
// It's used for language option check and setting.
func (b *baseController) Prepare() {
	// Reset language option.
	b.Lang = "" // this field is from i18n.Locale.

	// 1. Get language information from 'Accept-Language'.
	al := b.Ctx.Request.Header.Get("Accept-Language")
	if len(al) > 4 {
		al = al[:5] // Only compare first 5 letters.
		if i18n.IsExist(al) {
			b.Lang = al
		}
	}

	// 2. Default language is English.
	if len(b.Lang) == 0 {
		b.Lang = "en-US"
	}

	// Set template level language option.
	b.Data["Lang"] = b.Lang
}

// AppController handles the welcome screen that allows user to pick a technology and username.
type AppController struct {
	baseController // Embed to use methods that are implemented in baseController.
}

// Get implemented Get() method for AppController.
func (a *AppController) Get() {
	a.TplName = "welcome.html"
}

// Join method handles POST requests for AppController.
func (a *AppController) Join() {
	// Get form value.
	uname := a.GetString("uname")
	tech := a.GetString("tech")

	// Check valid.
	if len(uname) == 0 {
		a.Redirect("/", 302)
		return
	}

	switch tech {
	case "longpolling":
		a.Redirect("/lp?uname="+uname, 302)
	case "websocket":
		a.Redirect("/ws?uname="+uname, 302)
	default:
		a.Redirect("/", 302)
	}

	// Usually put return after redirect.
	return
}
