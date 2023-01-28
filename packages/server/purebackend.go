package main

import (
	"github.com/PureML-Inc/PureML/server/apis"
	"github.com/PureML-Inc/PureML/server/core"
	"github.com/PureML-Inc/PureML/server/datastore"
)

func main() {
	core.Bootstrap()
	datastore.InitDB()
	apis.Serve()
}
