package response

import (
	"net/http"

	"github.com/gin-gonic/gin"

	ccErr "cc/application/error"
)

func CCResponse(c *gin.Context, code ccErr.ErrorCode, err error) {
	e := ccErr.NewError(code, err)
	c.JSON(http.StatusOK, e.ToMap())
}

func Success(c *gin.Context, msg string, data interface{}) {
	if data == nil {
		data = struct{}{}
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": msg,
		"data":    data,
	})
}

func Resp(c *gin.Context, code ccErr.ErrorCode, msg string, data interface{}) {
	if data == nil {
		data = struct{}{}
	}
	if len(msg) == 0 {
		msg = ccErr.GetErrorMsg(code)
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": msg,
		"data":    data,
	})
}
