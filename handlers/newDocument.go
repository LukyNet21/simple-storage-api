package handlers

import (
	"storage-api/models"
	"storage-api/util"

	"github.com/gin-gonic/gin"
)

func NewDocument(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{"error": "No file uploaded"})
		return
	}

	if file == nil {
		c.JSON(400, gin.H{"error": "No file uploaded"})
		return
	}

	document := models.Document{
		Path:     util.RandString(60),
		Filename: file.Filename,
	}

	bucketId := c.Param("bucket_id")
	var bucket models.Bucket
	util.DB.First(&bucket, "id = ?", bucketId)
	if bucket.ID == 0 {
		c.JSON(400, gin.H{"error": "Bucket not found"})
		return
	}

	for {
		document.Path = util.RandString(60)
		var count int64
		util.DB.Model(&models.Document{}).Where("path = ? AND bucket_id = ?", document.Path, bucket.ID).Count(&count)
		if count == 0 {
			break
		}
	}

	document.BucketID = bucket.ID
	err = util.DB.Create(&document).Error
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.SaveUploadedFile(file, "data/buckets/"+bucket.Path+"/"+document.Path)

	c.JSON(200, gin.H{"message": "File uploaded successfully"})
}
