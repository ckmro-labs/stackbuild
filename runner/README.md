# go 1.11+ use docker api

* go-docker@v1.0.0, not master.
* 可以将GO111MODULE暂时设置为on，进行build, go get docker.io/go-docker 会检测出来v1.0.0,
* 将GOPATH/pkg/下的对应目录去除版本号拷贝至github.com or docker.io目录
* 恢复GO111MODULE=off
* rebuild