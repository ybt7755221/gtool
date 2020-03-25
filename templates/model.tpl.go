package templates

var ModelTpl = `package models

import (
	. "gpi/entities"
	DB "gpi/libraries/database"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type {{StructName}}Model struct {
}
//查找多条数据
func (u *{{StructName}}Model) Find(params map[string]interface{}) ([]{{StructName}}, error) {
	dbConn := DB.GetDB({{StructDB}})
	defer dbConn.Close()
	{{StructLcName}} := make([]{{StructName}}, 0)
	dbC := dbConn.Where("1")
	defer dbC.Close()
	reflect.TypeOf(params["conditions"])
	//where条件
	conditions := params["conditions"].(map[string]string)
	if len(conditions) > 0 {
		for key, val := range params["conditions"].(map[string]string) {
			if len(val) > 0 {
				dbC = dbC.And(key+" = ?", val)
			}
		}
	}
	//limit
	dbC = dbC.Limit(params["limit"].(int), params["offset"].(int))
	if params["sortField"] == "" {
		params["sortField"] = "id"
	}
	//排序
	sort := params["sort"].(map[string]string)
	fmt.Println(len(sort))
	if len(sort) > 0 {
		for key, val := range sort{
			if strings.ToLower(val) == "asc" {
				dbC = dbC.Asc(key)
			}else{
				dbC = dbC.Desc(key)
			}
		}
	}
	err := dbC.Find(&{{StructLcName}})
	return {{StructLcName}}, err
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
