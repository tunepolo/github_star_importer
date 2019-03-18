package main

import (
	"os"

	"github.com/urfave/cli"
)

func main() {
	newApp().Run(os.Args)
}

func newApp() *cli.App {
	app := cli.NewApp()

	app.Name = "Github Star Importer"
	app.Usage = "指定アカウントのStar情報を別アカウントにインポートする"
	app.Version = "0.0.1"

	app.Flags = globalFlags
	app.Action = doImport

	return app
}
