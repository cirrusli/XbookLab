package models

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetAllGroups() []gin.H {
	return []gin.H{
		{
			"id":      1,
			"title":   "深度阅读俱乐部",
			"cover":   "https://images.unsplash.com/photo-1544716278-ca5e3f4abd8c",
			"members": 1245,
			"topics":  89,
			"desc":    "专注深度阅读与思考的交流平台",
			"tagVal":  "reading",
			"tag":     "阅读",
		},
		{
			"id":      2,
			"title":   "科技前沿讨论组",
			"cover":   "https://images.unsplash.com/photo-1620712943543-bcc4688e7485",
			"members": 892,
			"topics":  156,
			"desc":    "探讨最新科技趋势与创新技术",
			"tagVal":  "technology",
			"tag":     "科技",
		},
		{
			"id":      3,
			"title":   "古典文学研究会",
			"cover":   "https://images.unsplash.com/photo-1544716278-ca5e3f4abd8c",
			"members": 673,
			"topics":  234,
			"desc":    "研究古典文学作品的学术小组",
			"tagVal":  "literature",
			"tag":     "文学",
		},
		{
			"id":      4,
			"title":   "心理学读书会",
			"cover":   "https://images.unsplash.com/photo-1544716278-ca5e3f4abd8c",
			"members": 1102,
			"topics":  187,
			"desc":    "心理学爱好者交流与学习平台",
			"tagVal":  "psychology",
			"tag":     "心理学",
		},
		{
			"id":      5,
			"title":   "历史探索者",
			"cover":   "https://images.unsplash.com/photo-1544716278-ca5e3f4abd8c",
			"members": 756,
			"topics":  143,
			"desc":    "探索历史真相，分享历史见解",
			"tagVal":  "history",
			"tag":     "历史",
		},
		{
			"id":      6,
			"title":   "商业思维训练营",
			"cover":   "https://images.unsplash.com/photo-1551288049-bebda4e38f71",
			"members": 983,
			"topics":  211,
			"desc":    "提升商业思维与决策能力",
			"tagVal":  "business",
			"tag":     "商业",
		},
	}
}

func GetAllFriendsBooks() []gin.H {
	count := rand.Intn(11) + 10 // 生成10-20之间的随机数

	books := make([]gin.H, count)
	for i := 0; i < count; i++ {
		books[i] = gin.H{
			"id":           i + 1,
			"name":         "用户" + strconv.Itoa(i+1),
			"avatar":       "https://picsum.photos/100/100?random=" + strconv.Itoa(i),
			"reading":      "正在阅读",
			"bookTitle":    "书籍" + strconv.Itoa(i+1),
			"bookCover":    "https://picsum.photos/150/200?random=" + strconv.Itoa(i),
			"bookProgress": rand.Intn(100) + 1,
			"bookComment":  "这是一本很好的书",
			"lastActive":   time.Now().Format("2006-01-02 15:04"),
			"type":         strconv.Itoa(rand.Intn(2) + 1), // 1或2
		}
	}
	return books
}
