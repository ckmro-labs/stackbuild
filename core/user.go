package core

import "context"

type (
	// User represents a user of the system.
	User struct {
		ID              string `bson:"_id" json:"id,omitempty"`
		Login           string `bson:"login" json:"login,omitempty"`
		Password        string `bson:"-" json:"password,omitempty"`
		EncryptPassword string `bson:"encryptPassword" json:"-"`
		Email           string `bson:"email" json:"email,omitempty"`
		Admin           bool   `bson:"admin"`
		Active          bool   `bson:"active" json:"active,omitempty"`
		Avatar          string `bson:"avatar" json:"avatar,omitempty"`
		Created         int64  `bson:"created" json:"created,omitempty"`
		Updated         int64  `bson:"updated" json:"updated,omitempty"`
		LastLogin       int64  `bson:"lastLogin" json:"last_login,omitempty"`
		Token           string `bson:"token" json:"token,omitempty"` //登录后需要更新它
		Synced          int64  `bson:"synced" json:"-,omitempty"`
	}

	// UserStore 用户存储接口.
	UserStore interface {
		// Find returns a user from the datastore.
		Find(context.Context, string) (*User, error)
		// FindLogin returns a user from the datastore by username.
		FindLogin(context.Context, string) (*User, error)
		// FindToken returns a user from the datastore by token.
		FindToken(context.Context, string) (*User, error)
		// Create persists a new user to the datastore.
		Create(context.Context, *User) error
	}
	// UserService 远程用户操作接口, 如github.com用户
	UserService interface {
		// Find returns the authenticated user.
		Find(ctx context.Context, access, refresh string) (*User, error)
	}
)
