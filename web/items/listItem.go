package items

import (
	"TGU-MAP/models"
	"TGU-MAP/service"
	"TGU-MAP/utils"
	"context"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"strconv"
)

var ctx = context.Background()

func HandleInsertData(c *gin.Context) {
	var data []models.ListItem
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := service.ListItemClient.InsertData(&data); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	service.RDB.Del(ctx, "list_items")
	c.JSON(200, gin.H{"status": "data inserted successfully"})
}

func HandleGetData(c *gin.Context) {
	// 尝试从缓存中获取数据
	cachedData, err := service.RDB.Get(ctx, "list_items").Result()
	if errors.Is(err, redis.Nil) {
		// 缓存未命中，从数据库中获取数据
		data, err := service.ListItemClient.FetchData()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		// 缓存数据
		service.RDB.Set(ctx, "list_items", string(utils.Marshal(*data)), 0)

		c.JSON(200, data)
	} else if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	} else {
		// 返回缓存的数据
		var data []models.ListItem
		if err := json.Unmarshal([]byte(cachedData), &data); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, data)
	}
}

func HandleUpdateNode(c *gin.Context) {
	var item models.ListItem
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	elemID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid ID"})
		return
	}

	if err := service.ListItemClient.UpdateNodeByID(&item, uint(elemID)); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// 清除缓存
	service.RDB.Del(ctx, "list_items")

	c.JSON(200, gin.H{"status": "node updated successfully"})
}

func HandleDeleteNode(c *gin.Context) {
	id := c.Param("id")
	elemID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid ID"})
		return
	}

	if err := service.ListItemClient.DeleteNodeByID(uint(elemID)); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// 清除缓存
	service.RDB.Del(ctx, "list_items")

	c.JSON(200, gin.H{"status": "node deleted successfully"})
}
