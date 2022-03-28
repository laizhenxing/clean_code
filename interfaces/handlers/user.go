package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"cc/application"
	"cc/domain/entity"
)

type UserHandler struct {
	ui *application.UserInteractor
}

func NewUserHandler(ui *application.UserInteractor) *UserHandler {
	return &UserHandler{ui: ui}
}

func (uh *UserHandler) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user entity.User
		err := c.BindJSON(&user)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"msg": "参数错误"})
			return
		}
		hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"msg": "password hash error", "error": err})
			return
		}
		user.Password = string(hashed)

		err = uh.ui.CreateUser(user)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"msg": "create failed", "error": err})
			return
		}
		c.JSON(http.StatusOK, gin.H{"msg": "create user success"})
	}
}

func (uh *UserHandler) FindOne() gin.HandlerFunc {
	return func(c *gin.Context) {
		sid := c.Param("id")
		id, err := strconv.ParseUint(sid, 10, 64)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"msg": "参数错误"})
			return
		}
		user, err := uh.ui.Get(id)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"msg": "create failed"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"msg": "get user success", "data": user})
	}
}

func (uh *UserHandler) FindAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := uh.ui.GetAll()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"msg": "get users fail"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"msg": "success", "data": users})
	}
}


func (uh *UserHandler) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		sid := c.Param("id")
		id, err := strconv.ParseUint(sid, 10, 64)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"msg": "参数错误"})
			return
		}

		var user entity.User
		err = c.BindJSON(&user)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"msg": "参数错误"})
			return
		}
		user.ID = id
		err = uh.ui.Update(user)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"msg": "update user failed"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"msg": "update user success"})
	}
}

func (uh *UserHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		sid := c.Param("id")
		id, err := strconv.ParseUint(sid, 10, 64)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"msg": "参数错误"})
			return
		}
		err = uh.ui.Delete(id)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"msg": "delete user failed"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"msg": "delete user success"})
	}
}
