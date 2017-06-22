package gitlab

import (
	"errors"
	"fmt"
	"github.com/qjw/git-notify/remote"
	"github.com/qjw/go-wx-sdk/utils"
	"net"
	"net/url"
)

const (
	apiPrefix = "/api/v4%s?private_token=%s"
)

const (
	getProjects = "/projects"
	getHooks    = "/projects/%d/hooks"
	addHook     = "/projects/%d/hooks"
	getUser     = "/users/%d"
	deleteHook  = "/projects/%d/hooks/%d"
)

type CommonError struct {
	Message string `json:"message,omitempty"`
}

// Opts defines configuration options.
type Opts struct {
	URL   string
	Token string
}

// New returns a Remote implementation that integrates with Gitlab, an open
// source Git service. See https://gitlab.com
func New(opts Opts) (remote.Remote, error) {
	url, err := url.Parse(opts.URL)
	if err != nil {
		return nil, err
	}
	host, _, err := net.SplitHostPort(url.Host)
	if err == nil {
		url.Host = host
	}

	return &Gitlab{
		Url:   opts.URL + apiPrefix,
		Token: opts.Token,
	}, nil
}

type Gitlab struct {
	utils.ApiBase
	Url   string
	Token string
}

func (this Gitlab) GetProjects() ([]remote.Project, error) {
	var res []remote.Project
	if err := this.DoGet(this.Url, &res, getProjects, this.Token); err == nil {
		return res, nil
	} else {
		return []remote.Project{}, err
	}
}

func (this Gitlab) GetHooks(projID int64) ([]remote.ProjHook, error) {
	var res []remote.ProjHook
	if err := this.DoGet(this.Url, &res, fmt.Sprintf(getHooks, projID), this.Token); err == nil {
		return res, nil
	} else {
		return []remote.ProjHook{}, err
	}
}

func (this Gitlab) AddHook(projID int64, param *remote.ProjHookParam) error {
	var res CommonError
	if err := this.DoPostObject(this.Url, param, &res, fmt.Sprintf(addHook, projID), this.Token); err == nil {
		if len(res.Message) > 0 {
			return errors.New(res.Message)
		} else {
			return nil
		}
	} else {
		return err
	}
}

func (this Gitlab) DeleteHook(projID int64, hookID int64) error {
	var res struct{}
	if err := this.DoDelete(this.Url, &res, fmt.Sprintf(deleteHook, projID, hookID), this.Token); err == nil {
		return nil
	} else {
		return err
	}
}

func (this Gitlab) GetUser(userID int64) (*remote.User, error) {
	var res remote.User
	if err := this.DoGet(this.Url, &res, fmt.Sprintf(getUser, userID), this.Token); err == nil {
		return &res, nil
	} else {
		return nil, err
	}
}
