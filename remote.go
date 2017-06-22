package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/qjw/git-notify/remote"
	"github.com/qjw/git-notify/remote/gitlab"
)


func GetGitlabFlags() []cli.Flag {
	return []cli.Flag{
		cli.BoolFlag{
			EnvVar: "GIT_NOTIFY_GITLAB",
			Name:   "gitlab",
			Usage:  "gitlab是否启用",
		},
		cli.StringFlag{
			EnvVar: "GIT_NOTIFY_GITLAB_URL",
			Name:   "gitlab-server",
			Usage:  "gitlab仓库地址",
			Value:  "https://gitlab.com",
		},
		cli.StringFlag{
			EnvVar: "GIT_NOTIFY_GITLAB_TOKEN",
			Name:   "gitlab-patoken",
			Usage:  "gitlab personal access token",
		},
	}
}

// helper function to setup the remote from the CLI arguments.
func setupRemote(c *cli.Context) (remote.Remote, error) {
	switch {
	case c.Bool("gitlab"):
		return setupGitlab(c)
	default:
		return nil, fmt.Errorf("version control system not configured")
	}
}


// helper function to setup the Gitlab remote from the CLI arguments.
func setupGitlab(c *cli.Context) (remote.Remote, error) {
	return gitlab.New(gitlab.Opts{
		URL:         c.String("gitlab-server"),
		Token:      c.String("gitlab-patoken"),
	})
}

