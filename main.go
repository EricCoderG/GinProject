package main

import (
	"GinProject/common"
	"github.com/gin-gonic/gin"
)

func main() {

	common.InitDB()
	r := gin.Default()
	r = CollectRoute(r)
	err := r.Run()
	if err != nil {
		panic(err.Error()) // listen and serve on
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
