项目 <a href="{{ .Project.Homepage }}">{{ .Project.Name }}</a>【<a href="{{ .Project.Homepage }}">{{ .Project.Homepage }}</a>】 有问题（ISSUE）变化。
提交人： <a href="{{ .Issue.Url }}">{{ .Issue.Author.Name }}</a>【{{ .Issue.Author.Email }}】
分配给： {{if gt .Issue.AssigneeID 0}}<a href="{{ .Issue.Url }}">{{ .Issue.Assignee.Name }}</a>【{{ .Issue.Assignee.Email }}】{{end}}
操作： <a href="{{ .Issue.Url }}">{{ .Issue.Action }}</a>
状态： <a href="{{ .Issue.Url }}">{{ .Issue.State }}</a>
标题： <a href="{{ .Issue.Url }}">{{ .Issue.Title }}</a>
正文： <a href="{{ .Issue.Url }}">{{ .Issue.Description }}</a>