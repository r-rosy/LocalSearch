package model

import (
	"fmt"
)

type Maps struct {
	Id 			int			`json:"id" gorm:"column:id"`
	Host        string      `json:"host" gorm:"column:host"`
	Guest		string      `json:"guest" gorm:"column:guest"`
}

func Add(host string,guest string) error{
	query := fmt.Sprintf("%s = ?", "host")
	if err := MysqlDb.Db.Where(query,host).Error;err!=nil {
		return err
	}

	if err:=MysqlDb.Db.Create(&Maps{Host: host,Guest: guest}).Error;err!=nil {
		return err
	}
	return nil
}

func DelGuest(id int) error{
	query := fmt.Sprintf("%s = ?", "id")
	err := MysqlDb.Db.Where(query, id).Delete(&Maps{}).Error
	if err!=nil {
		return err
	}
	return nil
}

func DelHost(host string) error{
	query := fmt.Sprintf("%s = ?", "host")
	err := MysqlDb.Db.Where(query, host).Delete(&Maps{}).Error
	if err!=nil {
		return err
	}
	return nil
}

func Sea(host string) ([]string,error){
	maps:=[]Maps{}
	query := fmt.Sprintf("%s = ?", "host")
	err := MysqlDb.Db.Where(query, host).Find(maps).Error
	if err!=nil {
		return nil,err
	}

	result:=[]string{}
	for _,Map:= range maps {
		result=append(result,Map.Guest)
	}
	return result,nil
}