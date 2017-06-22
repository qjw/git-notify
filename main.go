package main

import (
	"os"
	"github.com/ianschenck/envflag"
	"github.com/codegangsta/cli"
	"time"
	"math/rand"
)

func main() {
	rand.Seed(time.Now().Unix())
	envflag.Parse()
	app := cli.NewApp()
	app.Name = "git-notify"
	app.Version = "0.0.0.1"
	app.Usage = "git-notify命令行工具"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			EnvVar: "GIT_NOTIFY_DEBUG",
			Name:   "d,debug",
			Usage:  "开启调试模式",
		},
	}
	app.Commands = []cli.Command{
		serverCmd,
		hookCmd,
	}
	app.Run(os.Args)
}