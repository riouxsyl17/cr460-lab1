package datastore

import mgo "gopkg.in/mgo.v2"

// Datastore Datastore
type Datastore struct {
	dbname  string
	session *mgo.Session
}

// New datastore
func New(uri string, db string) (ds *Datastore, err error) {
	ds = &Datastore{}
	ds.session, err = mgo.Dial(uri)
	ds.dbname = db
	return
}

// GetDBSession to the database
func (ds *Datastore) GetDBSession() *mgo.Database {

	return ds.session.Copy().DB(ds.dbname)

}
