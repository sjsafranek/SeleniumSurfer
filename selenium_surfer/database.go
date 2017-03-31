package main

import "fmt"
import "github.com/schollz/jsonstore"

const (
	DEFAULT_DATABASE_FILE string = "db.json.gz"
)

var (
	DatabaseFile string = DEFAULT_DATABASE_FILE
	Database     jsonstore.JSONStore
)

type Job struct {
	Message string
}

func init() {
	//Database = new(jsonstore.JSONStore)
	Database, err := jsonstore.Open(DatabaseFile)
	if nil != err {
		Database = new(jsonstore.JSONStore)
	}
	Database.Set("job1", Job{"This is a job"})

	// Saving will automatically gzip if .gz is provided
	if err = jsonstore.Save(Database, DatabaseFile); err != nil {
		panic(err)
	}

	var job Job
	err = Database.Get("job1", &job)
	if nil != err {
		panic(err)
	}

	fmt.Printf("%v\n", job)

	//Ligneous.Info("job")
}
