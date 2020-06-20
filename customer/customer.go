package customer

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func MainAPI() *gin.Engine {
	r := gin.Default()

	r.Use(authMiddleware)

	r.POST("/customers", createCustomersHandler)
	r.GET("/customers/:id", getCustomerByIDHandler)
	r.GET("/customers", getCustomersHandler)
	r.PUT("/customers/:id", updateCustomersHandler)
	r.DELETE("/customers/:id", deleteCustomersHandler)
	return r
}
