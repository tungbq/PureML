package main

import (
	"github.com/PureML-Inc/PureML/server/apis"
	"github.com/PureML-Inc/PureML/server/core"
	"github.com/PureML-Inc/PureML/server/datastore"
)

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	contact@pureml.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization
//	@description				Header for logged in user format: Bearer {token}
func main() {
	core.Bootstrap()
	datastore.InitDB()
	apis.Serve()
}
