package handlers

import (
	"storage-api/models"
	"storage-api/util"

	"github.com/gin-gonic/gin"
)

func ListBuckets(c *gin.Context) {
	var buckets []models.Bucket
	util.DB.Find(&buckets)

	c.JSON(200, buckets)
}
