package controllers

import (
	"context"
	"fmt"
	"go-jwt-auth/helpers"
	helper "go-jwt-auth/helpers"
	"go-jwt-auth/models"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go.playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)
var userCollection *mongo.Collection= database.OpenCollection(database.Client,"user")
var validate=validator.New()
func HashPassword()
func VerifyPassword()
func Signup()
func Login()
func GetUser()
func GetUser() gin.HandlerFunc{
	return func (c *gin.Context){
		userId:=c.Param("user_id")
		
		if err:=helper.MatchUserTypeToUid(c, userId);err!= nil {
			c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
			return 
		}
		var ctx,cancel=context.WithTimeout(context.Background(),100*time.Second);
	}
}