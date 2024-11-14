package handlers

import (
	"os"

	"storage-api/models"
	"storage-api/util"

	"github.com/gin-gonic/gin"
)

func DeleteDocument(c *gin.Context) {
	documentPath := c.Param("path")
	if documentPath == "" {
		c.JSON(400, gin.H{"error": "Document path is required"})
		return
	}

	var document models.Document
	err := util.DB.First(&document, "path = ?", documentPath).Error
	if err != nil {
		c.JSON(400, gin.H{"error": "Document not found"})
		return
	}

	err = util.DB.Delete(&document).Error
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var bucket models.Bucket
	err = util.DB.First(&bucket, "id = ?", document.BucketID).Error
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	err = os.Remove("data/buckets/" + bucket.Path + "/" + document.Path)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Document deleted"})
}
