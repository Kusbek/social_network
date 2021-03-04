package handler

import (
	"log"

	"github.com/gin-gonic/gin"
)

func errorResponse(c *gin.Context, code int, err error) {
	log.Println(err.Error())
	c.String(code, err.Error())
}
