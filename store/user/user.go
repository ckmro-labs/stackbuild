package user

import (
	"context"
	"fmt"
	"time"

	"github.com/laidingqing/stackbuild/core"
	"github.com/laidingqing/stackbuild/store/shared/db"
	"gopkg.in/mgo.v2/bson"
)

//UsersCollKey users mongo collection
const UsersCollKey = "users"

// errs definition
var (
	ErrorLoginNameExisted = fmt.Errorf("login name existed")
)

// New returns a new UserStore.
func New(db *db.SessionStore) core.UserStore {
	return &userStore{db}
}

type userStore struct {
	db *db.SessionStore
}

// Find returns a user from the datastore.
func (s *userStore) Find(ctx context.Context, id string) (*core.User, error) {
	defer s.db.Close()
	var user *core.User
	err := s.db.C(UsersCollKey).FindId(id).One(&user)
	return user, err
}

// FindLogin returns a user from the datastore by username.
func (s *userStore) FindLogin(ctx context.Context, login string) (*core.User, error) {
	query := bson.M{
		"login": login,
	}
	defer s.db.Close()
	var user *core.User
	err := s.db.C(UsersCollKey).Find(query).One(&user)
	return user, err
}

// FindToken returns a user from the datastore by token.
func (s *userStore) FindToken(ctx context.Context, token string) (*core.User, error) {
	query := bson.M{
		"token": token,
	}
	defer s.db.Close()
	var user *core.User
	err := s.db.C(UsersCollKey).Find(query).One(&user)
	return user, err
}

//Create persists a user.
func (s *userStore) Create(ctx context.Context, user *core.User) error {
	defer s.db.Close()
	//find unquie key
	if u, _ := s.FindLogin(ctx, user.Login); u.Login == user.Login {
		return ErrorLoginNameExisted
	}
	user.ID = bson.NewObjectId().Hex()
	user.Created = time.Now().UnixNano()
	err := s.db.C(UsersCollKey).Insert(user)
	return err
}
