package core

import "context"

type (
	// User represents a user of the system.
	User struct {
		ID              string           `json:"id"`
		Login           string           `json:"login"`
		Email           string           `json:"email"`
		Admin           bool             `json:"admin"`
		Active          bool             `json:"active"`
		Avatar          string           `json:"avatar"`
		Syncing         bool             `json:"syncing"`
		Synced          int64            `json:"synced"`
		Created         int64            `json:"created"`
		Updated         int64            `json:"updated"`
		LastLogin       int64            `json:"last_login"`
		Token           string           `json:"token"` //登录后需要更新它
		Authentications []Authentication `json:"authentications"`
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
