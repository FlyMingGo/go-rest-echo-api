/**
 * @Author: Frank
 * @Description: main
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2020/7/7
 */
package main

import (
	log "github.com/jeanphorn/log4go"
	"github.com/labstack/echo"
	stderr "log"
	"testproject/config"
	"testproject/router"
)

func main() {
	log.LoadConfiguration("./log4go.json")
	e := echo.New()
	router.InitRouter(e)                 //register requests, address: /login &/users & /user/:id and bind controllers.
	err := e.Start(":" + config.ApiPort) //start http server, listening 80 port.
	if err != nil {
		stderr.Fatal(err.Error())
	}
}
