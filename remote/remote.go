package remote

type Project struct {
	ID                int64  `json:"id"`
	Description       string `json:"description,omitempty"`
	DefaultBranch     string `json:"default_branch,omitempty"`
	Name              string `json:"name"`
	NameWithNamespace string `json:"name_with_namespace,omitempty"`
	CreatedAt         string `json:"created_at,omitempty"`
	AvatarUrl         string `json:"avatar_url,omitempty"`
}

type ProjHook struct {
	ID        int64  `json:"id"`
	ProjectID int64  `json:"project_id"`
	Url       string `json:"url"`
	CreatedAt string `json:"created_at"`
}

// Event有默认值，不要omitempty
type ProjHookParam struct {
	Url                 string `json:"url"`
	PushEvents          bool   `json:"push_events"`
	IssuesEvents        bool   `json:"issues_events"`
	MergeRequestsEvents bool   `json:"merge_requests_events"`
	TagPushEvents       bool   `json:"tag_push_events"`
	NoteEvents          bool   `json:"note_events"`
	JobEvents           bool   `json:"job_events"`
	PipelineEvents      bool   `json:"pipeline_events"`
	WikiEvents          bool   `json:"wiki_events"`
}

type User struct {
	ID        int64  `json:"id"`
	Name      string `json:"name,omitempty"`
	Email     string `json:"email,omitempty"`
	Username  string `json:"username,omitempty"`
	AvatarUrl string `json:"avatar_url,omitempty"`
}

type Remote interface {
	GetProjects() ([]Project, error)
	GetHooks(projID int64) ([]ProjHook, error)
	AddHook(projID int64, param *ProjHookParam) error
	DeleteHook(projID int64, hookID int64) error
	GetUser(userID int64) (*User, error)
}

type HookNotify struct {
	ObjectKind string `json:"object_kind"`

	// push event
	EventName         string `json:"event_name,omitempty"`
	Ref               string `json:"ref,omitempty"`
	CheckoutSha       string `json:"checkout_sha,omitempty"`
	TotalCommitsCount int    `json:"total_commits_count,omitempty"`
	Name              string `json:"user_name,omitempty"`
	Email             string `json:"user_email,omitempty"`
	AvatarUrl         string `json:"user_avatar,omitempty"`

	User             User       `json:"user"`
	Project          ProjectObj `json:"project"`
	ObjectAttributes *struct {
		Title       string `json:"title,omitempty"`
		Description string `json:"description,omitempty"`
		AssigneeID  int64  `json:"assignee_id,omitempty"`
		AuthorID    int64  `json:"author_id,omitempty"`

		State  string `json:"state,omitempty"`
		Url    string `json:"url,omitempty"`
		Action string `json:"action,omitempty"`

		// note
		DiscussionID string `json:"discussion_id,omitempty"`
		Note         string `json:"note,omitempty"`
		NoteableType string `json:"noteable_type,omitempty"`

		// merge_request
		TargetBranch string      `json:"target_branch,omitempty"`
		SourceBranch string      `json:"source_branch,omitempty"`
		MergeStatus  string      `json:"merge_status,omitempty"`
		LastCommit   *CommitObj  `json:"last_commit,omitempty"`
		Source       *ProjectObj `json:"source,omitempty"`
		Target       *ProjectObj `json:"target,omitempty"`
	} `json:"object_attributes,omitempty"`

	Issue        *IssueObj        `json:"issue,omitempty"`
	Note         *NoteObj         `json:"note,omitempty"`
	MergeRequest *MergeRequestObj `json:"merge_request,omitempty"`
	// push
	Commits []struct {
		CommitObj
		Modified []string `json:"modified,omitempty"`
		Added    []string `json:"added,omitempty"`
		Removed  []string `json:"removed,omitempty"`
	} `json:"commits,omitempty"`
	Commit *CommitBriefObj `json:"commit,omitempty"`
}

type IssueObj struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	AssigneeID  int64  `json:"assignee_id,omitempty"`
	Assignee    *User  `json:"assignee,omitempty"`
	AuthorID    int64  `json:"author_id,omitempty"`
	Author      *User  `json:"author,omitempty"`

	State  string `json:"state,omitempty"`
	Url    string `json:"url,omitempty"`
	Action string `json:"action,omitempty"`
}

type NoteObj struct {
	DiscussionID string `json:"discussion_id,omitempty"`
	Note         string `json:"note,omitempty"`
	NoteableType string `json:"noteable_type,omitempty"`
	Url          string `json:"url,omitempty"`
	AuthorID     int64  `json:"author_id,omitempty"`
	Author       *User  `json:"author,omitempty"`
}

type MergeRequestObj struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	AssigneeID  int64  `json:"assignee_id,omitempty"`
	Assignee    *User  `json:"assignee,omitempty"`
	AuthorID    int64  `json:"author_id,omitempty"`
	Author      *User  `json:"author,omitempty"`

	TargetBranch string      `json:"target_branch,omitempty"`
	SourceBranch string      `json:"source_branch,omitempty"`
	MergeStatus  string      `json:"merge_status,omitempty"`
	LastCommit   *CommitObj  `json:"last_commit,omitempty"`
	Source       *ProjectObj `json:"source,omitempty"`
	Target       *ProjectObj `json:"target,omitempty"`

	State  string `json:"state,omitempty"`
	Url    string `json:"url,omitempty"`
	Action string `json:"action,omitempty"`
}

type CommitObj struct {
	ID        string `json:"id,omitempty"`
	Message   string `json:"message,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
	Url       string `json:"url,omitempty"`
	Author    struct {
		Name  string `json:"name,omitempty"`
		Email string `json:"email,omitempty"`
	} `json:"author,omitempty"`
}

type CommitBriefObj struct {
	EventName         string `json:"event_name,omitempty"`
	Ref               string `json:"ref,omitempty"`
	CheckoutSha       string `json:"checkout_sha,omitempty"`
	TotalCommitsCount int    `json:"total_commits_count,omitempty"`
}

type ProjectObj struct {
	Name          string `json:"name,omitempty"`
	Description   string `json:"description,omitempty,omitempty"`
	DefaultBranch string `json:"default_branch,omitempty,omitempty"`
	Namespace     string `json:"namespace,omitempty,omitempty"`
	AvatarUrl     string `json:"avatar_url,omitempty,omitempty"`
	GitSshUrl     string `json:"git_ssh_url,omitempty"`
	GitHttpUrl    string `json:"git_http_url,omitempty"`
	Url           string `json:"url,omitempty"`
	SshUrl        string `json:"ssh_url,omitempty"`
	HttpUrl       string `json:"http_url,omitempty"`
	WebUrl        string `json:"web_url,omitempty"`
	Homepage      string `json:"homepage,omitempty"`
}
