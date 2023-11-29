package controller

import (
	"errors"
	"net/http"
	"onlinemarketplace/driver"
	"onlinemarketplace/helper"
	"onlinemarketplace/model"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func Signup(c *gin.Context) {
	res := helper.NewResponse()
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		log.Error("Payload Error :: " + err.Error())
		c.JSON(http.StatusBadRequest, res.ErrorResponse(errors.New("Invalid input data")))
		return
	}
	user.Password, err = helper.Encryption(user.Password, os.Getenv("SecretKey"))
	if err != nil {
		log.Error("Encrytion Error :: " + err.Error())
		c.JSON(http.StatusInternalServerError, res.ErrorResponse(errors.New("Can't Decryt the string")))
		return
	}

	if err := driver.CreateUser(&user); err != nil {
		log.Error("DB Issue :: ", err)
		c.JSON(http.StatusInternalServerError, res.ErrorResponse(err))
		return
	}
	helper.LogInfo(user, res.Response(struct{}{}))
	c.JSON(http.StatusOK, res.Response(struct{}{}))
}

func Login(c *gin.Context) {
	res := helper.NewResponse()
	user := model.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Error("Payload error :: " + err.Error())
		c.JSON(http.StatusBadRequest, res.ErrorResponse(errors.New("Invalid input data")))
		return
	}
	pass, err := helper.Encryption(user.Password, os.Getenv("SecretKey"))
	if err != nil {
		log.Error("Encrytion error :: " + err.Error())
		c.JSON(http.StatusInternalServerError, res.ErrorResponse(errors.New("Can't Decryt the string")))
		return
	}
	user.Password = pass
	count, err := driver.CountUser(&user)
	if err != nil {
		log.Error("DB issue :: ", err)
		c.JSON(http.StatusInternalServerError, res.ErrorResponse(err))
		return
	}
	if count == 0 {
		c.JSON(http.StatusInternalServerError, res.ErrorResponse(errors.New("User not found")))
		return
	}

	user.Token, err = helper.GenerateToken(&user)
	if err != nil {
		log.Error("Token generation issue :: ", err)
		c.JSON(http.StatusInternalServerError, res.ErrorResponse(err))
	}
	if err := driver.UpdateUserToken(user); err != nil {
		log.Error("DB issue :: ", err)
		c.JSON(http.StatusInternalServerError, res.ErrorResponse(err))
		return
	}

	helper.LogInfo(user, res.Response(struct{}{}))
	c.JSON(http.StatusOK, res.Response(map[string]string{"token": user.Token}))
}

func GetProduct(c *gin.Context) {
	res := helper.NewResponse()
	var prod []model.Product

	if err := driver.GetProduct(&prod); err != nil {
		log.Error("DB issue", err)
		c.JSON(http.StatusInternalServerError, res.ErrorResponse(err))
		return
	}

	helper.LogInfo(prod, res.Response(prod))
	c.JSON(http.StatusOK, res.Response(prod))
}

func CreateProduct(c *gin.Context) {
	res := helper.NewResponse()
	var prod model.Product
	err := c.ShouldBindJSON(&prod)
	if err != nil {
		log.Error("Payload error :: " + err.Error())
		c.JSON(http.StatusBadRequest, res.ErrorResponse(errors.New("Invalid input data")))
		return
	}

	if err := driver.CreateProduct(&prod); err != nil {
		log.Error("DB issue ::", err)
		c.JSON(http.StatusInternalServerError, res.ErrorResponse(err))
		return
	}

	helper.LogInfo(prod, res.Response(struct{}{}))
	c.JSON(http.StatusCreated, res.Response(struct{}{}))
}

func GetProductById(c *gin.Context) {
	res := helper.NewResponse()
	var prod = model.Product{}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, res.ErrorResponse(errors.New("Invalid input data")))
		return
	}
	Isduplicate, err := driver.GetProductById(id, &prod)
	if Isduplicate {
		log.Error("Duplicate issue :: ", err)
		c.JSON(http.StatusNotFound, res.ErrorResponse(err))
		return
	}
	if err != nil {
		log.Error("DB issue :: ", err)
		c.JSON(http.StatusInternalServerError, res.ErrorResponse(err))
		return
	}

	helper.LogInfo(prod, res.Response(prod))
	c.JSON(http.StatusOK, res.Response(prod))
}

func UpdateProduct(c *gin.Context) {
	res := helper.NewResponse()
	prod := model.Product{}
	if err := c.ShouldBindJSON(&prod); err != nil {
		log.Error("Payload error :: " + err.Error())
		c.JSON(http.StatusBadRequest, res.ErrorResponse(errors.New("Invalid input data")))
		return
	}

	if err := driver.UpdateProductById(c.Param("id"), &prod); err != nil {
		log.Error("DB issue :: ", err)
		c.JSON(http.StatusInternalServerError, res.ErrorResponse(errors.New("Some technical error")))
		return
	}

	helper.LogInfo(prod, res.Response(struct{}{}))
	c.JSON(http.StatusOK, res.Response(struct{}{}))
}

func DeleteProduct(c *gin.Context) {
	res := helper.NewResponse()
	var prod = model.Product{}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Error("Invalid product Id :: ", c.Param("id"))
		c.JSON(http.StatusBadRequest, res.ErrorResponse(errors.New("Invalid product Id")))
		return
	}

	if err := driver.DeleteProductById(id, &prod); err != nil {
		log.Error("DB issue ::", err)
		c.JSON(http.StatusInternalServerError, res.ErrorResponse(errors.New("Some technical error")))
		return
	}

	helper.LogInfo(prod, res.Response(struct{}{}))
	c.JSON(http.StatusOK, res.Response(struct{}{}))
}
