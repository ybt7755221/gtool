package templates

var ModelTpl = `package models

import (
	. "{{productName}}/entities"
	DB "{{productName}}/libraries/database"
	"errors"
	"fmt"
	"strings"
)

type {{StructName}}Model struct {
}

//查找单条数据
func(u *{{StructName}}Model) FindOne(conditions *{{StructName}}) (*{{StructName}}, error) {
	dbConn := DB.GetDB(Gin)
	defer dbConn.Close()
	has, err := dbConn.Get(conditions)
	if has {
		return conditions, err
	}else{
		return nil, err
	}
}
//查找多条数据
func(u *{{StructName}}Model) Find(conditions *{{StructName}}, pagination *Pagination )  ([]{{StructName}}, error) {
	dbConn := DB.GetDB({{StructDB}})
	defer dbConn.Close()
	//获取分页信息
	pageinfo := getPagingParams(pagination)
	limit := pageinfo["limit"].(int)
	offset := pageinfo["offset"].(int)
	dbC := dbConn.Limit(limit, offset)
	defer dbC.Close()
	{{StructLcName}}Page := new({{StructName}}PageDao)
	{{StructLcName}}Page.PageNum = pagination.PageNum
	{{StructLcName}}Page.PageSize = pagination.PageSize
	//排序
	sort := pageinfo["sort"].(map[string]string)
	if len(sort) > 0 {
		for key, val := range sort{
			if strings.ToLower(val) == "asc" {
				dbC = dbC.Asc(key)
			}else{
				dbC = dbC.Desc(key)
			}
		}
	}
	//执行查找
	err := dbC.Find(&{{StructLcName}}Page.List, conditions)
	return {{StructLcName}}Page.List, err
}

//查找所有数据
func (u *{{StructName}}Model) FindAll(conditions *{{StructName}}, sort map[string]string) ([]{{StructName}}, error) {
	{{StructLcName}}List := new({{StructName}}PageDao).List
	dbConn := DB.GetDB(Gin).NoCache()
	//排序
	if len(sort) > 0 {
		for key, val := range sort {
			if strings.ToLower(val) == "asc" {
				dbConn = dbConn.Asc(key)
			} else {
				dbConn = dbConn.Desc(key)
			}
		}
	}
	err := dbConn.Find(&{{StructLcName}}List, conditions)
	defer dbConn.Close()
	return {{StructLcName}}List, err
}

//查找多条数据-分页
func(u *{{StructName}}Model) FindPaging(conditions *{{StructName}}, pagination *Pagination )  (*{{StructName}}PageDao, error) {
	dbConn := DB.GetDB({{StructDB}})
	defer dbConn.Close()
	//获取分页信息
	pageinfo := getPagingParams(pagination)
	limit := pageinfo["limit"].(int)
	offset := pageinfo["offset"].(int)
	dbC := dbConn.Limit(limit, offset)
	defer dbC.Close()
	{{StructLcName}}Page := new({{StructName}}PageDao)
	{{StructLcName}}Page.PageNum = pagination.PageNum
	{{StructLcName}}Page.PageSize = pagination.PageSize
	//排序
	sort := pageinfo["sort"].(map[string]string)
	if len(sort) > 0 {
		for key, val := range sort{
			if strings.ToLower(val) == "asc" {
				dbC = dbC.Asc(key)
			}else{
				dbC = dbC.Desc(key)
			}
		}
	}
	//执行查找
	err := dbC.Find(&{{StructLcName}}Page.List, conditions)
	total, err := dbC.Count(conditions)
	if err == nil {
		{{StructLcName}}Page.Total = total
	}
	return {{StructLcName}}Page, err
}
//根据id查找单条数据
func (u *{{StructName}}Model) GetById(id int) (*{{StructName}}, error) {
	fmt.Println(id)
	{{StructLcName}} := &{{StructName}}{Id: id}
	dbConn := DB.GetDB({{StructDB}})
	_, err := dbConn.Get({{StructLcName}})
	defer dbConn.Close()
	return {{StructLcName}}, err
}
//插入
func (u *{{StructName}}Model) Insert({{StructLcName}} *{{StructName}}) (err error) {
	dbConn := DB.GetDB({{StructDB}})
	affected, err := dbConn.Insert({{StructLcName}})
	defer dbConn.Close()
	if err != nil {
		return err
	}
	if affected < 1 {
		err = errors.New("插入影响行数: 0" )
		return err
	}
	return err
}
//根据id更新
func (u *{{StructName}}Model) UpdateById(id int, {{StructLcName}} *{{StructName}}) (affected int64, err error) {
	dbConn := DB.GetDB({{StructDB}})
	affected, err = dbConn.Id(id).Update({{StructLcName}})
	defer dbConn.Close()
	return
}`
