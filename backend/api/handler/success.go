package handler

import "github.com/gin-gonic/gin"

func successResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(code, data)
}
