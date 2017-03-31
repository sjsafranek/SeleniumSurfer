package main

//import "os"
import "fmt"
import "github.com/schollz/jsonstore"

const (
	DEFAULT_DATABASE_FILE string = "db.json.gz"
)

var (
	DatabaseFile string = DEFAULT_DATABASE_FILE
	Database     jsonstore.JSONStore
	Tasks        Jobs
)

type Jobs struct {
	Jobs map[string]int
}

func (self Jobs) Add(message string, minutes int) {
	Ligneous.Debug("[Database] Adding task ", message, " ", minutes)
	self.Jobs[message] = minutes
}

func SyncDatabase() {
	Ligneous.Info("[Database] Sync database")

	Database.Set("tasks", &Tasks)

	// Saving will automatically gzip if .gz is provided
	if err := jsonstore.Save(&Database, DatabaseFile); err != nil {
		panic(err)
	}
}

func init() {

	Tasks = Jobs{make(map[string]int)}

	//Database = new(jsonstore.JSONStore)
	Database, err := jsonstore.Open(DatabaseFile)
	if nil != err {
		Database = new(jsonstore.JSONStore)
	}

	err = Database.Get("tasks", &Tasks)
	if nil != err {
		// Saving will automatically gzip if .gz is provided
		if err = jsonstore.Save(Database, DatabaseFile); err != nil {
			panic(err)
		}
	}

	fmt.Printf("%v\n", Tasks)

	//	Tasks.Add(fmt.Sprintf("%v", time.Now()), 1)

	//	os.Exit(0)

	//data := Database.GetAll("*")
	//fmt.Printf("%v\n", data)
	//Ligneous.Info("job")
}
