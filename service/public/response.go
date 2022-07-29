package public

import (
	"GolangScaffold/global/errCode"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Handler200(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": errCode.SuccessEn, "data": data})
	return
}

func Handler201(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, gin.H{"code": http.StatusCreated, "msg": errCode.CreatedSuccessEn, "data": data})
	return
}

func Handler204(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": http.StatusNoContent, "msg": errCode.SuccessEn, "data": nil})
	return
}

func Handler400(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": msg, "data": data})
	return
}

func Handler404(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "msg": errCode.NotFoundEn, "data": nil})
	return
}

func Handler403(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"code": http.StatusForbidden, "msg": errCode.NoForbiddenEn, "data": nil})
	return
}

func Handler500(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "msg": msg, "data": data})
	return
}