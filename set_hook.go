package main

import (
	"encoding/json"
	"github.com/codegangsta/cli"
	"log"
	"github.com/qjw/git-notify/remote"
	"strings"
	"fmt"
)

var hookCmd = cli.Command{
	Name:  "hook",
	Usage: "给每个项目设置hook，触发消息到本服务器",
	Action: func(c *cli.Context) {
		if err := hook(c); err != nil {
			log.Fatal(err)
		}
	},
	Flags: CombineSliceArray([]cli.Flag{
		cli.StringFlag{
			EnvVar: "GIT_NOTIFY_CALLBACK_URL",
			Name:   "callback_url",
			Usage:  "回调地址",
			Value:  "",
		},
		cli.StringFlag{
			EnvVar: "GIT_NOTIFY_EVENT_LIST",
			Name:   "event_list",
			Usage:  "需要关注的时间，note|issue|push|merge_request 四种，用|分隔",
			Value:  "merge_request",
		},
		cli.BoolFlag{
			EnvVar: "GIT_NOTIFY_FORCE",
			Name:   "force",
			Usage:  "强制更新hook，默认已经存在callback的hook就不再作处理",
		},
	},
		GetGitlabFlags(),
	),
}

func hook(c *cli.Context) error {
	remoteApi, err := setupRemote(c)
	if err != nil {
		return err
	}

	projs, err := remoteApi.GetProjects()
	if err != nil{
		return err
	}

	data, err := json.MarshalIndent(projs, "", " ")
	if err != nil{
		return err
	}

	callback_url := c.String("callback_url")
	event_list := strings.ToLower(c.String("event_list"))
	events := strings.Split(event_list,"|")
	if len(events) == 0{
		return fmt.Errorf("invalid evnet list %s",event_list)
	}

	eventsFlags := map[string]bool{
		"note":false,
		"issue":false,
		"push":false,
		"merge_request":false,
	}
	for k, _ := range eventsFlags {
		for _,v := range events{
			if k == v{
				eventsFlags[k] = true
				break
			}
		}
	}

	for _,v := range projs{
		if v.ID > 0 && len(v.Name) > 0{
			hooks,_ := remoteApi.GetHooks(v.ID)
			data, _ = json.MarshalIndent(hooks, "", " ")
			log.Printf("project %s\n%s\n\n",v.NameWithNamespace,string(data))

			var existFlag bool = false
			for _,h := range hooks{
				if h.Url == callback_url{
					if !c.Bool("force"){
						existFlag = true
					}else{
						if err := remoteApi.DeleteHook(h.ProjectID,h.ID);err != nil{
							return err
						}
					}
					break
				}
			}

			if !existFlag{
				err = remoteApi.AddHook(v.ID,&remote.ProjHookParam{
					Url:callback_url,
					PushEvents:eventsFlags["push"],
					IssuesEvents:eventsFlags["issue"],
					MergeRequestsEvents:eventsFlags["merge_request"],
					NoteEvents:eventsFlags["note"],
					TagPushEvents:false,
					JobEvents:false,
					PipelineEvents:false,
					WikiEvents:false,
				})
				if err != nil{
					log.Print(err.Error())
				}
			}
		}
	}
	return nil
}
