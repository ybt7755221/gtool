package service

import (
	et "pm-system/entities"
	"pm-system/models"
)

type WrProjectsService struct {
}
/**
 * 根据多条件查询数据-单条
 */
func (c *WrProjectsService) FindOne(conditions *et.WrProjects) (*et.WrProjects, error) {
	wrProjectsModel := models.WrProjectsModel{}
	return wrProjectsModel.FindOne(conditions)
}
/**
 * 根据多条件查询数据
 */
func (c *WrProjectsService) Find(conditions *et.WrProjects, pagination *et.Pagination) ([]et.WrProjects, error) {
	wrProjectsModel := models.WrProjectsModel{}
	wrProjectsPage, err := wrProjectsModel.Find(conditions, pagination)
	if err != nil {
		return nil, err
	}
	return wrProjectsPage, nil
}
/**
 * 获取所有
 */
func (c *WrProjectsService) FindAll(condition *et.WrProjects, sort map[string]string) ([]et.WrProjects, error) {
	wrProjectsModel := models.WrProjectsModel{}
	wrProjectsList, err := wrProjectsModel.FindAll(condition, sort)
	return wrProjectsList, err
}
/**
 * 根据多条件查询数据-分页
 */
func (c *WrProjectsService) FindPaging(conditions *et.WrProjects, pagination *et.Pagination) (*et.WrProjectsPageDao, error) {
	wrProjectsModel := models.WrProjectsModel{}
	wrProjectsPage, err := wrProjectsModel.FindPaging(conditions, pagination)
	if err != nil {
		return nil, err
	}
	return wrProjectsPage, nil
}

func (c *WrProjectsService) FindById(id int) (*et.WrProjects, error) {
	wrProjectsModel := models.WrProjectsModel{}
	return wrProjectsModel.GetById(id)
}

func (c *WrProjectsService) Insert(wrProjects *et.WrProjects) (err error) {
	wrProjectsModel := models.WrProjectsModel{}
	err = wrProjectsModel.Insert(wrProjects)
	if err != nil {
		return err
	}
	return nil
}

func (c *WrProjectsService) UpdateById(id int, wrProjects *et.WrProjects) (has int64, err error) {
	wrProjectsModel := models.WrProjectsModel{}
	has, err = wrProjectsModel.UpdateById(id, wrProjects)
	return
}