package driver

import (
	"errors"
	"onlinemarketplace/model"
	"strconv"
	"strings"
)

func CreateUser(user *model.User) error {
	if err := GetDBConn().Where("username != ?", user.UserName).Create(user).Error; err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return errors.New("User already exist")
		}
		return errors.New("Some technical issue")
	}
	return nil
}

func CountUser(user *model.User) (int64, error) {
	var count int64
	if err := GetDBConn().Find(user).Count(&count).Error; err != nil {
		return -1, errors.New("Some technical issue")
	}
	return count, nil
}

func UpdateUserToken(user model.User) error {
	if err := GetDBConn().Model(&user).Where("username = ?", user.UserName).Updates(map[string]interface{}{"token": user.Token}).Error; err != nil {
		return errors.New("Some technical issue")
	}
	return nil
}

func CreateProduct(prod *model.Product) error {
	switch {
	case prod.Name == "":
		return errors.New("Name shouldn't be empty")
	case prod.Description == "":
		return errors.New("Description shouldn't be empty")
	case prod.Price <= 0:
		return errors.New("Price shouldn't be zero or negative")
	case len(prod.Name) > 100:
		return errors.New("Name shouldn't be greater than 100 character")
	case len(prod.Description) > 200:
		return errors.New("Description shouldn't be greater than 200 character")
	}
	err := GetDBConn().Model(&prod).Create(prod).Error
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return errors.New("Product already exist")
		}
		return errors.New("Some technical issue")
	}
	return nil
}

func GetProduct(prod *[]model.Product) error {
	if err := GetDBConn().Model(&prod).Find(prod).Error; err != nil {
		return errors.New("Some technical issue")
	}
	return nil
}

func UpdateProductById(idstr string, prod *model.Product) error {
	if prod.Name == "" || prod.Price <= 0 || prod.Description == "" {
		return errors.New("Invalid input data")
	}
	id, err := strconv.Atoi(idstr)
	if err != nil {
		return errors.New("Invalid product Id")
	}
	if err := GetDBConn().Model(&prod).Where("id = ?", id).Updates(map[string]interface{}{"Name": prod.Name, "Description": prod.Description, "Price": prod.Price}).Error; err != nil {
		return errors.New("Some technical issue")
	}
	return nil
}

func DeleteProductById(id int, prod *model.Product) error {
	if err := GetDBConn().Where("id = ?", id).Model(&prod).Delete(&prod).Error; err != nil {
		return errors.New("Some technical issue")
	}
	return nil
}

func GetProductById(id int, prod *model.Product) (bool, error) {
	if err := GetDBConn().Model(&prod).Where("id = ?", id).Find(&prod).Error; err != nil {
		if err.Error() == "record not found" {
			return true, errors.New("Resource not found")
		}
		return false, errors.New("Some technical issue")
	}
	return false, nil
}
