package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"gopkg.in/redis.v5"
	"strconv"
	"github.com/qjw/git-notify/qywechat"
	"github.com/qjw/git-notify/gateway"
	"github.com/qjw/git-notify/utils"
	"github.com/qjw/kelly"
)


func GetCorpFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			EnvVar: "GIT_NOTIFY_CORP_ID",
			Name:   "corp_id",
			Usage:  "企业微信corpid",
			Value:  "",
		},
		cli.StringFlag{
			EnvVar: "GIT_NOTIFY_CORP_SECRET",
			Name:   "corp_secret",
			Usage:  "企业微信应用secret",
			Value:  "",
		},
		cli.Int64Flag{
			EnvVar: "GIT_NOTIFY_CORP_SECRET",
			Name:   "corp_appid",
			Usage:  "企业微信应用id",
			Value:  0,
		},
		cli.StringFlag{
			EnvVar: "GIT_NOTIFY_CORP_SECRET",
			Name:   "receive_department",
			Usage:  "接收消息的部门",
			Value:  "",
		},
	}
}

func CombineSliceArray(first []cli.Flag, lst ...[]cli.Flag) []cli.Flag {

	for _, item := range lst {
		first = append(first, item...)
	}
	return first
}

var serverCmd = cli.Command{
	Name:  "server",
	Usage: "启动账户服务器",
	Action: func(c *cli.Context) {
		if err := server(c); err != nil {
			log.Fatal(err)
		}
	},
	Flags: CombineSliceArray([]cli.Flag{
		cli.StringFlag{
			EnvVar: "GIT_NOTIFY_HOST",
			Name:   "host",
			Usage:  "服务器主机",
			Value:  "localhost",
		},
		cli.IntFlag{
			EnvVar: "GIT_NOTIFY_PORT",
			Name:   "port",
			Usage:  "服务器端口",
			Value:  13577,
		},
		cli.StringFlag{
			EnvVar: "GIT_NOTIFY_REDIS_HOST",
			Name:   "redis_host",
			Usage:  "redis服务器主机",
			Value:  "localhost",
		},
		cli.IntFlag{
			EnvVar: "GIT_NOTIFY_REDIS_PORT",
			Name:   "redis_port",
			Usage:  "redis服务器端口",
			Value:  6379,
		},
		cli.IntFlag{
			EnvVar: "GIT_NOTIFY_REDIS_DB",
			Name:   "redis_db",
			Usage:  "redis服务器数据库",
			Value:  9,
		},
		cli.StringFlag{
			EnvVar: "GIT_NOTIFY_REDIS_PASSWORD",
			Name:   "redis_password",
			Usage:  "redis服务器密码",
			Value:  "",
		},
	},
		GetCorpFlags(),
		GetGitlabFlags(),
	),
}

/**
初始化Redis
*/
func initRedis(c *cli.Context) *redis.Client{
	log.Print("start to init redis")
	redisClient := redis.NewClient(&redis.Options{
		Addr:     c.String("redis_host") + ":" + strconv.Itoa(c.Int("redis_port")),
		Password: c.String("redis_password"),
		DB:       c.Int("redis_db"),
	})
	if err := redisClient.Ping().Err(); err != nil {
		log.Fatal("failed to connect redis")
	}
	log.Print("init redis success")
	return redisClient
}

func server(c *cli.Context) error {
	utils.InitLog(c)
	redisClient := initRedis(c)

	remoteApi, err := setupRemote(c)
	if err != nil {
		return err
	}

	r := kelly.NewClassic()

	wxapi := qywechat.InitMessageApi(c, redisClient)
	gateway.InitializeApiRoutes(r.Group("/gateway"),remoteApi,wxapi)
	r.Run(c.String("host") + ":" + strconv.Itoa(c.Int("port")))
	return nil
}
