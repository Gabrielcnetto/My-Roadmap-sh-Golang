package main

import (
	"fmt"
	"net/http"
	"tidy/controllers"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Role string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateJWT(userId string) (string, error) {
	NewClaim := Claims{Role: "Admin", RegisteredClaims: jwt.RegisteredClaims{
		Subject:   userId,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, NewClaim)
	return token.SignedString([]byte("skkkkkkkkkkkkkkkkk"))
}

func VerifyToken(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (any, error) {
		return []byte("skkkkkkkkkkkkkkkkk"), nil
	})
	if err != nil {
		return "", err
	} else if claims, ok := token.Claims.(*Claims); ok {
		return claims.Role, nil
	} else {
		return "", fmt.Errorf("Claims error")
	}
}

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/", func(c *gin.Context) {
		UserId := c.GetHeader("userId")
		if UserId == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "userId not found",
			})
			c.Abort()
			return
		}
		jwt, err := GenerateJWT(UserId)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			c.Abort()
			return
		}
		c.Set("Token", jwt)
	}, controllers.Hellow)
	router.GET("/using", func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "token not found",
			})
			c.Abort()
			return
		}
		role, err := VerifyToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			c.Abort()
			return
		}
		c.Set("role", role)
	}, controllers.Hellow2)
	router.Run(":8080")
}
