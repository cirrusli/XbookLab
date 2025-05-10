package handlers

import (
	"net/http"
	"strconv"

	"xbooklab/models"

	"github.com/gin-gonic/gin"
)

type FollowUserRequest struct {
	UserID       uint `json:"user_id"`
	FollowUserID uint `json:"follow_user_id"`
}

type UnfollowUserRequest struct {
	UserID       uint `json:"user_id"`
	FollowUserID uint `json:"follow_user_id"`
}

type GetFollowingRequest struct {
	UserID uint `json:"user_id"`
}

type GetFollowersRequest struct {
	UserID uint `json:"user_id"`
}

type FollowListResponse struct {
	UserList []models.User `json:"user_list"`
}

type FollowResponse struct {
	Code    uint   `json:"code"`
	Message string `json:"message"`
}

// FollowUser 关注用户
func FollowUser(c *gin.Context) {
	userID := c.GetUint("userID")
	followedID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的用户ID"})
		return
	}

	if err := models.CreateFollow(userID, uint(followedID)); err != nil {
		if err == models.ErrCannotFollowSelf {
			c.JSON(http.StatusBadRequest, gin.H{"error": "不能关注自己"})
		} else if err == models.ErrAlreadyFollowed {
			c.JSON(http.StatusBadRequest, gin.H{"error": "已关注该用户"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "关注失败"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "关注成功"})
}

// UnfollowUser 取消关注用户
func UnfollowUser(c *gin.Context) {
	userID := c.GetUint("userID")
	followedID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的用户ID"})
		return
	}

	if err := models.DeleteFollow(userID, uint(followedID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "取消关注失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "取消关注成功"})
}

// GetFollowing 获取关注列表
func GetFollowing(c *gin.Context) {
	userID := c.GetUint("userID")
	follows, err := models.GetFollowing(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取关注列表失败"})
		return
	}

	// 提取关注用户信息
	followingUsers := make([]gin.H, len(follows))
	for i, follow := range follows {
		followingUsers[i] = gin.H{
			"user_id":  follow.Followed.UserID,
			"username": follow.Followed.Username,
			"avatar":   follow.Followed.Avatar,
		}
	}

	c.JSON(http.StatusOK, gin.H{"following": followingUsers})
}

// GetFollowers 获取粉丝列表
func GetFollowers(c *gin.Context) {
	userID := c.GetUint("userID")
	follows, err := models.GetFollowers(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取粉丝列表失败"})
		return
	}

	// 提取粉丝用户信息
	followers := make([]gin.H, len(follows))
	for i, follow := range follows {
		followers[i] = gin.H{
			"user_id":  follow.Follower.UserID,
			"username": follow.Follower.Username,
			"avatar":   follow.Follower.Avatar,
		}
	}

	c.JSON(http.StatusOK, gin.H{"followers": followers})
}
