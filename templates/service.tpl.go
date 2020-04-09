package templates

var ServiceTpl = `package service

import (
	et "gpi/entities"
	"gpi/models"
)

type {{StructName}}Service struct {
}
/**
 * 根据多条件查询数据
 */
func (c *{{StructName}}Service) Find(conditions *et.{{StructName}}, pagination *et.Pagination) (*et.{{StructName}}PageDao, error) {
	{{StructLcName}}Model := models.{{StructName}}Model{}
	{{StructLcName}}Page, err := {{StructLcName}}Model.Find(conditions, pagination)
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
