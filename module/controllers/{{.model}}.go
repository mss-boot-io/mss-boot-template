/*
 * @Author: lwnmengjing
 * @Date: 2022/3/10 22:43
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2022/3/10 22:43
 */

package controllers

import (
	"net/http"

	"github.com/casdoor/casdoor-go-sdk/auth"
	"github.com/gin-gonic/gin"

	"github.com/mss-boot-io/mss-boot/pkg/response"
	"github.com/mss-boot-io/mss-boot/pkg/response/curd"
	"github.com/mss-boot-io/mss-boot/service/tenant/form"
)

func init() {
	e := new({{.model | ToCamel}})
	e.TableName = "{{.model}}"
	e.Auth = true
	e.CreateReq = &form.{{.model | ToCamel}}CreateReq{}
	e.UpdateReq = &form.{{.model | ToCamel}}UpdateReq{}
	e.GetReq = &form.{{.model | ToCamel}}GetReq{}
	e.GetResp = &form.{{.model | ToCamel}}GetResp{}
	e.DeleteReq = &form.{{.model | ToCamel}}DeleteReq{}
	e.ListReq = &form.{{.model | ToCamel}}ListReq{}
	e.ListResp = make([]form.{{.model | ToCamel}}ListItem, 0)
	response.AppendController(e)
}

type {{.model | ToCamel}} struct {
	curd.DefaultController
}

// Create 创建
// @Summary 创建{{.model}}
// @Description 创建{{.model}}
// @Tags {{.model}}
// @Accept  application/json
// @Product application/json
// @Param data body form.{{.model}}CreateReq true "data"
// @Success 200 {object} response.Response
// @Router /{{.service | ToPath}}/api/v1/{{.model | ToPath}} [post]
// @Security Bearer

// Update 更新
// @Summary 更新{{.model}}
// @Description 更新{{.model}}
// @Tags {{.model}}
// @Accept  application/json
// @Product application/json
// @Param id path string true "id"
// @Param data body form.{{.model}}UpdateReq true "data"
// @Success 200 {object} response.Response
// @Router /{{.service | ToPath}}/api/v1/{{.model | ToPath}}/{id} [put]
// @Security Bearer

// Delete 删除
// @Summary 删除{{.model}}
// @Description 删除{{.model}}
// @Tags {{.model}}
// @Accept  application/json
// @Product application/json
// @Param id path string true "id"
// @Success 200 {object} response.Response
// @Router /{{.service | ToPath}}/api/v1/{{.model | ToPath}}/{id} [delete]
// @Security Bearer

// Get 获取
// @Summary 获取{{.model}}
// @Description 获取{{.model}}
// @Tags {{.model}}
// @Accept  application/json
// @Product application/json
// @Param id path string true "id"
// @Success 200 {object} response.Response{data=form.{{.model}}GetResp}
// @Router /{{.service | ToPath}}/api/v1/{{.model | ToPath}}/{id} [get]
// @Security Bearer

// List 列表
// @Summary 列表{{.model}}
// @Description 列表{{.model}}
// @Tags {{.model}}
// @Accept  application/json
// @Product application/json
// @Param name query string false "租户名称"
// @Param page query string false "当前页"
// @Param pageSize query string false "每页容量"
// @Success 200 {object} response.Page{data=[]form.{{.model}}ListItem}
// @Router /{{.service | ToPath}}/api/v1/{{.model | ToPath}} [get]
// @Security Bearer

func (e *{{.model | ToCamel}}) Other(r *gin.RouterGroup) {
}