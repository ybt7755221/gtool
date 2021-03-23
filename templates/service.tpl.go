package templates

var ServiceTpl = `package service

import (
	et "pm-system/entities"
	"pm-system/models"
)

type {{StructName}}Service struct {
}
/**
 * 根据多条件查询数据-单条
 */
func (c *{{StructName}}Service) FindOne(conditions *et.{{StructName}}) (*et.{{StructName}}, error) {
	{{StructLcName}}Model := models.{{StructName}}Model{}
	return {{StructLcName}}Model.FindOne(conditions)
}
/**
 * 根据多条件查询数据
 */
func (c *{{StructName}}Service) Find(conditions *et.{{StructName}}, pagination *et.Pagination) ([]et.{{StructName}}, error) {
	{{StructLcName}}Model := models.{{StructName}}Model{}
	{{StructLcName}}Page, err := {{StructLcName}}Model.Find(conditions, pagination)
	if err != nil {
		return nil, err
	}
	return {{StructLcName}}Page, nil
}
/**
 * 获取所有
 */
func (c *{{StructName}}Service) FindAll(condition *et.{{StructName}}, sort map[string]string) ([]et.{{StructName}}, error) {
	{{StructLcName}}Model := models.{{StructName}}Model{}
	{{StructLcName}}List, err := {{StructLcName}}Model.FindAll(condition, sort)
	return {{StructLcName}}List, err
}
/**
 * 根据多条件查询数据-分页
 */
func (c *{{StructName}}Service) FindPaging(conditions *et.{{StructName}}, pagination *et.Pagination) (*et.{{StructName}}PageDao, error) {
	{{StructLcName}}Model := models.{{StructName}}Model{}
	{{StructLcName}}Page, err := {{StructLcName}}Model.FindPaging(conditions, pagination)
	if err != nil {
		return nil, err
	}
	return {{StructLcName}}Page, nil
}

func (c *{{StructName}}Service) FindById(id int) (*et.{{StructName}}, error) {
	{{StructLcName}}Model := models.{{StructName}}Model{}
	return {{StructLcName}}Model.GetById(id)
}

func (c *{{StructName}}Service) Insert({{StructLcName}} *et.{{StructName}}) (err error) {
	{{StructLcName}}Model := models.{{StructName}}Model{}
	err = {{StructLcName}}Model.Insert({{StructLcName}})
	if err != nil {
		return err
	}
	return nil
}

func (c *{{StructName}}Service) UpdateById(id int, {{StructLcName}} *et.{{StructName}}) (has int64, err error) {
	{{StructLcName}}Model := models.{{StructName}}Model{}
	has, err = {{StructLcName}}Model.UpdateById(id, {{StructLcName}})
	return
}`
