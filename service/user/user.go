package user

import (
	"context"

	"github.com/drone/go-scm/scm"
	"github.com/laidingqing/stackbuild/core"
)

type service struct {
	client *scm.Client //TODO need new instance by context, not inject.
}

// New returns a new User service that provides source code management interface.
func New() core.UserService {
	return &service{}
}

func (s *service) Find(ctx context.Context, access, refresh string) (*core.User, error) {
	ctx = context.WithValue(ctx, scm.TokenKey{}, &scm.Token{
		Token:   access,
		Refresh: refresh,
	})
	src, _, err := s.client.Users.Find(ctx)
	if err != nil {
		return nil, err
	}
	dst := &core.User{
		Login:  src.Login,
		Email:  src.Email,
		Avatar: src.Avatar,
	}
	if !src.Created.IsZero() {
		dst.Created = src.Created.Unix()
	}
	if !src.Updated.IsZero() {
		dst.Updated = src.Updated.Unix()
	}
	return dst, nil
}
