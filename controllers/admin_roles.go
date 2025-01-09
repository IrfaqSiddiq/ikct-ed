package controllers

import (
	"ikct-ed/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllRoles(c *gin.Context){
	roles,err :=models.GetAllRoles()

	if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failed",
            "message": "failed to get roles",
			"error": err.Error(),
		})
        return
    }

	c.JSON(http.StatusOK,gin.H{
		"status": "ok",
		"roles": roles,
	})

}