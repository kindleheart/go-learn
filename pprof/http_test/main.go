package main

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"math"
	"math/rand"
	"net/http"
)

type IndexReq struct {
	Count int `form:"count"`
}

func main() {
	r := gin.New()
	r.GET("/index", func(c *gin.Context) {
		var indexReq IndexReq
		err := c.ShouldBindQuery(&indexReq)
		if err != nil {
			c.String(http.StatusBadRequest, "参数有误")
			return
		}
		arr := make([]int, indexReq.Count)
		for i := 0; i < indexReq.Count; i++ {
			arr[i] = rand.Intn(math.MaxInt32)
		}
		c.JSON(
			http.StatusOK,
			gin.H{
				"arr":   arr,
				"count": indexReq.Count,
			},
		)
	})
	pprof.Register(r)
	r.Run()
}
