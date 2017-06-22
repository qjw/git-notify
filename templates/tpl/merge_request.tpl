项目 <a href="{{ .Project.Homepage }}">{{ .Project.Name }}</a>【<a href="{{ .Project.Homepage }}">{{ .Project.Homepage }}</a>】 有Pull Request变化。
提交人： <a href="{{ .MergeRequest.Url }}">{{ .MergeRequest.Author.Name }}</a>【{{ .MergeRequest.Author.Email }}】
分配给： {{if gt .MergeRequest.AssigneeID 0}}<a href="{{ .MergeRequest.Url }}">{{ .MergeRequest.Assignee.Name }}</a>【{{ .MergeRequest.Assignee.Email }}】{{end}}
操作： <a href="{{ .MergeRequest.Url }}">{{ .MergeRequest.Action }}</a>
状态： <a href="{{ .MergeRequest.Url }}">{{ .MergeRequest.State }}</a>
源分支： <a href="{{ .MergeRequest.Url }}">{{ .MergeRequest.SourceBranch }}</a>
目标分支： <a href="{{ .MergeRequest.Url }}">{{ .MergeRequest.TargetBranch }}</a>
标题： <a href="{{ .MergeRequest.Url }}">{{ .MergeRequest.Title }}</a>
正文： <a href="{{ .MergeRequest.Url }}">{{ .MergeRequest.Description }}</a>
