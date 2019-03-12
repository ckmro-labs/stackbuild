# About StackBuild

* Containerization Env
* Continuous integration
* Auto Deploy 
* Kubernets & Docker Swarm

# MODULE

* Vcs: 代码仓库授权后访问仓库所需接口，CI部份包含WebHook
* Pipeline: 管道集成构建模块，CI/CD核心过程，运行于容器
* Builder: 处理构建结果: 日志/发布镜像等
* Runner: 构建容器调度
* Deploy: 部署模块，目标对象为K8s/Swarm，使用YAML或Helm
* Api: platform apis.