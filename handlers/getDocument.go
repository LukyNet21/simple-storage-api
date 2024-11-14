package handlers

import (
	"storage-api/models"
	"storage-api/util"

	"github.com/gin-gonic/gin"
)

func GetDocument(c *gin.Context) {
	documentPath := c.Param("path")
	bucketId := c.Param("bucket_id")
	if documentPath == "" || bucketId == "" {
		c.JSON(400, gin.H{"error": "Document ID and bucket ID are required"})
		return
	}
	var bucket models.Bucket
	err := util.DB.First(&bucket, "id = ?", bucketId).Error
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	var document models.Document
	err = util.DB.First(&document, "path = ? AND bucket_id = ?", documentPath, bucket.ID).Error
	if err != nil {
		c.JSON(400, gin.H{"error": "Document not found"})
		return
	}

	c.FileAttachment("data/buckets/"+bucket.Path+"/"+document.Path, document.Filename)
}
