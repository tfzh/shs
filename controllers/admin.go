package controllers

import (
    "net/http"
    //"log"
    "github.com/gin-gonic/gin"
    . "../models"
)

func Test(c *gin.Context) {
    // var msg struct {
    //         Name    string `json:"user" xml:"user"`
    //         Message string
    //         Number  int
    //     }
        //num := 333
        num := GetAllAdmin()
        // msg.Name = "Lena"
        // msg.Message = "hey"
        // msg.Number = num
        c.JSON(http.StatusOK, num)
}