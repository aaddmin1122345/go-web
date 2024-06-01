package main

import (
	"fmt"
	"go-web/route"
	//"go-web/utils"
)

func main() {
	fmt.Println("web服务:\thttp://127.0.0.1:8080")
	routes := route.MyRouteImpl{}
	routes.Init()

}
