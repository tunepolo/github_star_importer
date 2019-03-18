package main

import (
	"os"

	"github.com/urfave/cli"
)

var Version string = "0.0.1"

func main() {
	newApp().Run(os.Args)
}

func newApp() *cli.App {
	app := cli.NewApp()

	app.Name = "Github Star Importer"
	app.Usage = "指定アカウントのStar情報を別アカウントにインポートする"
	app.Version = Version
	app.Author = "Yuichi Tsunematsu"
	app.Email = "tunepolo@gmail.com"
	app.Flags = globalFlags
	app.Action = doImport

	return app
}
