package router

import (
	"finalProject/controllers"
	"finalProject/middleware"
	"finalProject/repository"
	"finalProject/service"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	photoRepository := repository.NewPhotoRepository()
	photoService := service.NewPhotoService(photoRepository)
	photoController := controllers.NewPhotoController(photoService)

	commentRepository := repository.NewCommentRepository()
	commentService := service.NewCommentService(commentRepository)
	commentController := controllers.NewCommentController(commentService)

	socialMediaRepository := repository.NewSocialMediaRepository()
	socialMediaService := service.NewSocialMediaService(socialMediaRepository)
	socialMediaController := controllers.NewSocialMediaController(socialMediaService)

	userRouter := router.Group("/users")
	{
		userRouter.POST("/register", userController.UserRegister)
		userRouter.POST("/login", userController.UserLogin)

		userRouter.Use(middleware.Authentication())
		userRouter.PUT("/:userId", middleware.UserAuthorization(), userController.UpdateUser)
		userRouter.DELETE("/:userId", middleware.UserAuthorization(), userController.DeleteUser)
	}

	photoRouter := router.Group("/photos")
	{
		photoRouter.Use(middleware.Authentication())
		photoRouter.GET("/", photoController.GetAllPhotos)
		photoRouter.POST("/", photoController.CreatePhoto)
		photoRouter.DELETE("/:photoId", middleware.PhotoAuthorization(), photoController.DeletePhoto)
		photoRouter.PUT("/:photoId", middleware.PhotoAuthorization(), photoController.UpdatePhoto)
	}

	commentRouter := router.Group("/comments")
	{
		commentRouter.Use(middleware.Authentication())
		commentRouter.GET("/", commentController.GetAllComments)
		commentRouter.POST("/", commentController.CreateComment)
		commentRouter.DELETE("/:commentId", middleware.CommentAuthorization(), commentController.DeleteComment)
		commentRouter.PUT("/:commentId", middleware.CommentAuthorization(), commentController.UpdateComment)
	}

	socialMediaRouter := router.Group("/socialmedias")
	{
		socialMediaRouter.Use(middleware.Authentication())
		socialMediaRouter.GET("/", socialMediaController.GetAllSocialMedias)
		socialMediaRouter.POST("/", socialMediaController.CreateSocialMedia)
		socialMediaRouter.DELETE("/:socialMediaId", middleware.SocialMediaAuthorization(), socialMediaController.DeleteSocialMedia)
		socialMediaRouter.PUT("/:socialMediaId", middleware.SocialMediaAuthorization(), socialMediaController.UpdateSocialMedia)
	}

	return router
}
