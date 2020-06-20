package customer

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nandanai/finalexam/database"
)

func updateCustomersHandler(c *gin.Context) {
	id := c.Param("id")
	stmt, err := database.Conn().Prepare("SELECT id, name, email, status FROM customers where id=$1")
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	row := stmt.QueryRow(id)

	customer := Customer{}

	err = row.Scan(&customer.ID, &customer.Name, &customer.Email, &customer.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	stmt, err = database.Conn().Prepare("UPDATE customers SET name=$2, email=$3, status=$4 WHERE id=$1;")
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if _, err := stmt.Exec(id, customer.Name, customer.Email, customer.Status); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, customer)
}
