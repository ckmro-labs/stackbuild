package core

import "context"

type (
	// User represents a user of the system.
	User struct {
		ID        int64  `json:"id"`
		Login     string `json:"login"`
		Email     string `json:"email"`
		Admin     bool   `json:"admin"`
		Active    bool   `json:"active"`
		Avatar    string `json:"avatar"`
		Syncing   bool   `json:"syncing"`
		Synced    int64  `json:"synced"`
		Created   int64  `json:"created"`
		Updated   int64  `json:"updated"`
		LastLogin int64  `json:"last_login"`
		Token     string `json:"-"`
		Refresh   string `json:"-"`
		Expiry    int64  `json:"-"`
		Hash      string `json:"-"`
	}

	// UserStore 用户存储接口.
	UserStore interface {
	}
	// UserService 远程用户操作接口, 如github.com用户
	UserService interface {
		// Find returns the authenticated user.
		Find(ctx context.Context, access, refresh string) (*User, error)
	}
)
