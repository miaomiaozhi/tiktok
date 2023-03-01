package main

import (
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
	ds "feed/ds"
	"time"
)

func setupRouter() *gin.Engine {
	// r.GET("/someProtoBuf", func(c *gin.Context) {
	// 	reps := []int64{int64(1), int64(2)}
	// 	label := "test"
	// 	// protobuf 的具体定义写在 testdata/protoexample 文件中。
	// 	data := &protoexample.Test{
	// 		Label: &label,
	// 		Reps:  reps,
	// 	}
	// 	// 请注意，数据在响应中变为二进制数据
	// 	// 将输出被 protoexample.Test protobuf 序列化了的数据
	// 	c.ProtoBuf(http.StatusOK, data)
	// })

	r := gin.Default()
	r.GET("/feed/", func(c *gin.Context) {
		data := &ds.DouyinFeedResponse {
			StatusCode: 0,
			StatusMsg: "None",
			VideoList: nil,
			NextTime: time.Now().Unix(),
		}
		dataFlect := data.ProtoReflect()
		c.ProtoBuf(http.StatusOK, dataFlect)
	})
	return r
}

func main() {
	go service.RunMessageServer()

	r := setupRouter()
	r.Run(":8080")
}