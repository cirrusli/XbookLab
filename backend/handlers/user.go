package handlers

import (
	"net/http"
	"strconv"
	"xbooklab/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
	Bio      string `json:"bio"`
	Email    string `json:"email"`
}

type UpdateUserRequest struct {
	Username string `json:"name"`
	Avatar   string `json:"avatar"`
	Bio      string `json:"bio"`
}

func Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	user := models.User{
		Username: req.Username,
		Password: string(hashedPassword),
		Avatar:   req.Avatar,
		Bio:      req.Bio,
	}
	if err := models.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "注册失败" + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"code": 200,
		"data":"",
		"message": "注册成功，请登录"})
}

func GetUserProfile(c *gin.Context) {
	userID := c.GetUint("userID")
	if userID == 0 {
		userID = 1
	}
	user, err := models.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "用户不存在"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{
			"id":            user.UserID,
			"name":          user.Username,
			"avatar":        user.Avatar,
			"bio":           user.Bio,
			"followCount":   6,
			"followMeCount": 31,
			"topicCount":    15,
			"bookCount":     20,
		},
		"message": "用户信息获取成功",
	})
}

func UpdateUserProfile(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	var req UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := models.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "用户不存在"})
		return
	}
	user.Username = req.Username
	user.Avatar = req.Avatar
	user.Bio = req.Bio

	if err := models.UpdateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新用户信息失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": "",
		"message": "用户信息更新成功"})
}

func GetUserList(c *gin.Context) {
	users, err := models.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取用户列表失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	userID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "用户 ID 格式错误"})
		return
	}
	if err := models.DeleteUser(uint(userID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "删除用户失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "用户删除成功"})
}
