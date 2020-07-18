/**
 * @Author: Frank
 * @Description: Endpoint to handle requests
 * @File:  handler
 * @Version: 1.0.0
 * @Date: 2020/7/7
 */
package handler

import (
	"database/sql"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/jeanphorn/log4go"
	"github.com/labstack/echo"
	"net/http"
	"testproject/config"
	"time"
)

// ApiLogin - User login and get JWT
// URL : /login
// Method: POST
// Body:
/*
 * {
 *	"username":"admin",
 *	"password":"admin"
 }
*/
// Output: JSON Web Token, or JSON encoded exception.
func ApiLogin(c echo.Context) error {
	type Req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	req := &Req{}
	if err := c.Bind(req); err != nil {
		log.LOGGER("default").Error("wrong login information")
		return c.JSON(http.StatusBadRequest, "Login failed")
	}
	sqlQuery := fmt.Sprintf("SELECT id FROM user_login WHERE username = '%s' AND password = '%s'",
		req.Username,
		req.Password)
	db, err := sql.Open("mysql", config.ConnStr)
	defer db.Close()
	log.LOGGER("default").Debug("SqlQuery:%s\n", sqlQuery)
	var uid int
	err = db.QueryRow(sqlQuery).Scan(&uid)
	switch {
	case err == sql.ErrNoRows:
		log.LOGGER("default").Error("wrong login information user=%s\n", req.Username)
		return c.JSON(http.StatusBadRequest, "Login failed")
	case err != nil:
		log.LOGGER("default").Error("InternalServerError")
		return c.JSON(http.StatusInternalServerError, "InternalServerError")
	default:
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["uid"] = uid
		claims["exp"] = time.Now().Add(time.Second * 3600).Unix()
		tokenString, err := token.SignedString([]byte(config.JwtSecretKey))
		if err != nil {
			log.LOGGER("default").Error("InternalServerError")
			return c.JSON(http.StatusInternalServerError, "InternalServerError")
		}
		log.LOGGER("default").Info("JWT=%s user=%s\n", tokenString, req.Username)
		return c.JSON(http.StatusCreated, tokenString)
	}
}
