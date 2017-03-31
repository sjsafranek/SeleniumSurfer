package main

import "github.com/schollz/jsonstore"

const (
	DEFAULT_DATABASE_FILE string = "db.json.gz"
)

var (
	DATABASE_FILE string = DEFAULT_DATABASE_FILE
	DB            Database
)

type Database struct {
	Store *jsonstore.JSONStore
	Data  map[string]int
}

func (self Database) Add(message string, minutes int) {
	Ligneous.Debug("[Database] Add item ", message, " ", minutes)
	self.Data[message] = minutes
	self.Sync()
}

func (self Database) Sync() {
	//Ligneous.Info("[Database] Sync database")
	self.Store.Set("tasks", &self.Data)

	// Saving will automatically gzip if .gz is provided
	if err := jsonstore.Save(self.Store, DATABASE_FILE); err != nil {
		panic(err)
	}
}

func InitDatabase() {

	store, err := jsonstore.Open(DATABASE_FILE)
	if nil != err {
		store = new(jsonstore.JSONStore)
	}

	DB = Database{Store: store, Data: make(map[string]int)}

	err = DB.Store.Get("tasks", &DB.Data)
	if nil != err {
		DB.Sync()
	}

	Ligneous.Trace(DB.Data)

}
