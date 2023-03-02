package main

import (
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	go service.RunMessageServer()

	r := setupRouter()
	r.Run(":8080")
}