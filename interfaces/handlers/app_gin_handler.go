package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"cc/application"
	ccErr "cc/application/error"
	"cc/domain/entity"
	"cc/interfaces/handlers/response"
)

type AppGinHandler struct {
	appInteractor application.AppInteractor
}

// 注入 AppInteractor
func NewAppGinHandler(appInteractor application.AppInteractor) *AppGinHandler {
	return &AppGinHandler{appInteractor}
}

func (handler *AppGinHandler) Add() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code ccErr.ErrorCode
		var app entity.App
		err := c.BindJSON(&app)
		if err != nil {
			code = ccErr.ErrParameterValidate
			response.CCResponse(c, code, err)
			return
		}
		err = handler.appInteractor.CreateApp(app)
		if err != nil {
			code = ccErr.ErrParameterValidate
			response.CCResponse(c, code, err)
			return
		}

		response.Success(c, "success", nil)
	}
}

func (handler *AppGinHandler) FindAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code ccErr.ErrorCode
		apps, err := handler.appInteractor.FindAll()
		if err != nil {
			code = ccErr.ErrParameterValidate
			response.CCResponse(c, code, err)
			return
		}

		response.Success(c, "success", apps)
	}
}

func (handler *AppGinHandler) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code ccErr.ErrorCode
		sid := c.Param("id")
		id, err := strconv.ParseUint(sid, 10, 64)
		if err != nil {
			code = ccErr.ErrParameterValidate
			response.CCResponse(c, code, err)
			return
		}
		app, err := handler.appInteractor.FindOne(id)
		if err != nil {
			code = ccErr.ErrAppNotFound
			response.CCResponse(c, code, err)
			return
		}

		response.Success(c, "success", app)
	}
}

func (handler *AppGinHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code ccErr.ErrorCode
		sid := c.Param("id")
		id, err := strconv.ParseUint(sid, 10, 64)
		if err != nil {
			code = ccErr.ErrParameterValidate
			response.CCResponse(c, code, err)
			return
		}
		err = handler.appInteractor.Delete(id)
		if err != nil {
			code = ccErr.ErrAppDeleted
			response.CCResponse(c, code, err)
			return
		}

		response.Success(c, "success", nil)
	}
}

func (handler *AppGinHandler) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code ccErr.ErrorCode
		sid := c.Param("id")
		id, err := strconv.ParseUint(sid, 10, 64)
		if err != nil {
			code = ccErr.ErrParameterValidate
			response.CCResponse(c, code, err)
			return
		}
		var app entity.App
		err = c.BindJSON(&app)
		if err != nil {
			code = ccErr.ErrParameterValidate
			response.CCResponse(c, code, err)
			return
		}
		app.ID = id
		err = handler.appInteractor.Update(app)
		if err != nil {
			code = ccErr.ErrAppUpdated
			response.CCResponse(c, code, err)
			return
		}

		response.Success(c, "success", nil)
	}
}