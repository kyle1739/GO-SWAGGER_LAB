// Package classification ToDo API
//
// 待辦事項DEMO
//
// Terms Of Service:
//
// 使用風險請自行承擔
//
//     Schemes: http
//     Host: localhost:9080
//     BasePath: /api/v1
//     Version: 0.0.1
//
//     Consumes:
//     - application/json
//     - application/xml
//
//     Produces:
//     - application/json
//     - application/xml
//
//
// swagger:meta
package main

import(
	"todoapi/router"
	Gdb "todoapi/databases"
)

func main(){
	defer Gdb.Db.Close()
	router := router.InitRouter()

	router.Run(":9080")
}
