package controllers

import "github.com/gin-gonic/gin"

func Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "localhost", false, true)

	c.JSON(200, gin.H{"success": "user logged out"})
}
