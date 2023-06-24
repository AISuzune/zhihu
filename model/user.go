package model

import (
	"github.com/dgrijalva/jwt-go"
)

type User struct {
	ID       int    `json:"id"`
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
