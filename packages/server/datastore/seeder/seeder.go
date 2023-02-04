package main

import (
	"log"
	"os"

	"github.com/PureML-Inc/PureML/server/datastore/impl"
	"github.com/PureML-Inc/PureML/server/datastore/seeds"
)

func main() {
	var forTestDb bool
	if len(os.Args) < 2 {
		forTestDb = false
	} else {
		forTestDb = os.Args[1] == "test"
	}
	var ds *impl.Datastore
	if forTestDb {
		ds = impl.NewSQLiteDatastore("../../tests/data")
	} else {
		ds = impl.NewSQLiteDatastore()
	}
	dbConn, err := ds.DB.DB()
	if err != nil {
		log.Fatalf("Couldn't establish database connection: %s", err)
	}
	defer dbConn.Close()

	for _, seed := range seeds.All() {
		if err := seed.Run(ds.DB); err != nil {
			log.Fatalf("Running seed '%s', failed with error: %s", seed.Name, err)
		}
	}

	log.Println("Seeding complete")
}
