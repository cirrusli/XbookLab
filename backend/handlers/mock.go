package handlers

import (
	"math/rand"
	"net/http"
	"xbooklab/models"

	"github.com/gin-gonic/gin"
)

func GetGroups(c *gin.Context) {
	allGroups := models.GetAllGroups()

	// 如果有tag查询参数，则进行过滤
	if tag := c.Query("tag"); tag != "" {
		filteredGroups := make([]gin.H, 0)
		for _, group := range allGroups {
			if group["tagVal"] == tag {
				filteredGroups = append(filteredGroups, group)
			}
		}
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"data":    filteredGroups,
			"message": "获取小组列表成功",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"data":    allGroups,
		"message": "获取小组列表成功",
	})

}

func GetFriendsBooks(c *gin.Context) {
	allFriendsBooks := models.GetAllFriendsBooks()
	filteredBooks := allFriendsBooks

	// 如果有type查询参数，则进行过滤
	if typeParam := c.Query("type"); typeParam != "0" {
		filteredBooks = make([]gin.H, 0)
		for _, book := range allFriendsBooks {
			if book["type"] == typeParam {
				filteredBooks = append(filteredBooks, book)
			}
		}
	}

	// 随机排序
	for i := range filteredBooks {
		j := rand.Intn(i + 1)
		filteredBooks[i], filteredBooks[j] = filteredBooks[j], filteredBooks[i]
	}

	c.JSON(http.StatusOK, gin.H{
		"code":  200,
		"data":  filteredBooks,
		"total": len(filteredBooks),
	})
}
