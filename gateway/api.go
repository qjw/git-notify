package gateway

import (
	"net/http"
	"io/ioutil"
	"log"
	"encoding/json"
	"reflect"
	"github.com/qjw/git-notify/remote"
	"github.com/qjw/git-notify/qywechat"
	"github.com/qjw/kelly"
)

func copyObj(from interface{},to interface{}){
	fValue1 := reflect.ValueOf(from)
	if fValue1.Kind() != reflect.Ptr {
		panic("invalid varibal to sign (need to be pointer)")
	}
	fValue := fValue1.Elem()

	tValue := reflect.ValueOf(to)
	if tValue.Kind() != reflect.Ptr {
		panic("invalid varibal to sign (need to be pointer)")
	}
	tValue = tValue.Elem()

	tp := tValue.Type()
	count := tp.NumField()
	if count == 0 {
		return
	}

	for i := 0; i < count; i++ {
		field := tp.Field(i)
		fieldV := tValue.Field(i)

		if field.Anonymous {
			// todo
			continue
		}

		fv := fValue.FieldByName(field.Name)
		if fv.Kind() == reflect.Invalid{
			continue
		}
		if fv.Type().Kind() != field.Type.Kind(){
			panic("invalid cast")
		}

		fieldV.Set(fv)
		fv.Set(reflect.Zero(fv.Type()))
	}
}

func getUser(userID int64,remoteApi remote.Remote) *remote.User{
	if userID == 0{
		return nil
	}
	if user,err := remoteApi.GetUser(userID);err == nil {
		return user
	}else {
		log.Print(err.Error())
		return nil
	}
}

func fiterCommit(msg *remote.HookNotify,remoteApi remote.Remote){
	if msg.ObjectKind == "push"{
		if msg.Commit == nil{
			msg.Commit = &remote.CommitBriefObj{}
		}
		copyObj(msg,&msg.User)
		copyObj(msg,msg.Commit)
	}else if msg.ObjectKind == "merge_request"{
		if msg.MergeRequest == nil{
			msg.MergeRequest = &remote.MergeRequestObj{}
		}
		copyObj(msg.ObjectAttributes,msg.MergeRequest)
		msg.MergeRequest.Assignee = getUser(msg.MergeRequest.AssigneeID,remoteApi)
		msg.MergeRequest.Author = getUser(msg.MergeRequest.AuthorID,remoteApi)
	}else if msg.ObjectKind == "note"{
		if msg.Note == nil{
			msg.Note = &remote.NoteObj{}
		}
		copyObj(msg.ObjectAttributes,msg.Note)
		msg.Note.Author = getUser(msg.Note.AuthorID,remoteApi)
		if msg.Issue != nil{
			msg.Issue.Assignee = getUser(msg.Issue.AssigneeID,remoteApi)
			msg.Issue.Author = getUser(msg.Issue.AuthorID,remoteApi)
		}else if msg.MergeRequest != nil{
			msg.MergeRequest.Assignee = getUser(msg.MergeRequest.AssigneeID,remoteApi)
			msg.MergeRequest.Author = getUser(msg.MergeRequest.AuthorID,remoteApi)
		}
	}else if msg.ObjectKind == "issue"{
		if msg.Issue == nil{
			msg.Issue = &remote.IssueObj{}
		}
		copyObj(msg.ObjectAttributes,msg.Issue)
		msg.Issue.Assignee = getUser(msg.Issue.AssigneeID,remoteApi)
		msg.Issue.Author = getUser(msg.Issue.AuthorID,remoteApi)
	}
	msg.ObjectAttributes = nil
}

func InitializeApiRoutes(grouter kelly.Router,remoteApi remote.Remote,wxApi qywechat.MessageApi) {
	grouter.GET("/hook",func(c *kelly.Context) {
		log.Print("get")
		c.WriteString(http.StatusOK,"ok")
	})

	grouter.POST("/hook",func(c *kelly.Context) {
		rawXMLMsgBytes, _ := ioutil.ReadAll(c.Request().Body)
		log.Print(string(rawXMLMsgBytes))
		var msg remote.HookNotify
		if err := json.Unmarshal(rawXMLMsgBytes,&msg);err != nil{
			log.Print(err.Error())
		}else{
			fiterCommit(&msg,remoteApi)
			wxApi.SendMessage(&msg)
			 data,_ := json.MarshalIndent(&msg,""," ")
			 log.Print(string(data))
		}
		c.WriteString(http.StatusOK,"ok")
	})
}

