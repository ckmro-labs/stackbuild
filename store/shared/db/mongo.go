package db

import (
	mgo "gopkg.in/mgo.v2"
)

// SessionStore  is the type for a database session
type SessionStore struct {
	Session *mgo.Session
}

var mainSession *mgo.Session

// NewSessionStore  returns a new SessionStore  with a copied session
func NewSessionStore() *SessionStore {
	ds := &SessionStore{
		Session: mainSession.Copy(),
	}
	return ds
}

//Connect .connect db.
func Connect(source string) error {
	var err error
	mainSession, err = mgo.Dial(source)
	return err
}

// Close close session.
func (ds *SessionStore) Close() {
	ds.Session.Close()
}
