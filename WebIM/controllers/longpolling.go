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
	"github.com/beego/samples/WebIM/models"
)

// LongPollingController handles long polling requests.
type LongPollingController struct {
	baseController
}

// Join method handles GET requests for LongPollingController.
func (L *LongPollingController) Join() {
	// Safe check.
	uname := L.GetString("uname")
	if len(uname) == 0 {
		L.Redirect("/", 302)
		return
	}

	// Join chat room.
	Join(uname, nil)

	L.TplName = "longpolling.html"
	L.Data["IsLongPolling"] = true
	L.Data["UserName"] = uname
}

// Post method handles receive messages requests for LongPollingController.
func (L *LongPollingController) Post() {
	L.TplName = "longpolling.html"

	uname := L.GetString("uname")
	content := L.GetString("content")
	if len(uname) == 0 || len(content) == 0 {
		return
	}

	publish <- newEvent(models.EVENT_MESSAGE, uname, content)
}

// Fetch method handles fetch archives requests for LongPollingController.
func (L *LongPollingController) Fetch() {
	lastReceived, err := L.GetInt("lastReceived")
	if err != nil {
		return
	}

	events := models.GetEvents(int(lastReceived))
	if len(events) > 0 {
		L.Data["json"] = events
		L.ServeJSON()
		return
	}

	// Wait for new message(s).
	ch := make(chan bool)
	waitingList.PushBack(ch)
	<-ch

	L.Data["json"] = models.GetEvents(int(lastReceived))
	L.ServeJSON()
}
