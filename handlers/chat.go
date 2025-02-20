package handlers

import (
	"chatroom/models"
	"encoding/json"
	"log"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/jinzhu/gorm"
)

// ChatHandler 处理聊天相关的请求
type ChatHandler struct {
	db       *gorm.DB
	clients  map[uint]*websocket.Conn
	mutex    sync.RWMutex
	upgrader websocket.Upgrader
}

// NewChatHandler 创建新的ChatHandler
func NewChatHandler(db *gorm.DB) *ChatHandler {
	return &ChatHandler{
		db:      db,
		clients: make(map[uint]*websocket.Conn),
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true // 允许所有来源的WebSocket连接
			},
		},
	}
}

// HandleWebSocket 处理WebSocket连接
func (h *ChatHandler) HandleWebSocket(c *gin.Context) {
	userID := c.GetUint("user_id")
	conn, err := h.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("WebSocket升级失败: %v", err)
		return
	}

	// 保存连接
	h.mutex.Lock()
	h.clients[userID] = conn
	h.mutex.Unlock()

	// 处理接收到的消息
	go h.handleMessages(userID, conn)
}

// handleMessages 处理接收到的消息
func (h *ChatHandler) handleMessages(userID uint, conn *websocket.Conn) {
	defer func() {
		conn.Close()
		h.mutex.Lock()
		delete(h.clients, userID)
		h.mutex.Unlock()
	}()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("读取消息失败: %v", err)
			break
		}

		var msg models.Message
		if err := json.Unmarshal(message, &msg); err != nil {
			log.Printf("解析消息失败: %v", err)
			continue
		}

		msg.FromID = userID
		// 保存消息到数据库
		if err := h.db.Create(&msg).Error; err != nil {
			log.Printf("保存消息失败: %v", err)
			continue
		}

		// 发送消息给接收者
		h.mutex.RLock()
		if toConn, ok := h.clients[msg.ToID]; ok {
			if err := toConn.WriteJSON(msg); err != nil {
				log.Printf("发送消息失败: %v", err)
			}
		}
		h.mutex.RUnlock()
	}
}
