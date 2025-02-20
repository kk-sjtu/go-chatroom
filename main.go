package main

import (
	"chatroom/handlers"
	"chatroom/models"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	// 连接数据库
	db, err := gorm.Open("sqlite3", "chatroom.db")
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}
	defer db.Close()

	// 自动迁移数据库表
	db.AutoMigrate(&models.User{}, &models.Message{}, &models.UserFriend{})

	// 创建处理器
	userHandler := handlers.NewUserHandler(db)
	chatHandler := handlers.NewChatHandler(db)

	// 设置路由
	r := gin.Default()

	// 用户相关路由
	r.POST("/register", userHandler.Register)
	r.POST("/add_friend", userHandler.AddFriend)
	r.GET("/friends", userHandler.GetFriends)

	// WebSocket路由
	r.GET("/ws", chatHandler.HandleWebSocket)

	// 启动服务器
	if err := r.Run(":8080"); err != nil {
		log.Fatal("服务器启动失败:", err)
	}
}
