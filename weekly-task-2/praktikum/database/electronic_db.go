package database

import (
	"praktikum-weekly-task-2/config"
	"praktikum-weekly-task-2/models"
)

func GetItems() (interface{}, error) {
	var items []models.Electronics

	err := config.DB.Find(&items).Error
	if err != nil {
		return nil, err
	}
	return items, nil
}

func GetItem(id string) (interface{}, error) {
	var item models.Electronics

	err := config.DB.Find(&item, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return item, nil
}

func CreateItem(item *models.Electronics) error {
	err := config.DB.Create(&item).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateItem(id string, item *models.Electronics) error {
	err := config.DB.Where("id = ?", id).Updates(&item).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteItem(id string) error {
	var items []models.Electronics

	err := config.DB.Where("id = ?", id).Delete(&items).Error
	if err != nil {
		return err
	}
	return nil
}

func GetCategory(cat_id int) (interface{}, error) {
	var items []models.Electronics

	err := config.DB.Where("kategori_Barang = ?", cat_id).Find(&items).Error
	if err != nil {
		return nil, err
	}
	return items, nil
}

func GetItemName(keyword string) (interface{}, error) {
	var item models.Electronics

	err := config.DB.Find(&item, "nama LIKE ?", keyword).Error
	if err != nil {
		return nil, err
	}
	return item, nil
}
