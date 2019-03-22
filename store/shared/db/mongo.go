package db

import (
	"log"

	mgo "gopkg.in/mgo.v2"
)

// SessionStore  is the type for a database session
type SessionStore struct {
	Database string
	Session  *mgo.Session
}

var mainSession *mgo.Session

// NewSessionStore  returns a new SessionStore  with a copied session
func NewSessionStore(host string, database string) *SessionStore {
	var err error
	mainSession, err = mgo.Dial(host)
	if err != nil {
		log.Fatalf("can't connect db..%v", host)
	}
	ds := &SessionStore{
		Session: mainSession.Copy(),
	}
	return ds
}

// C get collection.
func (ds *SessionStore) C(name string) *mgo.Collection {
	return ds.Session.DB(ds.Database).C(name)
}

// Close close session.
func (ds *SessionStore) Close() {
	ds.Session.Close()
}
