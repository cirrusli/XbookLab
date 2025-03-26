package handlers

import (
	"canigraduate/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Bio      string `json:"bio"`
	Email    string `json:"email"`
}

type UpdateUserRequest struct {
	Nickname string `json:"nickname"`
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
		Username:     req.Username,
		PasswordHash: string(hashedPassword),
		Nickname:     req.Nickname,
		Avatar:       req.Avatar,
		Bio:          req.Bio,
		Email:        req.Email,
	}
	if err := models.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "注册失败"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "注册成功"})
}

func GetUserProfile(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	user, err := models.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "用户不存在"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":       user.ID,
		"username": user.Username,
		"nickname": user.Nickname,
		"avatar":   user.Avatar,
		"bio":      user.Bio,
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

	user.Nickname = req.Nickname
	user.Avatar = req.Avatar
	user.Bio = req.Bio

	if err := models.UpdateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新用户信息失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "用户信息更新成功"})
}

func UploadAvatar(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	file, err := c.FormFile("avatar")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请上传头像文件"})
		return
	}

	// 这里应该实现文件保存逻辑，返回文件路径
	avatarPath := "/uploads/" + file.Filename
	if err := c.SaveUploadedFile(file, "."+avatarPath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "保存头像文件失败"})
		return
	}

	user, err := models.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "用户不存在"})
		return
	}

	user.Avatar = avatarPath
	if err := models.UpdateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新用户头像失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"avatar":  avatarPath,
		"message": "头像上传成功",
	})
}
