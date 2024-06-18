package listItem

import (
	"TGU-MAP/models"
	"TGU-MAP/service"
	"TGU-MAP/utils"
	"TGU-MAP/web/handler"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"net/http"
	"strconv"
)

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
	service.RDB.Del(handler.Ctx, "list_items")
	c.JSON(200, gin.H{"status": "data inserted successfully"})
}

func GetData() (string, error) {
	// 尝试从缓存中获取数据
	//TODO: 可能需要上写锁
	cachedData, err := service.RDB.Get(handler.Ctx, "list_items").Result()
	if errors.Is(err, redis.Nil) {
		// 缓存未命中，从数据库中获取数据
		data, err := service.ListItemClient.FetchData()
		if err != nil {
			return "", err
		}
		cache := string(utils.Marshal(*data))
		// 缓存数据
		service.RDB.Set(handler.Ctx, "list_items", cache, 0)

		return cache, nil
	} else if err != nil {
		return "", err
	} else {
		return cachedData, nil
	}
}

func HandleGetData(c *gin.Context) {
	if data, err := GetData(); err != nil {
		c.JSON(500, handler.StatusBad("出错了……", nil))
	} else {
		c.JSON(200, handler.StatusOK("查询成功", data))
	}

}

func HandleInsertNode(c *gin.Context) {
	var item models.ListItem
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, handler.StatusBad("参数错误", nil))
		return
	}
	id := c.Param("id")
	tmp, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(400, handler.StatusBad("参数错误", nil))
		return
	}
	elemID := uint(tmp)
	var pElemID *uint
	if elemID == 0 {
		pElemID = nil
	} else {
		pElemID = &elemID
	}
	if err := service.ListItemClient.InsertNodeByID(&item, pElemID); err != nil {
		c.JSON(500, handler.StatusBad("出错了……", nil))
		return
	}

	// 清除缓存
	service.RDB.Del(handler.Ctx, "list_items")
	data, _ := GetData()
	c.JSON(200, handler.StatusOK("添加成功", data))

}

func HandleUpdateNode(c *gin.Context) {
	var item models.ListItem
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(400, handler.StatusBad("参数错误", nil))
		return
	}

	id := c.Param("id")
	elemID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(400, handler.StatusBad("参数错误", nil))
		return
	}

	if err := service.ListItemClient.UpdateNodeByID(&item, uint(elemID)); err != nil {
		c.JSON(500, handler.StatusBad("出错了……", nil))
		return
	}

	// 清除缓存
	service.RDB.Del(handler.Ctx, "list_items")
	data, _ := GetData()
	c.JSON(200, handler.StatusOK("更新成功", data))
}

func HandleDeleteNode(c *gin.Context) {
	id := c.Param("id")
	elemID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(400, handler.StatusBad("参数错误", nil))
		return
	}

	if err := service.ListItemClient.DeleteNodeByID(uint(elemID)); err != nil {
		c.JSON(500, handler.StatusBad("出错了……", nil))
		return
	}

	// 清除缓存
	service.RDB.Del(handler.Ctx, "list_items")

	data, _ := GetData()
	c.JSON(200, handler.StatusOK("删除成功", data))
}
