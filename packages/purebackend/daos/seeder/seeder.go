package main

import (
	"fmt"
	"log"
	"os"

	impl "github.com/PureML-Inc/PureML/packages/purebackend/daos/datastore"
	"github.com/PureML-Inc/PureML/packages/purebackend/daos/seeds"
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
		fmt.Println("Seeding test database")
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
			log.Printf("Running seed '%s', failed with error: %s", seed.Name, err)
		}
	}

	log.Println("Seeding complete")
}
