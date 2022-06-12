package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ruzwn/restaurantBooking/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		restaurants := api.Group("/restaurants")
		{
			restaurants.POST("/", h.createRestaurant)
			restaurants.GET("/", h.getAllRestaurants)
			restaurants.GET("/:id", h.getRestaurantById)
			restaurants.PUT("/:id", h.updateRestaurant)
			restaurants.DELETE("/:id", h.deleteRestaurant)

			tables := restaurants.Group(":id/tables")
			{
				tables.POST("/", h.createTable)
				tables.GET("/", h.getAllTables)
				tables.GET("/:table_id", h.getTableById)
				tables.PUT("/:table_id", h.updateTable)
				tables.DELETE("/:table_id", h.deleteTable)
			}
		}
	}

	return router
}
