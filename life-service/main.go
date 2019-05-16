package main

import (
	"life-service/controller/life"
	"life-service/controller/user"
	"life-service/violet"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// http.HandleFunc("/home", controller.Home)
	// fmt.Println(http.ListenAndServe(":11223",nil))

	v := violet.Default()

	v.POST("/user/login", user.Login)
	v.POST("/user/register", user.Register)
	v.GET("/user/logout", user.Logout)

	v.POST("/life/save", life.Save)
	v.GET("/life/all/:page", life.All)
	v.POST("/life/delete/:id", life.Delete)

	v.MakeStatic("/static", "static")

	v.Start(":8089")
}
