package main

import (
	"storage-api/handlers"
	"storage-api/util"

	"github.com/gin-gonic/gin"
)

func main() {
	util.InitDb()

	r := gin.Default()
	v1 := r.Group("/api/v1")

	buckets := v1.Group("/buckets")
	buckets.GET("", handlers.ListBuckets)
	buckets.POST("", handlers.NewBucket)

	documents := v1.Group("/documents/:bucket_id")
	documents.POST("/new", handlers.NewDocument)
	documents.GET("/:path", handlers.GetDocument)
	documents.DELETE("/:path", handlers.DeleteDocument)

	r.Run(":8080")
}
