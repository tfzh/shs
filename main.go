package main

import (
     router "./router"
)
type Admin struct {
 Id     int    `json:"id" form:"id"`
 Name   string `json:"name" form:"name"`
 Mail   string `json:"mail" form:"mail"`
}
func main() {
    router :=router.InitRouter()
    router.Run(":8090")
}