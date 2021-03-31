package controller

import "github.com/gin-gonic/gin"

type tag struct {
}

type TagData struct {
	ID   int    `json:"id" example:"1" `
	Name string `json:"name" example:"测试"`
}
type HTTPError struct {
	Code int         `json:"code" example:"400"`
	Msg  string      `json:"msg" example:"status bad request"`
	Data interface{} `json:"data" example:""`
}
type HTTPOk struct {
	Code int         `json:"code" example:"0"`
	Msg  string      `json:"msg" example:""`
	Data interface{} `json:"data" example:""`
}

var Tag tag

// @Summary 获取标签
// @Tags Tag
// @Description 通过id获取标签
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param id query int true "标签id"
// @Success 200 {object} TagData
// @Failure 400 {object} HTTPError
// @Router /tag [get]
func (t *tag) Get(c *gin.Context) {

}

// @Summary 添加标签
// @Tags Tag
// @Description 添加标签
// @Accept  json
// @Produce  json
// @Param account body TagData true "标签内容"
// @Success 200 {object} HTTPOk
// @Failure 400 {object} HTTPError
// @Router /tag [POST]
func (t *tag) Add(c *gin.Context) {

}

func (t *tag) Update(c *gin.Context) {}

func (t *tag) Del(c *gin.Context) {

}
