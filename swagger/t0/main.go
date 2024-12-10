package main

// @title fck-swagger
// @license.name Apache 2.0
// @version 1.0
// @Description Fucking description
// @Security ApiKeyAuth
// @host localhost:8080
// @BasePath /v1

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
	_ "swg/docs"
	_ "swg/errcode"
)

// @summary main function summary
// @description main function description
// @tags main
// @router /ping [get]
func main() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/ping", fuack)
	r.GET("/ping2", fuack2)

	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")

}

func fuack(c *gin.Context) {
	c.String(200, "pong")
}

func fuack2(c *gin.Context) {
	c.String(200, "pongpong")
}

type Tag struct{}

// @Summary 获取多个标签
// @Produce  json
// @Param name query string false "标签名称" maxlength(100)
// @Param state query int false "状态" Enums(0, 1) default(1)
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.Tag "成功"
// @Router /api/v1/tags [get]
func (t Tag) List(c *gin.Context) {}

// @Summary 新增标签
// @Produce  json
// @Param name body string true "标签名称" minlength(3) maxlength(100)
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param created_by body string true "创建者" minlength(3) maxlength(100)
// @Success 200 {object} model.Tag "成功"
// @Router /api/v1/tags [post]
func (t Tag) Create(c *gin.Context) {}

// @Summary 更新标签
// @Produce  json
// @Param id path int true "标签 ID"
// @Param name body string false "标签名称" minlength(3) maxlength(100)
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param modified_by body string true "修改者" minlength(3) maxlength(100)
// @Success 200 {array} model.Tag "成功"
// @Router /api/v1/tags/{id} [put]
func (t Tag) Update(c *gin.Context) {}

// @Summary 删除标签
// @Produce  json
// @Param id path int true "标签 ID"
// @Success 200 {string} string "成功"
// @Router /api/v1/tags/{id} [delete]
func (t Tag) Delete(c *gin.Context) {}
