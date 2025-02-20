# Go聊天室应用

这是一个基于Go语言开发的实时聊天室应用，支持用户注册、添加好友和实时聊天功能。

## 技术栈

- Go 1.16+
- Gin Web框架
- GORM ORM框架
- SQLite3 数据库
- Gorilla WebSocket

## 项目结构

```
go-chatroom/
├── main.go          # 主程序入口
├── models/          # 数据模型
│   ├── user.go      # 用户模型
│   └── message.go   # 消息模型
└── handlers/        # 处理器
    ├── user.go      # 用户相关处理
    └── chat.go      # 聊天相关处理
```

## 功能特性

1. 用户管理
   - 用户注册
   - 添加好友
   - 获取好友列表

2. 实时聊天
   - WebSocket实时通信
   - 私聊功能
   - 消息持久化

## API接口

### 1. 用户注册
```http
POST /register
Content-Type: application/json

{
    "username": "用户名",
    "password": "密码"
}
```

### 2. 添加好友
```http
POST /add_friend
Content-Type: application/json

{
    "friend_id": 2
}
```

### 3. 获取好友列表
```http
GET /friends
```

### 4. WebSocket聊天
- 连接地址：`ws://localhost:8080/ws`
- 消息格式：
```json
{
    "content": "消息内容",
    "to_id": 2
}
```

## 快速开始

1. 克隆项目
```bash
git clone https://github.com/yourusername/go-chatroom.git
cd go-chatroom
```

2. 安装依赖
```bash
go mod download
```

3. 运行项目
```bash
go run main.go
```

服务器将在 http://localhost:8080 启动

## 数据库设计

### User表
- ID: 用户ID（主键）
- Username: 用户名（唯一）
- Password: 密码
- CreatedAt: 创建时间
- UpdatedAt: 更新时间

### Message表
- ID: 消息ID（主键）
- Content: 消息内容
- FromID: 发送者ID
- ToID: 接收者ID
- CreatedAt: 发送时间

### UserFriend表
- UserID: 用户ID
- FriendID: 好友ID

## 待优化功能

1. 添加用户认证（JWT）
2. 实现群聊功能
3. 添加消息历史记录查询
4. 添加在线状态显示
5. 支持图片和文件传输
6. 添加消息加密功能
7. 实现消息推送通知

## 贡献指南

欢迎提交Issue和Pull Request来帮助改进项目。

## 许可证

MIT License
build a chatroom by golang
