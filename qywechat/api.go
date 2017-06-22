package qywechat

import (
	"log"
	"github.com/codegangsta/cli"
	"gopkg.in/redis.v5"
	"strings"
	"github.com/qjw/git-notify/remote"
	"github.com/qjw/git-notify/utils"
	"github.com/qjw/go-wx-sdk/corp"
	"github.com/qjw/go-wx-sdk/cache"
)

type MessageApi interface {
	SendMessage(*remote.HookNotify)
}

type QywechatApi struct {
	WxApi *corp.CorpApi
	AppID int64
	DepID int64
}

func (this QywechatApi) SendMessage(msg *remote.HookNotify){
	if msg.ObjectKind == "push"{
		msgContent := utils.ExcuteTemplate("push.tpl",msg)
		this.sendTextMsg(msgContent)
		msgContent = utils.ExcuteTemplate("push_file.tpl",msg)
		this.sendFileMsg(msgContent,"push.txt")
	}else if msg.ObjectKind == "issue"{
		msgContent := utils.ExcuteTemplate("issue.tpl",msg)
		this.sendTextMsg(msgContent)
	}else if msg.ObjectKind == "merge_request"{
		msgContent := utils.ExcuteTemplate("merge_request.tpl",msg)
		this.sendTextMsg(msgContent)
	}else if msg.ObjectKind == "note"{
		var msgContent string
		if msg.MergeRequest != nil{
			msgContent = utils.ExcuteTemplate("note_merge_request.tpl",msg)
		}else if msg.Issue != nil{
			msgContent = utils.ExcuteTemplate("note_issue.tpl",msg)
		}else{
			log.Print("invalid note msg")
			return
		}
		this.sendTextMsg(msgContent)
	}
}

func (this QywechatApi) sendTextMsg(msg string){
	log.Printf("send msg to wx %s",msg)
	res,err := this.WxApi.SendTextMsg([]string{},[]int64{this.DepID},[]int64{},this.AppID,msg)
	if err != nil{
		log.Printf("send msg fail %s",err.Error())
	}else if res.ErrCode != 0{
		log.Printf("send msg fail %d|%s", res.ErrCode, res.ErrMsg)
	}
}

func (this QywechatApi) sendFileMsg(msg string,filename string){
	res,err := this.WxApi.UploadTmpMedia(strings.NewReader(msg),filename,"file")
	if err != nil{
		log.Printf("upload file %s fail %s",filename,err.Error())
	}else if res.ErrCode != 0{
		log.Printf("upload file %s fail %d|%s",filename, res.ErrCode, res.ErrMsg)
	}

	res2,err := this.WxApi.SendFileMsg([]string{},[]int64{this.DepID},[]int64{},this.AppID,res.MediaID)
	if err != nil{
		log.Printf("send file fail %s",err.Error())
	}else if res2.ErrCode != 0{
		log.Printf("send file fail %d|%s", res2.ErrCode, res2.ErrMsg)
	}
}

func InitMessageApi(c *cli.Context,redisClient *redis.Client) MessageApi{
	corpID := c.String("corp_id")
	corpSecret := c.String("corp_secret")
	corpAppID := c.Int64("corp_appid")
	receiveDepartment := c.String("receive_department")
	var receiveDepartmentID int64 = 0
	if len(corpID) == 0 || len(corpSecret) == 0 ||
		len(receiveDepartment) == 0 || corpAppID == 0{
		panic("invalid parameter")
	}

	context := corp.NewContext(
		&corp.Config{
			CorpID:     corpID,
			CorpSecret: corpSecret,
			Tag:        "gitlab",
		},
		cache.NewCache(redisClient),
	)
	wxApi := corp.NewCorpApi(context)

	dps,_ := wxApi.GetDepartments(nil)
	if dps.ErrCode != 0{
		panic("can NOT get departments")
	}
	for _,v := range dps.Departments{
		if v.Name == receiveDepartment{
			log.Printf("%d",v.ID)
			receiveDepartmentID = *v.ID
			break
		}
	}

	if receiveDepartmentID == 0{
		panic("can NOT find department named " + receiveDepartment)
	}

	return &QywechatApi{
		WxApi:wxApi,
		AppID:corpAppID,
		DepID:receiveDepartmentID,
	}
}