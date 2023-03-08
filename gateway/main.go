package main

import (
	"gateway/internal/router"
	"log"
)

func main() {
	// 创建路由
	r := router.New()
	err := r.Run()
	if err != nil {
		log.Fatalf("failed to run server %v", err)
	}
}
