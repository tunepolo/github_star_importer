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
			Usage: "import target username",
		},
		cli.StringFlag{
			Name:  "token",
			Value: "PERSONAL_ACCESS_TOKEN",
			Usage: "github personal access token for star import",
		},
	}

	app.Action = func(c *cli.Context) error {
		fromUser := c.String("from")
		token := c.String("token")

		// トークンを使ってGitHubアクセスのためのclientを生成
		ctx := context.Background()
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		)
		tc := oauth2.NewClient(ctx, ts)
		client := github.NewClient(tc)

		// fromUserがStarをつけたリポジトリ一覧を取得
		opt := &github.ActivityListStarredOptions{
			ListOptions: github.ListOptions{PerPage: 100},
		}
		var allRepositories []*github.StarredRepository
		for {
			repos, response, err := client.Activity.ListStarred(ctx, fromUser, opt)
			if err != nil {
				return err
			}
			allRepositories = append(allRepositories, repos...)
			if response.NextPage == 0 {
				break
			}
			opt.Page = response.NextPage
		}

		// 取得したリポジトリのリストを元にStarを付与
		for _, repo := range allRepositories {
			_, err := client.Activity.Star(ctx, repo.Repository.Owner.GetLogin(), repo.Repository.GetName())
			if err != nil {
				fmt.Printf("Starred : %s/%s\n", repo.Repository.Owner.GetLogin(), repo.Repository.GetName())
			}
		}
		return nil
	}

	app.Run(os.Args)
}
