# delta-task

## Description

Exemplar CI/CD task for Delta team interview.

## Build instructions

1. Build app

```bash
docker build -t hello-world-go .
```

2. Run container

```bash
docker run -p 8080:8080 hello-world-go
```

## Design decision

Should the tests be handled through Dockerfile, or through CI/CD tool?

Dockerfile option will make the testing CI/CD tool agnostic, but it will make
the Dockerfile much harder to maintain.

Using the CI/CD tool will be easier to maintain over the long term, but it will
be much harder to migrate to another CI/CD tool, because I will need to rely on
specific features of the tool—ie. GitHub actions.

Since my goal is to show my skills with CI/CD I will go through the GitHub actions route.
However, I currently believe there will be a redundant build of the application for unit tests.
Also, I suspect Go is using advanced strategies to recompile only relevant parts of code.
Since I will be building the app in virtual env pipeline, the compiler will not be able to use
these strategies, thus making the compilation slower and more resource intensive.

For my small application this is not a problem, but for a larger or growing application this could
grow into a huge problem as the application becomes larger. Due to time constraints I deem ignoring this
the best path forward.

If it were not for the testing of my CI/CD skills I would have gone with the Dockerfile option if
restricted similarly. Especially since I plan to migrate this project to ArgoCD when I finish it.

## Test if pod runs in Minikube

### For Linux

Either add `minikube ip` output to `/etc/hosts`:

```bash
IP=$(minikube ip) sudo bash -c "echo \"$IP hello-world.delta # delta hw task minikube redirect\" >> /etc/hosts"
```

and access in browser, or add custom DNS resolve record to `curl` command:

```bash
curl --resolve "hello-world.delta:80:$( minikube ip )" -i http://hello-world.delta
```

### For MacOS

The step 4 explains [here](https://kubernetes.io/docs/tasks/access-application-cluster/ingress-minikube/#create-an-ingress) how to run on MacOS.

Instructions, not tested:

```bash
# In new terminal
minikube tunnel

# Curl request
curl --resolve "hello-world.delta:80:127.0.0.1" -i http://hello-world.delta
```

## Comments

### I am pushing straight to master and using a lot of comments

I see this as a mistake, but it is time efficient. I prefer to use branches and
PRs for code changes and then squash all the commits into one with sensible
message to not pollute the main's commit history. I would never do this on
established codebase or on a project I plan to scale.

### Caching

The caching of the build process is terribly optimized. For production pipeline
I would focus much more time on this. For example container rebuilding, common
cache for Docker layers, rebuilding Go code efficiently with respect to what
Go compiler allows,

## Sources

The sources document my path through the task:

- Kagi as search engine
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
  - [Helm test](https://helm.sh/docs/topics/chart_tests/)

- Minikube:

  - CLI help
  - [Ingress in minikube](https://kubernetes.io/docs/tasks/access-application-cluster/ingress-minikube/)
  - [Ingress DNS](https://minikube.sigs.k8s.io/docs/handbook/addons/ingress-dns/#Linux)
    - Add minikube as DNS server for domain names to work correctly without tunnel

- Kubernetes:

  - kubectl CLI help
  - [Services](https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-types)
  - [Ingress](https://kubernetes.io/docs/concepts/services-networking/ingress/)

- GitHub actions:

  - [docs](https://docs.github.com/en/actions)
  - [Go starter project](https://github.com/actions/starter-workflows/blob/main/ci/go.yml)
  - [Build & Test Go docs](https://docs.github.com/en/actions/use-cases-and-examples/building-and-testing/building-and-testing-go)
  - Github docs for [Docker container action](https://docs.github.com/en/actions/sharing-automations/creating-actions/creating-a-docker-container-action) and [Dockerfile support](https://docs.github.com/en/actions/sharing-automations/creating-actions/dockerfile-support-for-github-actions)
  - [Docker build image starter action](https://github.com/actions/starter-workflows/blob/main/ci/docker-image.yml)
  - [Docker publish starter action](https://github.com/actions/starter-workflows/blob/main/ci/docker-publish.yml)
  - [Setup Go action](https://github.com/actions/setup-go/tree/main)

- ArgoCD
  - [Official docs](https://argo-cd.readthedocs.io/en/stable/getting_started/)
  - [ArgoCD Helm](https://argo-cd.readthedocs.io/en/stable/user-guide/helm/)
