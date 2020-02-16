package controllers

import (
	"W2OlineWinterAssignmentTest/models"
	"fmt"
)

type TagsController struct {
	BaseController
}

func (this *TagsController) Get(){
	tags := models.QueryTagsByStrings("tags")
	fmt.Println(models.GetTagsMap(tags))
	this.Data["Tags"] = models.GetTagsMap(tags)

	this.TplName = "tags.html"
}