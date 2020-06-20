package customer

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nandanai/finalexam/database"
)

func deleteCustomersHandler(c *gin.Context) {
	id := c.Param("id")

	stmt, err := database.Conn().Prepare("DELETE FROM customers WHERE id = $1")
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if _, err := stmt.Exec(id); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "customer deleted"})
}
