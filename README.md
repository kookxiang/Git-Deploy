你好！
很冒昧用这样的方式来和你沟通，如有打扰请忽略我的提交哈。我是光年实验室（gnlab.com）的HR，在招Golang开发工程师，我们是一个技术型团队，技术氛围非常好。全职和兼职都可以，不过最好是全职，工作地点杭州。
我们公司是做流量增长的，Golang负责开发SAAS平台的应用，我们做的很多应用是全新的，工作非常有挑战也很有意思，是国内很多大厂的顾问。
如果有兴趣的话加我微信：13515810775  ，也可以访问 https://gnlab.com/，联系客服转发给HR。
# Git-Deploy
A tiny program help you deploy your code via git web hook

## What this program do
```shell
cd /path/to/your/repository
git reset --hard HEAD
git pull
deploy.sh    # Only if there is a deploy.sh file
```

## Usage
First, build this tiny program via `go build`

Then, config your nginx like this:
```
    location /Admin/Deploy/YOUR-DEPLOY-KEY-HERE/ {
        proxy_pass  http://127.0.0.1:4321/path/to/your/repository;
    }
```

I suggest you to **add that deploy key** in path and **use https request** to prevent abuse.

You may notice the path in your request is exactly the path of repository.
So **please keep this little program behind firewall**, use nginx for proxy instead.
