package templates

var ControllerTpl = `package controllers

import (
	et "gpi/entities"
	"gpi/service"
	"github.com/gin-gonic/gin"
	"strconv"
)
type {{StructName}}Controller struct {
	serv *service.{{StructName}}Service
}
// @Tags {{StructRoute}}表操作
// @Summary 【GetAll】根据条件获取信息
// @Description 根据条件获取信息
// @Accept html
// @Produce json
// @Param	page_num	query 	int		false	"页数，默认1"
// @Param	page_size	query 	int		false	"每夜条数，默认50"
// @Param	sort		query 	string	false	"排序。id desc, time asc"
// @Success 200 {object} SgrResp
// @Router /{{StructRoute}} [get]
func (c *{{StructName}}Controller) Find(ctx *gin.Context) {
	fieldsArr := []string{}
	//处理分页参数
	params := getPagingParams(ctx)
	//处理查询条件
	params["conditions"] = getParams(ctx, fieldsArr, et.{{StructName}}{})
	{{StructFcName}}List, err := c.serv.Find(params)
	if err != nil {
		resError(ctx, et.EntityFailure, err.Error())
		return
	}
	resSuccess(ctx, {{StructFcName}}List)
}
// @Tags {{StructRoute}}表操作
// @Summary 【GetOne】根据id获取信息
// @Description 根据id获取信息
// @Accept html
// @Produce json
// @Param   id		path	string 	false	"主键id"
// @Success 200 {object} SgrResp
// @Router /{{StructRoute}}/{id} [get]
func (c *{{StructName}}Controller) FindById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	{{StructFcName}}, err := c.serv.FindById(id)
	if err != nil {
		resError(ctx, et.EntityFailure, err.Error())
	}else{
		resSuccess(ctx, {{StructFcName}})
	}
}
// @Tags {{StructRoute}}表操作
// @Summary 【create】创建{{StructRoute}}信息
// @Description 创建{{StructRoute}}信息
// @Accept x-www-form-urlencoded
// @Produce json
// @Success 200 {object} SgrResp
// @Router /{{StructRoute}} [post]
func (c *{{StructName}}Controller) Create(ctx *gin.Context) {
	{{StructFcName}} := new(et.{{StructName}})
	getPostStructData(ctx, {{StructFcName}})
	if err := c.serv.Insert({{StructFcName}}); err != nil {
		resError(ctx, et.EntityFailure, err.Error())
		return
	}
	resSuccess(ctx, {{StructFcName}})
}
// @Tags {{StructRoute}}表操作
// @Summary 【update】根据id更新数据
// @Description 根据id更新数据
// @Accept x-www-form-urlencoded
// @Produce json
// @Param   id	body	string 	true	"主键更新依据此id"
// @Success 200 {object} SgrResp
// @Router /{{StructRoute}}/update-by-id [put]
func (c * {{StructName}}Controller) UpdateById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.PostForm("id"))
	{{StructFcName}} := new(et.{{StructName}})
	getPostStructData(ctx, {{StructFcName}})
	has, err := c.serv.UpdateById(id, {{StructFcName}})
	if err != nil {
		resError(ctx, et.EntityFailure, err.Error())
	}else{
		if has == 0 {
			resError(ctx, et.EntityFailure, "影响行数0")
		}else{
			resSuccess(ctx, gin.H{
				"update_count":has,
			})
		}
	}
}`
