package handler

import (
	"search/model"
	"strconv"
	"github.com/gin-gonic/gin"
)

func Add(c *gin.Context) {
	host:=c.Query("host")
	guest:=c.Query("guest")
	err:=model.Add(host,guest)
	if err!=nil {
		c.JSON(400, model.Response{
			Code:    400,
			Message: "Failed",
			Data:    err,
		})
		return
	} 
	c.JSON(200, model.Response{
		Code:    200,
		Message: "ok",
		Data:    "successful",
	})
}

func Sea(c *gin.Context) {
	host:=c.Query("host")
	guests,err:=model.Sea(host)
	if err!=nil {
		c.JSON(400, model.Response{
			Code:    400,
			Message: "Failed",
			Data:    err,
		})
		return
	}
	c.JSON(200, model.Response{
		Code:    200,
		Message: "ok",
		Data:    guests,
	})
}

func Del(c *gin.Context) {

	value:=c.Query("value")
	method:=c.Query("method")

	if method=="host" {
		err:=model.DelHost(value)
		if err!=nil {
			c.JSON(400, model.Response{
				Code:    400,
				Message: "Failed",
				Data:    err,
			})
			return
		}
	}

	if method=="guest" {

		id,err:=strconv.Atoi(value)

		if err!=nil {
			c.JSON(400, model.Response{
				Code:    400,
				Message: "Failed",
				Data:    err,
			})
			return
		}

		err=model.DelGuest(id)

		if err!=nil {
			c.JSON(400, model.Response{
				Code:    400,
				Message: "Failed",
				Data:    err,
			})
			return
		}
	}
	
	c.JSON(200, model.Response{
		Code:    200,
		Message: "ok",
		Data:    "successful",
	})
}
