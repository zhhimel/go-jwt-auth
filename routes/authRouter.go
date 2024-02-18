package routes
import(
	controller "go-jwt-auth/controllers"
	"github.com/gin-gonic/gin"
)
func AuthRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.POST("users/signup",controller.Signup())
	incomingRoutes.POST("users/login",controller.Login())
}