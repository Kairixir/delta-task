# delta-task

## Description

Exemplar CI/CD task for Delta team interview.

## Requirements

- k8s cluster with CoreDNS able to access the internet
- argoCD CLI
- docker

## Docker build instructions

1. Build app

```bash
docker build -t hello-world-go .
```

2. Run container

```bash
docker run -p 8080:8080 hello-world-go
```

## Setup Minikube

### Setup coreDNS

To resolve correctly in minikube you need to setup external DNS resolver for coreDNS.

Open coreDNS'es configmap:

```bash
kubectl edit configmap/coredns -n kube-system
```

and add Google's DNS server (or any you prefer) to the `forward` section
instead of the present `resolv.conf`:

```yaml
forward . 8.8.8.8 8.8.4.4 1.1.1.1 {
max_concurrent 1000
}
# forward . /etc/resolv.conf {
#    max_concurrent 1000
# }
```

### Test deployed cluster

#### For Linux

Either add `minikube ip` output to `/etc/hosts`:

```bash
IP=$(minikube ip) sudo bash -c "echo \"$IP hello-world.delta # delta hw task minikube redirect\" >> /etc/hosts"
```

and access in browser, or add custom DNS resolve record to `curl` command:

```bash
curl --resolve "hello-world.delta:80:$( minikube ip )" -i http://hello-world.delta
```

#### For MacOS

The step 4 explains [here](https://kubernetes.io/docs/tasks/access-application-cluster/ingress-minikube/#create-an-ingress) how to run on MacOS.

Instructions, not tested:

```bash
# In new terminal
minikube tunnel

# Curl request
curl --resolve "hello-world.delta:80:127.0.0.1" -i http://hello-world.delta
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

### CD in argo

First I tried to use syncOptions in argo to completely rebuild and restart the
deployment. I did it mainly because it worked as a CLI command and I was running
out of time. I did not manage to make it work (since I believe CLI argocd works
differently then using syncOptions for automatic sync).

After looking into the issue I found two possible solutions:

1. Use [Argocd Image Updater](https://argocd-image-updater.readthedocs.io/en/stable/) to keep track of docker container versions in the registry
   and update the application whenever there is an image with new version.

2. Update Helm chart versions which should (?) trigger the redeployment within the k8s
   cluster.

3. Update Helm chart image tags and use semantic versioning when deploying containers
   (ie. no latest)

I am inclined to do solution 2. It seems like the best option because it preserves GitOps (ie.
git is the single source of truth) and would require minimal changes during deploy.
The con is I am not certain changing the version in Chart.yaml would trigger redeployment
as I expect it.

Solution 1 seems like an interesting solution due to the tools' documentation; however,
it breaks the principle of single source of truth in git. It depends on Container registry,
not git. At the same time the authors of the tool do not recommend using it in production
environments and

Solution 3 seems like it should always work, but it seems hardest to maintain and implement.

Probably would test the solution 2 in local env, and if it would not work, I would look more
into the Image Updater documentation. With more information I would then either use the tool
or try to build on Solution 1 and try to use templating to implement solution 3.

### Readiness and liveness probes

Now the readiness and liveness probes pollute log. I would definitely change and improve
them so the project has better monitoring and observability.

### Deploy branch

I don't like the way merge to deploy adds new commits. At first I thought it
was a good idea to track when changes are merged to the deploy branch and keep
main branch updated with deploy. However, this causes the remote update on main
branch. This breaks my local git repo.

I cannot make changes in the main branch, since the CI pipeline adds new
commit, and when I try to reconcile the changes during git pull, git refuses to
connect to any remote source for a few minutes.

Would try a `rebase` strategy instead of merge and see the results.

### HTTPS

The app is missing https.

### GitHub Actions

Even though I have used GitHub actions for the CI/CD pipeline, I would try to avoid them
in production. I would prefer to keep my pipelines tool-agnostic to enable easy migration
in case of dissatisfaction with the tool.

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
  - Minikube has issues with DNS:
    - [Guide to migrate to CoreDNS](https://kubernetes.io/docs/tasks/administer-cluster/coredns/)
    - [CoreDNS installation](https://coredns.io/manual/installation/)

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
