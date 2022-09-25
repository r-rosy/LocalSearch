package model

import (
	"fmt"
)

var test Maps

type Maps struct {
	Id    int    `json:"id" gorm:"column:id"`
	Host  string `json:"host" gorm:"column:host"`
	Guest string `json:"guest" gorm:"column:guest"`
}

func Add(host string, guest string) error {
	query := fmt.Sprintf("%s = ?", "host")
	if MysqlDb.Db.Where(query, host).First(&test); test.Host != "" {
		return fmt.Errorf("Aleardy_Exited")
	}

	if err := MysqlDb.Db.Create(&Maps{Host: host, Guest: guest}).Error; err != nil {
		return err
	}
	return nil
}

func DelGuest(id int) error {
	query := fmt.Sprintf("%s = ?", "id")
	err := MysqlDb.Db.Where(query, id).Delete(&Maps{}).Error
	if err != nil {
		return err
	}

	return nil
}

func DelHost(host string) error {
	query := fmt.Sprintf("%s = ?", "host")
	if MysqlDb.Db.Where(query, host).First(&test); test.Host == "" {
		return fmt.Errorf("NotExisted")
	}

	err := MysqlDb.Db.Where(query, host).Delete(&Maps{}).Error
	if err != nil {
		return err
	}

	return nil
}

func Sea(host string) ([]string, error) {
	var maps []Maps
	query := fmt.Sprintf("%s = ?", "host")
	err := MysqlDb.Db.Where(query, host).Find(&maps).Error
	if err != nil {
		return nil, err
	}

	var result []string
	for _, Map := range maps {
		result = append(result, Map.Guest)
	}
	return result, nil
}
