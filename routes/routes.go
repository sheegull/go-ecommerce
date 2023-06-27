package routes

import (
	"log"
	"os"

	"github.com/akhil/ecommerce-yt/controllers"
	"github.com/akhil/ecommerce-yt/database"
	"github.com/akhil/ecommerce-yt/middleware"
	"github.com/akhil/ecommerce-yt/routes"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
  incomingRoutes.POST("/users/singup", controllers.SignUp())
  incomingRoutes.POST("/users/login", controllers.Login())
  incomingRoutes.POST("/admin/addproduct", controllers.ProductViewerAdmin())
  incomingRoutes.GET("/users/productview", controllers.SearchProduct())
  incomingRoutes.GET("/users/search", controllers.SearchProductByQuery())
}

func main() {
  port := os.Getenv("PORT")
  if port == "" {
    port = "8000"
  }

  app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))

  router := gin.New()
  router.Use(gin.Logger())

  routes.UserRoutes(router)
  router.Use(middleware.Authentication())

  router.GET("/addtocart", app.AddToCart())
  router.GET("removeitem", app.RemoveItem())
  router.GET("/cartcheckout", app.BuyFromCart())
  router.GET("/instantbuy", app.InstantBuy())

  log.Fatal(router.Run(":" + port))
}
