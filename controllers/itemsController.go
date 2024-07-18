package itemsController

import (
	"encoding/json"
	"net/http"

	"github.com/bashori/learn-go/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var items []models.Items
	models.DB.Find(&items)
	c.JSON(http.StatusOK, gin.H{"Items": items})
}

func Create(c *gin.Context) {
	var item models.Items

	if err := c.ShouldBindJSON(&item); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
		return
	}

	models.DB.Create(&item)
	c.JSON(http.StatusOK, gin.H{"Item": item})
}

func Show(c *gin.Context) {
	var item models.Items
	id := c.Param("id")

	if err := models.DB.First(&item, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Message": "Cant find the data"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"Item": item})
}

func Update(c *gin.Context) {
	var item models.Items
	id := c.Param("id")

	if err := c.ShouldBindJSON(&item); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
		return
	}

	if models.DB.Model(&item).Where("id = ?", id).Updates(&item).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Message": "Cant update"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message": "Updated Successfully"})
}

func Delete(c *gin.Context) {
	var item models.Items

	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&item); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
		return
	}

	id, _ := input.Id.Int64()
	if models.DB.Delete(&item, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Message": "Cant delete item"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message": "Deleted Successfully"})
}