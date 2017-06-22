项目 <a href="{{ .Project.Homepage }}">{{ .Project.Name }}</a>【<a href="{{ .Project.Homepage }}">{{ .Project.Homepage }}</a>】 有新的提交（PUSH）。
提交者： <a href="#">{{ .User.Name }}</a>【{{ .User.Email }}】
Commit数量： <a href="#">{{ .Commit.TotalCommitsCount }}</a>
分支：<a href="#">{{ .Commit.Ref }}</a>
Sha：<a href="#">{{ .Commit.CheckoutSha }}</a>