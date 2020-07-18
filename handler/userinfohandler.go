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
	_ "github.com/go-sql-driver/mysql"
	log "github.com/jeanphorn/log4go"
	"github.com/labstack/echo"
	"net/http"
	"testproject/config"
	"testproject/model"
)

// ApiGetUserInfo : Get user's information
// URL : /user/1
// Parameters: int id
// Method: GET
// Output: JSON encoded user object, or JSON encoded exception.
func ApiGetUserInfo(c echo.Context) error {
	id := c.Param("id")
	sqlQuery := fmt.Sprintf("SELECT id,name,age,city FROM user_info WHERE id = %s", id)
	db, err := sql.Open("mysql", config.ConnStr)
	if err != nil {
		log.LOGGER("default").Error("InternalServerError")
		return c.JSON(http.StatusInternalServerError, "InternalServerError")
	}
	defer db.Close()
	var uid sql.NullInt32
	var name sql.NullString
	var age sql.NullInt32
	var city sql.NullString
	err = db.QueryRow(sqlQuery).Scan(&uid, &name, &age, &city)
	log.LOGGER("default").Debug("SqlQuery:%s\n", sqlQuery)
	switch {
	case err == sql.ErrNoRows:
		log.LOGGER("default").Error("No user information uid=%s\n", id)
		return c.JSON(http.StatusBadRequest, "no user information uid="+id)
	case err != nil:
		log.LOGGER("default").Error("InternalServerError")
		return c.JSON(http.StatusInternalServerError, "InternalServerError")
	default:
		result := &model.UserInfo{}
		result.Id = uid.Int32
		result.Name = name.String
		result.Age = age.Int32
		result.City = city.String
		log.LOGGER("default").Info("UserInfo:\n", result)
		return c.JSON(http.StatusOK, result)
	}
}

// ApiGetAllUserInfo : Get all users' information
// URL : /users
// Parameters: none
// Method: GET
// Output: JSON encoded users object array, or JSON encoded exception.
func ApiGetAllUserInfo(c echo.Context) error {
	sqlQuery := `SELECT id,name,age,city from user_info`
	db, err := sql.Open("mysql", config.ConnStr)
	defer db.Close()
	if err != nil {
		log.LOGGER("default").Error("InternalServerError")
		return c.JSON(http.StatusInternalServerError, "InternalServerError")
	}
	rows, err := db.Query(sqlQuery)
	if err != nil {
		log.LOGGER("default").Error("InternalServerError")
		return c.JSON(http.StatusInternalServerError, "InternalServerError")
	}
	defer rows.Close()
	log.LOGGER("default").Debug("SqlQuery:%s\n", sqlQuery)
	var totalResult []model.UserInfo
	for rows.Next() {
		result := model.UserInfo{}
		err := rows.Scan(
			&result.Id,
			&result.Name,
			&result.Age,
			&result.City)
		if err != nil {
			log.LOGGER("default").Error("InternalServerError")
			return c.JSON(http.StatusInternalServerError, "InternalServerError")
		}
		totalResult = append(totalResult, result)
	}
	if len(totalResult) == 0 {
		log.LOGGER("default").Error("No user information")
		return c.JSON(http.StatusBadRequest, "no user information")
	}
	log.LOGGER("default").Info("AllUserInfo:\n", totalResult)
	return c.JSON(http.StatusOK, totalResult)
}

// ApiGreatUser - Creat User Information
// URL : /user
// Method: POST
// Body:
/*
 * {
 *	"name":"John",
 *	"age":"34",
 *	"city":"Hope"
 }
*/
// Output: JSON Web Token, or JSON encoded exception.
func ApiCreatUser(c echo.Context) error {
	type Req struct {
		Name string `json:"name"`
		Age  string `json:"age"`
		City string `json:"city"`
	}
	req := &Req{}
	if err := c.Bind(req); err != nil {
		log.LOGGER("default").Error("wrong user information")
		return c.JSON(http.StatusBadRequest, "Creat failed")
	}
	sqlQuery := fmt.Sprintf("INSERT INTO user_info (name, age, city) VALUES ('%s', '%s', '%s')",
		req.Name,
		req.Age,
		req.City)
	db, err := sql.Open("mysql", config.ConnStr)
	defer db.Close()
	log.LOGGER("default").Debug("SqlQuery:%s\n", sqlQuery)
	_, err = db.Exec(sqlQuery)
	if err != nil {
		log.LOGGER("default").Error("InternalServerError")
		return c.JSON(http.StatusInternalServerError, "InternalServerError")
	}
	log.LOGGER("default").Info("GreatUserInfo:\n", req)
	return c.JSON(http.StatusOK, req)
}

// ApiUpdateUser - Update User Information
// URL : /user
// Method: PUT
// Body:
/*
 * {
 *	"id":"7",
 *	"name":"John",
 *	"age":"34",
 *	"city":"Victoria"
 }
*/
// Output: JSON encoded user object, or JSON encoded exception.
func ApiUpdateUser(c echo.Context) error {
	type Req struct {
		Id   string `json:"id"`
		Name string `json:"name"`
		Age  string `json:"age"`
		City string `json:"city"`
	}
	req := &Req{}
	if err := c.Bind(req); err != nil {
		log.LOGGER("default").Error("wrong login information")
		return c.JSON(http.StatusBadRequest, "Login failed")
	}
	sqlQuery := fmt.Sprintf("UPDATE user_info SET name = '%s', age = '%s', city = '%s' where id = %s",
		req.Name,
		req.Age,
		req.City,
		req.Id)
	db, err := sql.Open("mysql", config.ConnStr)
	defer db.Close()
	log.LOGGER("default").Debug("SqlQuery:%s\n", sqlQuery)
	_, err = db.Exec(sqlQuery)
	if err != nil {
		log.LOGGER("default").Error("InternalServerError")
		return c.JSON(http.StatusInternalServerError, "InternalServerError")
	}
	log.LOGGER("default").Info("GreatUserInfo:\n", req)
	return c.JSON(http.StatusOK, req)
}

// ApiDeleteUserInfo : Delete user's information
// URL : /user/7
// Parameters: int id
// Method: DELETE
// Output: JSON encoded user object, or JSON encoded exception.
func ApiDeleteUserInfo(c echo.Context) error {
	id := c.Param("id")
	sqlQuery := fmt.Sprintf("SELECT id,name,age,city FROM user_info WHERE id = %s", id)
	db, err := sql.Open("mysql", config.ConnStr)
	if err != nil {
		log.LOGGER("default").Error("InternalServerError")
		return c.JSON(http.StatusInternalServerError, "InternalServerError")
	}
	defer db.Close()
	var uid sql.NullInt32
	var name sql.NullString
	var age sql.NullInt32
	var city sql.NullString
	err = db.QueryRow(sqlQuery).Scan(&uid, &name, &age, &city)
	log.LOGGER("default").Debug("SqlQuery:%s\n", sqlQuery)
	switch {
	case err == sql.ErrNoRows:
		log.LOGGER("default").Error("No user information uid=%s\n", id)
		return c.JSON(http.StatusBadRequest, "no user information uid="+id)
	case err != nil:
		log.LOGGER("default").Error("InternalServerError")
		return c.JSON(http.StatusInternalServerError, "InternalServerError")
	default:
		sqlQuery := fmt.Sprintf("DELETE FROM user_info WHERE id = %s", id)
		db, err := sql.Open("mysql", config.ConnStr)
		defer db.Close()
		if err != nil {
			log.LOGGER("default").Error("InternalServerError")
			return c.JSON(http.StatusInternalServerError, "InternalServerError")
		}
		log.LOGGER("default").Debug("SqlQuery:%s\n", sqlQuery)
		res, err := db.Exec(sqlQuery)
		if err != nil {
			log.LOGGER("default").Error("InternalServerError")
			return c.JSON(http.StatusInternalServerError, "InternalServerError")
		}
		count, err := res.RowsAffected()
		if err != nil {
			log.LOGGER("default").Error("InternalServerError")
			return c.JSON(http.StatusInternalServerError, "InternalServerError")
		}
		if count > 0 {
			userInfo := model.UserInfo{}
			userInfo.Id = uid.Int32
			userInfo.Name = name.String
			userInfo.Age = age.Int32
			userInfo.City = city.String
			log.LOGGER("default").Info("GreatUserInfo:\n", userInfo)
			return c.JSON(http.StatusOK, userInfo)
		}
		return c.JSON(http.StatusInternalServerError, "InternalServerError")
	}
}
