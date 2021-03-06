package main

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
	"github.com/urfave/cli"
	"golang.org/x/oauth2"
)

var globalFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "from",
		Usage: "import target username",
	},
	cli.StringFlag{
		Name:  "token",
		Usage: "github personal access token for star import",
	},
}

func doImport(c *cli.Context) error {
	fromUser := c.String("from")
	token := c.String("token")

	if fromUser == "" || token == "" {
		return cli.NewExitError("insufficient parameters supplied to the command", 1)
	}

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
			return err
		}
		fmt.Printf("Starred : %s/%s\n", repo.Repository.Owner.GetLogin(), repo.Repository.GetName())
	}

	return nil
}
