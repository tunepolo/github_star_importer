package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "Github Star Importer"
	app.Usage = "指定アカウントのStar情報を別アカウントにインポートする"
	app.Version = "0.0.1"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "from, f",
			Value: "tune",
			Usage: "import FROM username",
		},
		cli.StringFlag{
			Name:  "to, t",
			Value: "tunepolo",
			Usage: "import TO username",
		},
		cli.StringFlag{
			Name:  "token",
			Value: "PERSONAL_ACCESS_TOKEN",
			Usage: "github personal access token for star import",
		},
	}

	app.Action = func(c *cli.Context) error {
		fmt.Printf("from: %+v\n", c.String("from"))
		fmt.Printf("to: %+v\n", c.String("to"))
		fmt.Printf("token: %+v\n", c.String("token"))
		return nil
	}

	app.Run(os.Args)
}
