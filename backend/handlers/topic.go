package handlers

import (
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"xbooklab/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GetTopicListRequest struct {
	Offset    uint `json:"offset"`
	Limit     uint `json:"limit"`
	TagFilter uint `json:"tag_filter"`
}

type GetTopicListResponse struct {
	Topics []Topic `json:"topics"`
	Total  uint    `json:"total"`
}

type CreateTopicRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	TagName string `json:"tag"`
}

type CreateTopicResponse struct {
	TopicID uint   `json:"topic_id"`
	Code    uint   `json:"code"`
	Message string `json:"message"`
}

type GetTopicDetailRequest struct {
	TopicID uint `json:"topic_id"`
}

type GetTopicDetailResponse struct {
	Topic Topic `json:"topic"`
}

type LikeTopicRequest struct {
	UserID  uint `json:"user_id"`
	TopicID uint `json:"topic_id"`
}

type LikeTopicResponse struct {
	Code    uint   `json:"code"`
	Message string `json:"message"`
}

type DeleteTopicRequest struct {
	TopicID uint `json:"topic_id"`
}

type DeleteTopicResponse struct {
	Code    uint   `json:"code"`
	Message string `json:"message"`
}

type Topic struct {
	TopicID      uint      `json:"topic_id"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	Author       string    `json:"author"`
	AuthorUserID uint      `json:"author_user_id"`
	Views        uint      `json:"views"`
	LikeCount    uint      `json:"like_count"`
	Comments     uint      `json:"comments"`
	TagName      string    `json:"tag_name"`
	CreatedAt    time.Time `json:"created_at"`
	Date         string    `json:"date"`
	Timestamp    int64     `json:"timestamp"`
}

// 创建话题
func CreateTopic(c *gin.Context) {
	var topic models.Topic
	var req CreateTopicRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID := c.GetUint("userID")
	if userID == 0 {
		userID = 1
	}
	topic.Title = req.Title
	topic.Content = req.Content
	topic.TagName = req.TagName
	topic.AuthorUserID = userID
	topic.LikeCount = 0
	// if len(topic.Tag) == 0 {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "必须选择至少一个分类标签"})
	// 	return
	// }

	if err := models.DB.Create(&topic).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建话题失败"})
		return
	}

	// 处理标签关联
	if req.TagName != "" {
		// 查找或创建标签
		var tag models.Tag
		if err := models.DB.Where("tag_name = ?", req.TagName).FirstOrCreate(&tag, models.Tag{TagName: req.TagName}).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "处理标签失败"})
			return
		}

		// 创建话题标签关联
		topicTag := models.TopicTag{
			TopicID: topic.TopicID,
			TagID:   tag.TagID,
		}
		if err := models.DB.Create(&topicTag).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "创建标签关联失败"})
			return
		}
	}

	c.JSON(http.StatusCreated, gin.H{
		"code":    200,
		"message": "创建成功",
		"data":    topic,
	})
}

// 获取话题列表
func GetTopics(c *gin.Context) {
	var topics []Topic
	query := models.DB

	// 获取tag筛选参数
	tagID := c.Query("tag")
	if tagID != "" {
		query = query.Joins("JOIN topic_tags ON topics.topic_id = topic_tags.topic_id").
			Where("topic_tags.tag_id = ?", tagID)
	}
	timeRange := c.Query("range")
	if timeRange != "" {
		switch timeRange {
		case "today":
			query = query.Where("created_at >= ?", time.Now().In(time.Local).AddDate(0, 0, -1))
		case "week":
			query = query.Where("created_at >= ?", time.Now().In(time.Local).AddDate(0, 0, -7))
		case "month":
			query = query.Where("created_at >= ?", time.Now().In(time.Local).AddDate(0, -1, 0))
		default:
			query = query.Where("created_at >= ?", time.Now().In(time.Local).AddDate(0, 0, -7))
		}
	}
	if err := query.Select("topics.topic_id, topics.title, topics.content, topics.author_user_id, topics.like_count, topics.created_at").Find(&topics).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取话题列表失败"})
		return
	}

	// 查询每个话题的标签名称和其他信息
	for i := range topics {
		var tag models.Tag
		if err := models.DB.Table("topic_tags").
			Select("tags.tag_name").
			Joins("join tags on topic_tags.tag_id = tags.tag_id").
			Where("topic_tags.topic_id = ?", topics[i].TopicID).
			First(&tag).Error; err == nil {
			topics[i].TagName = tag.TagName
		}
		topics[i].Date = topics[i].CreatedAt.Format("2006-01-02 15:04:05")
		topics[i].Timestamp = topics[i].CreatedAt.Unix()
		topics[i].Views = uint(1000 + rand.Intn(9001)) // 随机1000-10000之间的浏览量
		topics[i].Comments = uint(10 + rand.Intn(91))  // 随机10-100之间的评论数
		topics[i].Author = "书友圈圈友"
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"data":    topics,
		"message": "获取成功",
	})
}

// 获取单个话题
func GetTopicDetail(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的话题ID"})
		return
	}

	var topic models.Topic

	if err := models.DB.First(&topic, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "话题不存在"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "获取话题失败"})
		}
		return
	}

	// 查询话题关联的标签
	var tag models.Tag
	if err := models.DB.Table("topic_tags").
		Select("tags.tag_name").
		Joins("join tags on topic_tags.tag_id = tags.tag_id").
		Where("topic_tags.topic_id = ?", topic.TopicID).
		First(&tag).Error; err == nil {
		topic.TagName = tag.TagName
	}

	// 检查当前用户是否点赞过该话题
	userID := c.GetUint("userID")
	if userID > 0 {
		var like models.TopicLike
		if err := models.DB.Where("user_id = ? AND topic_id = ?", userID, topic.TopicID).First(&like).Error; err == nil {
			topic.IsLiked = 1
		} else {
			topic.IsLiked = 0
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"data":    topic,
		"message": "获取成功",
	})
}

// 更新话题
func UpdateTopic(c *gin.Context) {
	_, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的话题ID"})
		return
	}

	var topic models.Topic
	if err := c.ShouldBindJSON(&topic); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// if len(topic.Tag) == 0 {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "必须选择至少一个分类标签"})
	// 	return
	// }

	// if len(topic.Tag) == 0 {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "必须选择至少一个分类标签"})
	// 	return
	// }

	// topic.ID = uint(id)

	if err := models.DB.Save(&topic).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新话题失败"})
		return
	}

	c.JSON(http.StatusOK, topic)
}

// 删除话题
func DeleteTopic(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的话题ID"})
		return
	}

	if err := models.DB.Delete(&models.Topic{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "删除话题失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "话题删除成功"})
}
