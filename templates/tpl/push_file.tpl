项目 {{ .Project.Name }}【{{ .Project.Homepage }}】 有新的提交（PUSH）。
提交者： {{ .User.Name }}【{{ .User.Email }}】
Commit数量： {{ .Commit.TotalCommitsCount }}
分支：{{ .Commit.Ref }}
Sha：{{ .Commit.CheckoutSha }}
{{if gt .Commit.TotalCommitsCount 0}}
{{range $index, $element := .Commits }}
Commit: {{$element.ID}}
  作者： {{ $element.Author.Name }}【{{ $element.Author.Email }}】
  时间：{{ $element.Timestamp }}
  地址： {{ $element.Url }}
  备注： {{ $element.Message }}
{{end}}
{{end}}