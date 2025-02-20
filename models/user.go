package models

import (
	"time"
)

// User 用户模型
type User struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Username  string    `json:"username" gorm:"unique"`
	Password  string    `json:"-"`  // 密码不会在JSON中返回
	Friends   []User    `json:"friends" gorm:"many2many:user_friends;"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UserFriend 用户好友关系
type UserFriend struct {
	UserID   uint `gorm:"primary_key"`
	FriendID uint `gorm:"primary_key"`
}
