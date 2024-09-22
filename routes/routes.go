package routes

import (
	"github.com/gin-gonic/gin"
	"sumit.com/mise-link/middleware"
)

func RegisterRoutes(server *gin.Engine) {
	server.POST("/api/user/signup", createUser)
	server.POST("/api/user/login", login)

	authenticated := server.Group("/")
	authenticated.Use(middleware.IsLoggedIn)
	authenticated.POST("/api/outlet/create", createOutlet)
	authenticated.GET("/api/outlet/getOutletByUserId", getOutletsByUserId)
	authenticated.POST("/api/staff/create", createStaff)      // give the quey as outletId={outletId}
	authenticated.POST("/api/inventory/create", addInventory) // give the quey as outletId={outletId}
	authenticated.POST("/api/product/create", addProduct)     // give the quey as outletId={inventoryId}

}
