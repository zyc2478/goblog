package controllers

import (
	"fmt"
	"strconv"

	// "github.com/deepzz0/go-common/log"
	"github.com/deepzz0/goblog/helper"
	"github.com/deepzz0/goblog/models"
)

type CategoryController struct {
	Common
}

func (this *CategoryController) Get() {
	this.Layout = "homelayout.html"
	this.TplName = "groupTemplate.html"
	this.ListTopic()
}

func (this *CategoryController) ListTopic() {
	cat := this.Ctx.Input.Param(":cat")
	this.Leftbar(cat)
	category := models.Blogger.GetCategoryByID(cat)
	var name string = "暂无该分类"
	if category != nil && category.Extra != "" {
		name = category.Text
	}
	this.Data["Name"] = name
	this.Data["URL"] = fmt.Sprintf("%s/cat/%s", this.domain, category.ID)
	pageStr := this.Ctx.Input.Param(":page")
	page := 1
	if temp, err := strconv.Atoi(pageStr); err == nil {
		page = temp
	}
	topics, remainpage := models.TMgr.GetTopicsByCatgory(cat, page)
	if remainpage == -1 {
		this.Data["ClassOlder"] = "disabled"
		this.Data["UrlOlder"] = "#"
		this.Data["ClassNewer"] = "disabled"
		this.Data["UrlNewer"] = "#"
	} else {
		if page == 1 {
			this.Data["ClassOlder"] = "disabled"
			this.Data["UrlOlder"] = "#"
		} else {
			this.Data["ClassOlder"] = ""
			this.Data["UrlOlder"] = this.domain + "/cat/" + cat + fmt.Sprintf("/p/%d", page-1)
		}
		if remainpage == 0 {
			this.Data["ClassNewer"] = "disabled"
			this.Data["UrlNewer"] = "#"
		} else {
			this.Data["ClassNewer"] = ""
			this.Data["UrlNewer"] = this.domain + "/cat/" + cat + fmt.Sprintf("/p/%d", page+1)
		}
		var ts []*listOfTopic
		for _, topic := range topics {
			t := &listOfTopic{}
			t.ID = topic.ID
			t.Time = topic.CreateTime.Format(helper.Layout_y_m_d2)
			t.URL = fmt.Sprintf("%s/%s/%d.html", this.domain, topic.CreateTime.Format(helper.Layout_y_m_d), topic.ID)
			t.Title = topic.Title
			t.Preview = topic.Preview
			t.PCategory = topic.PCategory
			t.PTags = topic.PTags
			ts = append(ts, t)
		}
		this.Data["ListTopics"] = ts
	}
	this.Data["Title"] = name + " - " + models.Blogger.BlogName
}
