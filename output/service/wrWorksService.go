package service

import (
	et "pm-system/entities"
	"pm-system/models"
)

type WrWorksService struct {
}
/**
 * 根据多条件查询数据-单条
 */
func (c *WrWorksService) FindOne(conditions *et.WrWorks) (*et.WrWorks, error) {
	wrWorksModel := models.WrWorksModel{}
	return wrWorksModel.FindOne(conditions)
}
/**
 * 根据多条件查询数据
 */
func (c *WrWorksService) Find(conditions *et.WrWorks, pagination *et.Pagination) ([]et.WrWorks, error) {
	wrWorksModel := models.WrWorksModel{}
	wrWorksPage, err := wrWorksModel.Find(conditions, pagination)
	if err != nil {
		return nil, err
	}
	return wrWorksPage, nil
}
/**
 * 获取所有
 */
func (c *WrWorksService) FindAll(condition *et.WrWorks, sort map[string]string) ([]et.WrWorks, error) {
	wrWorksModel := models.WrWorksModel{}
	wrWorksList, err := wrWorksModel.FindAll(condition, sort)
	return wrWorksList, err
}
/**
 * 根据多条件查询数据-分页
 */
func (c *WrWorksService) FindPaging(conditions *et.WrWorks, pagination *et.Pagination) (*et.WrWorksPageDao, error) {
	wrWorksModel := models.WrWorksModel{}
	wrWorksPage, err := wrWorksModel.FindPaging(conditions, pagination)
	if err != nil {
		return nil, err
	}
	return wrWorksPage, nil
}

func (c *WrWorksService) FindById(id int) (*et.WrWorks, error) {
	wrWorksModel := models.WrWorksModel{}
	return wrWorksModel.GetById(id)
}

func (c *WrWorksService) Insert(wrWorks *et.WrWorks) (err error) {
	wrWorksModel := models.WrWorksModel{}
	err = wrWorksModel.Insert(wrWorks)
	if err != nil {
		return err
	}
	return nil
}

func (c *WrWorksService) UpdateById(id int, wrWorks *et.WrWorks) (has int64, err error) {
	wrWorksModel := models.WrWorksModel{}
	has, err = wrWorksModel.UpdateById(id, wrWorks)
	return
}