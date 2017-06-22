## 接收git仓库并发送消息至企业微信

目前只实现了[gitlab](https://gitlab.com)和[企业微信](http://work.weixin.qq.com/)，可以扩展至[github](https://github.com/)，[bitbucket](https://bitbucket.org/ )等仓库，以及[slack](https://slack.com/)等企业工具

尚未支持配置hook的事件类型，目前会订阅push，merge_request,issue和note事件。

# 参数配置
``` bash
king@king:~/tmp/test$ /tmp/git_notify -h
NAME:
   git-notify - git-notify命令行工具

USAGE:
   git_notify [global options] command [command options] [arguments...]

VERSION:
   0.0.0.1

COMMANDS:
     server   启动账户服务器
     hook     给每个项目设置hook，触发消息到本服务器
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   -d, --debug    开启调试模式 [$GIT_NOTIFY_DEBUG]
   --help, -h     show help
   --version, -v  print the version

```
``` bash
king@king:~/tmp/test$ /tmp/git_notify hook -h
NAME:
   git_notify hook - 给每个项目设置hook，触发消息到本服务器

USAGE:
   git_notify hook [command options] [arguments...]

OPTIONS:
   --callback_url value    回调地址 [$GIT_NOTIFY_CALLBACK_URL]
   --gitlab                gitlab是否启用 [$GIT_NOTIFY_GITLAB]
   --gitlab-server value   gitlab仓库地址 (default: "https://gitlab.com") [$GIT_NOTIFY_GITLAB_URL]
   --gitlab-patoken value  gitlab personal access token [$GIT_NOTIFY_GITLAB_TOKEN]

```
``` bash
king@king:~/tmp/test$ /tmp/git_notify server -h
NAME:
   git_notify server - 启动账户服务器

USAGE:
   git_notify server [command options] [arguments...]

OPTIONS:
   --host value                服务器主机 (default: "localhost") [$GIT_NOTIFY_HOST]
   --port value                服务器端口 (default: 13577) [$GIT_NOTIFY_PORT]
   --redis_host value          redis服务器主机 (default: "localhost") [$GIT_NOTIFY_REDIS_HOST]
   --redis_port value          redis服务器端口 (default: 6379) [$GIT_NOTIFY_REDIS_PORT]
   --redis_db value            redis服务器数据库 (default: 9) [$GIT_NOTIFY_REDIS_DB]
   --redis_password value      redis服务器密码 [$GIT_NOTIFY_REDIS_PASSWORD]
   --corp_id value             企业微信corpid [$GIT_NOTIFY_CORP_ID]
   --corp_secret value         企业微信应用secret [$GIT_NOTIFY_CORP_SECRET]
   --corp_appid value          企业微信应用id (default: 0) [$GIT_NOTIFY_CORP_SECRET]
   --receive_department value  接收消息的部门 [$GIT_NOTIFY_CORP_SECRET]
   --gitlab                    gitlab是否启用 [$GIT_NOTIFY_GITLAB]
   --gitlab-server value       gitlab仓库地址 (default: "https://gitlab.com") [$GIT_NOTIFY_GITLAB_URL]
   --gitlab-patoken value      gitlab personal access token [$GIT_NOTIFY_GITLAB_TOKEN]
```

在OPTIONS中，左边是命令行参数，例如【--host 0.0.0.0】（中间无等号），【$GIT_NOTIFY_HOST】是环境变量参数。

Redis用于支持企业微信token的缓存

**receive_department** 指定用于接收消息的**部门名称** ，请务必确保它存在

需要确保gitlab personal access token有足够的权限
1. 设置project hook
2. 查询user（hook通知中很多人员信息都只有ID）

## 典型的用法
设置hook，若新增了project，需要重新手动运行
``` bash
git_notify hook --gitlab=true \
    --gitlab-server=http://127.0.0.1:8080 \
    --gitlab-patoken=123456789 \
    --callback_url=http://qjw.ycy.qiujinwu.com/gateway/hook
```

启动服务器（实例中开启了调试模式）
``` bash
git_notify -d server --gitlab=true \
    --gitlab-server=http://127.0.0.1:8080 \
    --gitlab-patoken=123456789 \
    --corp_id=123456789123456 \
    --corp_secret=f1a2dsf1asaadasd121dfdsadf  \
    --corp_appid=1000002 --receive_department=研发部门
```

# 通知模板
目前尚未支持模板自定义，代码在<https://github.com/qjw/git-notify/tree/master/templates/tpl>，语法是标准的golang template语法

修改之后，需要运行如下命令，重新将资源打包到go代码中
``` bash
go generate github.com/qjw/git-notify/templates/
```


# 编译
运行一下命令，在release目录生成git-notify
``` bash
make deps gen test
make build
```

编译之后，可以直接从Dockerfile创建一个很小的docker镜像
