package middleware

import (
	"first_task/exception"
	"first_task/helper"
	"time"

	"net/http"
)

func AuthAdmin(request *http.Request) {
	cookie, err := request.Cookie("token")
	if err != nil {
		panic(exception.NewUnauthorizedError(err.Error()))
	}

	tokenString := cookie.Value

	claims, err := helper.ParseToken(tokenString)
	if err != nil {
		panic(exception.NewUnauthorizedError(err.Error()))
	}
	claims.GetExpirationTime()

	
	if claims["role"] != "Admin" {
		panic(exception.NewUnauthorizedError(""))
	}
 
	if claims["exp"].(int64) < time.Now().Unix(){
		panic(exception.NewUnauthorizedError("token expired"))
	}
}
