package users

func UserList(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Users",
	})
}
