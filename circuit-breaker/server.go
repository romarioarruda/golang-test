package server

import(
	"fmt"
	"github.com/gin-gonic/gin"
)

var startTime time.Time = time.Now()

func server() {
	e := gin.Default()
	e.GET("/ping", func(ctx *gin.Context) {
		if time.Since(startTime) < 5*time.Second {
			ctx.String(http.StatusInternalServerError, "pong")
			return
		}
		ctx.String(http.StatusOK, "pong")
	})

	fmt.Printf("Starting server at port 8080\n")
	e.Run(":8080")
}