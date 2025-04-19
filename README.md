# delta-task

## Description

Exemplar CI/CD task for Delta team interview.

## Build instructions

1. Setup `ssh-agent` on your system with key to access `Kairixir/delta-task`
2. Build app

```bash
docker build -t hello-world-go --ssh=default .
```

3. Run container

```bash
docker run -p 8080:8080 hello-world-go
```

## Sources

The sources document my path through the task:

- Perplexity `delta-task` space
- Go docs:
  - [Getting started](https://go.dev/doc/tutorial/getting-started)
  - [Create module](https://go.dev/doc/tutorial/create-module)
  - [learnxinyminutes - Go](https://learnxinyminutes.com/go/)
  - [fmt](https://pkg.go.dev/fmt@go1.24.2)
  - [log](https://pkg.go.dev/log@go1.24.2)
  - [net/http](https://pkg.go.dev/net/http@go1.24.2)
  - [net/http/httptest](https://pkg.go.dev/net/http/httptest)
  - [testing](https://pkg.go.dev/testing)
  - [os](https://pkg.go.dev/os@go1.24.2)
- [DockerHub](https://hub.docker.com/_/golang/tags?name=alpine)
- Docker docs:
  - [Manage secrets](https://docs.docker.com/build/building/secrets/)
- [My project for Dockerfile inspiration](https://github.com/Kairixir/PA234/blob/main/hw01/Dockerfile)
- GitHub docs:

  - [SSH keys](https://docs.github.com/en/authentication/connecting-to-github-with-ssh/managing-deploy-keys)
    - Annoying. GitHub cannot do Project Deploy token, only ssh key. GitLab can :/

- Helm docs:

  - [Cheatsheet](https://helm.sh/docs/intro/cheatsheet/)

- Minikube:
  - help
