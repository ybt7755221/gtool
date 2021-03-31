package controllers

import (
	et "pm-system/entities"
	"pm-system/service"
	"github.com/gin-gonic/gin"
	"strconv"
)
type WrWorksController struct {
	serv *service.WrWorksService
}
// @Tags user表操作
// @Summary 【GetOne】根据条件获取单条数据
// @Description 根据条件获取信息
// @Accept html
// @Produce json
// @Success 200 {object} SgrResp
// @Router /user/one [get]
func (c *WrWorksController) FindOne(ctx *gin.Context) {
	wrWorks := new(et.WrWorks)
	getParamsNew(ctx, wrWorks)
	res, err := c.serv.FindOne(wrWorks)
	if err != nil {
		resError(ctx, et.EntityFailure, err.Error())
	}else{
		resSuccess(ctx, res)
	}
}
// @Tags works表操作
// @Summary 【GetAll】根据条件获取信息
// @Description 根据条件获取信息
// @Accept html
// @Produce json
// @Param	page_num	query 	int		false	"页数，默认1"
// @Param	page_size	query 	int		false	"每夜条数，默认50"
// @Param	sort		query 	string	false	"排序 {\"id\":\"desc\"}"
// @Success 200 {object} SgrResp
// @Router /works [get]
func (c *WrWorksController) Find(ctx *gin.Context) {
	wrWorks := new(et.WrWorks)
	getParamsNew(ctx, wrWorks)
	pagination := new(et.Pagination)
	pagination.PageNum, _ = strconv.Atoi(ctx.Query("page_num"))
	pagination.PageSize, _ = strconv.Atoi(ctx.Query("page_size"))
	pagination.SortStr = ctx.Query("sort")
	wrWorksList, err := c.serv.Find(wrWorks, pagination)
	if err != nil {
		resError(ctx, et.EntityFailure, err.Error())
		return
	}
	resSuccess(ctx, wrWorksList)
}
// @Tags works表操作
// @Summary 【GetAll】根据条件获取信息
// @Description 根据条件获取信息
// @Accept html
// @Produce json
// @Param	page_num	query 	int		false	"页数，默认1"
// @Param	page_size	query 	int		false	"每夜条数，默认50"
// @Param	sort		query 	string	false	"排序 {\"id\":\"desc\"}"
// @Success 200 {object} SgrResp
// @Router /works/page [get]
func (c *WrWorksController) FindPaging(ctx *gin.Context) {
	wrWorks := new(et.WrWorks)
	getParamsNew(ctx, wrWorks)
	pagination := new(et.Pagination)
	pagination.PageNum, _ = strconv.Atoi(ctx.Query("page_num"))
	pagination.PageSize, _ = strconv.Atoi(ctx.Query("page_size"))
	pagination.SortStr = ctx.Query("sort")
	wrWorksList, err := c.serv.FindPaging(wrWorks, pagination)
	if err != nil {
		resError(ctx, et.EntityFailure, err.Error())
		return
	}
	resSuccess(ctx, wrWorksList)
}
// @Tags works表操作
// @Summary 【GetOne】根据id获取信息
// @Description 根据id获取信息
// @Accept html
// @Produce json
// @Param   id		path	string 	false	"主键id"
// @Success 200 {object} SgrResp
// @Router /works/find-by-id/{id} [get]
func (c *WrWorksController) FindById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	wrWorks, err := c.serv.FindById(id)
	if err != nil {
		resError(ctx, et.EntityFailure, err.Error())
	}else{
		resSuccess(ctx, wrWorks)
	}
}
// @Tags works表操作
// @Summary 【create】创建works信息
// @Description 创建works信息
// @Accept x-www-form-urlencoded
// @Produce json
// @Success 200 {object} SgrResp
// @Router /works [post]
func (c *WrWorksController) Create(ctx *gin.Context) {
	wrWorks := new(et.WrWorks)
	getPostStructData(ctx, wrWorks)
	if err := c.serv.Insert(wrWorks); err != nil {
		resError(ctx, et.EntityFailure, err.Error())
		return
	}
	resSuccess(ctx, wrWorks)
}
// @Tags works表操作
// @Summary 【update】根据id更新数据
// @Description 根据id更新数据
// @Accept x-www-form-urlencoded
// @Produce json
// @Param   id	body	string 	true	"主键更新依据此id"
// @Success 200 {object} SgrResp
// @Router /works/update-by-id [put]
func (c * WrWorksController) UpdateById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.PostForm("id"))
	wrWorks := new(et.WrWorks)
	getPostStructData(ctx, wrWorks)
	has, err := c.serv.UpdateById(id, wrWorks)
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
}

func (c *WrWorksController) FindAll(ctx *gin.Context) {
	wrWorks := new(et.WrWorks)
	getParamsNew(ctx, wrWorks)
	sort := getSort(ctx.Query("sort"))
	wrWorksList, err := c.serv.FindAll(wrWorks, sort)
	if err != nil {
		resError(ctx, et.EntityFailure, err.Error())
		return
	}
	resSuccess(ctx, wrWorksList)
}
