package routes
import(
	controller "go-jwt-auth/controllers"
	"go-jwt-auth/middleware"
	"github.com/gin-gonic/gin"
)
func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users",controller.GetUsers())
	incomingRoutes.GET("/users/:user_id",controller.GetUser())
}