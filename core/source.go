package core

import (
	"context"
	"time"
)

type (
	//SourceAuth for 3rd auth information.
	SourceAuth struct {
		ID       string    `bson:"_id" json:"_id"`
		UserID   string    `bson:"userId" json:"id"`
		UID      string    `bson:"uid" json:"uid"`
		AuthName string    `bson:"name" json:"name"`
		Token    string    `bson:"token" json:"token"`
		Refresh  string    `bson:"refresh" json:"refresh_token"`
		Expired  int64     `bson:"expired" json:"expired"`
		Created  time.Time `bson:"created" json:"created_at"`
		Updated  time.Time `bson:"updated" json:"updated_at"`
	}

	//SourceAuthStore store for authentication.
	SourceAuthStore interface {
		Create(context.Context, *SourceAuth) error
		Find(context.Context, string, string) (*SourceAuth, error)
		List(context.Context, string) ([]*SourceAuth, error)
	}
)
