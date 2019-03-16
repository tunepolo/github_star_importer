package main

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/github"
	"github.com/urfave/cli"
	"golang.org/x/oauth2"
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
		ctx := context.Background()
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: c.String("token")},
		)
		tc := oauth2.NewClient(ctx, ts)

		client := github.NewClient(tc)
		repositories, _, err := client.Repositories.List(ctx, "", nil)

		for _, repo := range repositories {
			fmt.Printf("%+v\n", repo.GetFullName())
		}

		return err
	}

	app.Run(os.Args)
}
