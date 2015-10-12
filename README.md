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
