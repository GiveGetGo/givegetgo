package res

import (
	"github.com/GiveGetGo/shared/types"
	"github.com/gin-gonic/gin"
)

// func ResponseSuccessWithData()
func ResponseSuccessWithData(c *gin.Context, status int, event string, userid uint, response types.Response, data interface{}) {
	c.JSON(status, types.FullResponseWithData{
		FullResponse: types.FullResponse{
			Event: event,
			Code:  response.Code,
			Msg:   response.Msg,
		},
		Data: data,
	})
}

// func ResponseSuccess()
func ResponseSuccess(c *gin.Context, status int, event string, userid uint, response types.Response) {
	c.JSON(status, types.FullResponse{
		Event: event,
		Code:  response.Code,
		Msg:   response.Msg,
	})
}

// func ResponseError()
func ResponseError(c *gin.Context, status int, response types.Response) {
	c.JSON(status, types.Response{
		Code: response.Code,
		Msg:  response.Msg,
	})
}
