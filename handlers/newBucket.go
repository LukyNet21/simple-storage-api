package handlers

import (
	"storage-api/models"
	"storage-api/util"

	"github.com/gin-gonic/gin"
)

func NewBucket(c *gin.Context) {
	var bucket models.Bucket
	c.BindJSON(&bucket)

	if bucket.Name == "" {
		c.JSON(400, gin.H{"error": "Bucket name is required"})
		return
	}

	for {
		bucket.Path = util.RandString(16)
		var count int64
		util.DB.Model(&models.Bucket{}).Where("path = ?", bucket.Path).Count(&count)
		if count == 0 {
			break
		}
	}

	err := util.DB.Create(&bucket).Error
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, bucket)
}
