package routes

import (
	"api-perpus-go/internal/handlers"
	"api-perpus-go/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterAPIRoutes(r *gin.Engine) {
	api := r.Group("/api")
	api.POST("/login", handlers.Login)
	api.Use(middleware.JWTAuthMiddleware())


	api.POST("/categories", handlers.CreateCategory)
	api.GET("/categories", handlers.ListCategories)
	api.PUT("/categories/:id", handlers.UpdateCategory)
	api.DELETE("/categories/:id", handlers.DeleteCategory)


	api.POST("/books", handlers.CreateBook)
	api.GET("/books", handlers.ListBooks)
	api.GET("/books/:id", handlers.GetBook)
	api.PUT("/books/:id", handlers.UpdateBook)
	api.DELETE("/books/:id", handlers.DeleteBook)


	api.POST("/members", handlers.CreateMember)
	api.GET("/members", handlers.ListMembers)
	api.GET("/members/:id", handlers.GetMember)
	api.PUT("/members/:id", handlers.UpdateMember)
	api.DELETE("/members/:id", handlers.DeleteMember)


	api.GET("/loans", handlers.GetLoans)
	api.GET("/loans/:id", handlers.GetLoanDetail)
	api.POST("/loans", handlers.CreateLoan)
	api.PUT("/loans/:id/return", handlers.ReturnLoan)
}
