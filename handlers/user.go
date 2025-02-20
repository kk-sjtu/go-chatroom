package handlers

import (
	"chatroom/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// UserHandler 处理用户相关的请求
type UserHandler struct {
	db *gorm.DB
}

// NewUserHandler 创建新的UserHandler
func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{db: db}
}

// Register 用户注册
func (h *UserHandler) Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "注册失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "注册成功", "user": user})
}

// AddFriend 添加好友
func (h *UserHandler) AddFriend(c *gin.Context) {
	userID := c.GetUint("user_id")
	var friendID uint
	if err := c.ShouldBindJSON(&struct{ FriendID uint `json:"friend_id"` }{friendID}); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 检查好友是否存在
	var friend models.User
	if err := h.db.First(&friend, friendID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "好友不存在"})
		return
	}

	// 添加好友关系
	friendship := models.UserFriend{
		UserID:   userID,
		FriendID: friendID,
	}
	if err := h.db.Create(&friendship).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "添加好友失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "添加好友成功"})
}

// GetFriends 获取好友列表
func (h *UserHandler) GetFriends(c *gin.Context) {
	userID := c.GetUint("user_id")
	var user models.User
	if err := h.db.Preload("Friends").First(&user, userID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取好友列表失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"friends": user.Friends})
}
