# DevKit


## Application Template

Go, Python, Spring, ASP.NET, Angular, React,  NGINX, ...

```shell
$ devkit template golang
> App Name: demo

$ ls demo           
Dockerfile    chart         go.mod        main.go       public        skaffold.yaml

```


## Local Service Catalog

- Databases: PostgreSQL, MariaDB, Redis, Elasticsearch, ...
- Messaging: ActiveMQ, RabbitMQ, Kafka, ...
- Tools: Jenkins, SonarQube, Mailpit, ...
- Storage: MinIO, Vault, Artifactory, Nexus, ...

### Create instance

```shell
$ devkit postgres create

HAPPY_MAHAVIRA (f046e1237937)

ENDPOINT        TARGET                
localhost:55095 tcp://172.17.0.5:5432

DESCRIPTION     VALUE                                                               
Database        db                                                                 
Username        postgres                                                           
Password        r02ZxXR1E5                                                         
URI             postgresql://postgres:r02ZxXR1E5@localhost:55095/db?sslmode=disable
```

### List instances

```shell
$ devkit postgres list

happy_mahavira
```

### Display info

```shell
$ devkit postgres info
> happy_mahavira

HAPPY_MAHAVIRA (f046e1237937)

ENDPOINT        TARGET                
localhost:55095 tcp://172.17.0.5:5432

DESCRIPTION     VALUE                                                               
Database        db                                                                 
Username        postgres                                                           
Password        r02ZxXR1E5                                                         
URI             postgresql://postgres:r02ZxXR1E5@localhost:55095/db?sslmode=disable
```

### Follow logs

```shell
$ devkit postgres logs
> happy_mahavira

2024-07-14 12:09:29.661 UTC [1] LOG:  database system is ready to accept connections
```

### Start connected client

```shell
$ devkit postgres cli
> happy_mahavira

psql (16.3 (Debian 16.3-1.pgdg120+1))
Type "help" for help.

db=#
```

### Open shell

```shell
$ devkit postgres shell
> happy_mahavira

root@f046e1237937:/#
```

### Delete instance

```shell
$ devkit postgres delete
> happy_mahavira
```


## Analze Repository

### Static Analysis (using [Semgrep](https://semgrep.dev))

```shell
$ devkit sast
```

### Vulnerability Scanning (using [Trivy](https://github.com/aquasecurity/trivy))

```shell
$ devkit scan
```

### Lines of Code (using [cloc](https://github.com/AlDanial/cloc))

```shell
$ devkit cloc
```

### List big blobs

```shell
$ devkit git blobs
```

### Find leaks (using [Gitleaks](https://github.com/gitleaks/gitleaks))

```shell
$ devkit git leaks
```

### Delete file in history

```shell
$ devkit git purge /path/to/file1 /path/to/file2
```


## Analyze Images

### Inspect (using [Whaler](https://github.com/P3GLEG/Whaler))

```shell
$ devkit image inspect --image ubuntu
```

### Dockerfile Linting (using [Dockle](https://github.com/goodwithtech/dockle))

```shell
$ devkit image lint --image ubuntu
```

### Vulnerability Scanning (using [Trivy](https://github.com/aquasecurity/trivy))

```shell
$ devkit image scan --image ubuntu
```

### Show Bill of Material (using [Syft](https://github.com/anchore/syft))

```shell
$ devkit image bom --image ubuntu
```

### TUI Browser (using [Dive](https://github.com/wagoodman/dive))

```shell
$ devkit image browse --image ubuntu
```

## Utilities

### Local Web IDE

```shell
$ devkit code
> default

NAME    VALUE                 
URL     http://localhost:3000

Forward /Users/User/Projects to /workspace
```

### Local Web Server

Simple File Server with CORS and optional SPA support

```shell
$ devkit server [--port 3000] [--spa] [--index index.html]
```

### Local Proxy

Test your application for proxy support

```shell
$ devkit proxy [--port 3128] [--user username] [--password password]
```


## Install

### Requirements

- Running [Docker](https://docs.docker.com/get-docker/), [Podman](https://podman-desktop.io), [Rancher](https://rancherdesktop.io) or [Lima VM](https://lima-vm.io/docs/examples/)

#### MacOS / Linux

[Homebrew](https://brew.sh)

```
brew install adrianliechti/tap/devkit
```

#### Windows

[Scoop](https://scoop.sh)

```shell
scoop bucket add adrianliechti https://github.com/adrianliechti/scoop-bucket
scoop install adrianliechti/devkit
```