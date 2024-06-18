package feedback

import (
	"TGU-MAP/models"
	"TGU-MAP/service"
	"TGU-MAP/utils"
	"TGU-MAP/web/handler"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var pageSize = 10

func GetData(pageNum int) (string, int64, error) {
	offset := (pageNum - 1) * pageSize
	data, err := service.FeedbackClient.Paginate(offset, pageSize)
	if err != nil {
		return "", 0, err
	}
	total, err := service.FeedbackClient.Count()
	if err != nil {
		return "", 0, err
	}
	cache := string(utils.Marshal(*data))

	return cache, total, nil

}

func HandleGetData(c *gin.Context) {
	pageNum, err := strconv.Atoi(c.Query("pageNum"))
	if err != nil {
		c.JSON(http.StatusBadRequest, handler.StatusBad("参数错误", nil))
		return
	}
	if pageNum < 1 {
		pageNum = 1
	}
	if data, total, err := GetData(pageNum); err != nil {
		c.JSON(500, handler.StatusBad("出错了……", nil))
	} else {
		c.JSON(200, handler.StatusOK("查询成功", gin.H{
			"data":        data,
			"totalPage":   total/int64(pageSize) + 1,
			"currentPage": pageNum,
			"total":       total,
		}))
	}

}

func HandleInsertNode(c *gin.Context) {
	var item models.Feedback
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, handler.StatusBad("参数错误", nil))
		return
	}
	if err := service.FeedbackClient.InsertNode(&item); err != nil {
		c.JSON(500, handler.StatusBad("出错了……", nil))
		return
	}

	c.JSON(200, handler.StatusOK("添加成功", nil))

}

func HandleDeleteNode(c *gin.Context) {
	id := c.Param("id")
	pageNum, err := strconv.Atoi(c.Query("pageNum"))
	if err != nil {
		c.JSON(http.StatusBadRequest, handler.StatusBad("参数错误", nil))
		return
	}
	if pageNum < 1 {
		pageNum = 1
	}
	elemID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(400, handler.StatusBad("参数错误", nil))
		return
	}
	if err := service.FeedbackClient.DeleteNodeByID(uint(elemID)); err != nil {
		c.JSON(500, handler.StatusBad("出错了……", nil))
		return
	}
	if data, total, err := GetData(pageNum); err != nil {
		c.JSON(200, handler.StatusBad("删除成功", nil))
	} else {
		c.JSON(200, handler.StatusOK("删除成功", gin.H{
			"data":        data,
			"totalPage":   total/int64(pageSize) + 1,
			"currentPage": pageNum,
			"total":       total,
		}))
	}

}
