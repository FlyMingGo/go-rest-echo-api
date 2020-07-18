/**
 * @Author: Frank
 * @Description: router
 * @File:  router
 * @Version: 1.0.0
 * @Date: 2020/7/7
 */
package router

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"testproject/config"
	"testproject/handler"
)

func InitRouter(e *echo.Echo) {
	e.POST("/login", handler.ApiLogin)                                                      // POST /login login and get a JWT.
	e.GET("/user/:id", handler.ApiGetUserInfo, middleware.JWT([]byte(config.JwtSecretKey))) // GET /user/1 returns the user's information with id 1.
	e.GET("/users", handler.ApiGetAllUserInfo, middleware.JWT([]byte(config.JwtSecretKey))) // GET /users returns all users' information.
	e.POST("/user", handler.ApiCreatUser, middleware.JWT([]byte(config.JwtSecretKey))) //POST /user creat new user information
	e.PUT("/user", handler.ApiUpdateUser, middleware.JWT([]byte(config.JwtSecretKey))) //PUT /user update user's information
	e.DELETE("/user/:id", handler.ApiDeleteUserInfo, middleware.JWT([]byte(config.JwtSecretKey))) //DELETE /user delete user's information
}
