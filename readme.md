# jtnctl - Simplify kubectl Commands

`jtnctl` is a command-line interface (CLI) tool that simplifies the execution of `kubectl` commands, such as 'logs', 'port-forward', and 'exec'  It allows you to streamline Kubernetes operations with ease.
## Requirements

Before using `jtnctl`, make sure you have the following prerequisites installed on your system:

- **kubectl:** The Kubernetes command-line tool (`kubectl`) is required to interact with Kubernetes clusters. You can download and install `kubectl` by following the [official Kubernetes documentation](https://kubernetes.io/docs/tasks/tools/install-kubectl/)

## Installation

You can install `jtnctl` using the following methods:

### Option 1: Binary Releases (Cross-platform)

Download the latest release for your platform from the [Releases](https://github.com/zakyyudha/jtnctl/releases) page.

#### MacOS
```bash
~ $ wget https://github.com/zakyyudha/jtnctl/releases/download/v1.0.0/jtnctl_v1.0.0_darwin -O jtnctl
~ $ chmod +x jtnctl
~ $ sudo mv jtnctl /usr/local/bin/
```

#### Linux
```bash
~ $ wget https://github.com/zakyyudha/jtnctl/releases/download/v1.0.0/jtnctl_v1.0.0_linux_amd64 -O jtnctl
~ $ chmod +x jtnctl
~ $ sudo mv jtnctl /usr/local/bin/
```

### Option 2: Building from Source
Clone the repository and build jtnctl:
```bash
~ $ git clone https://github.com/zakyyudha/jtnctl.git
~ $ cd jtnctl
~ $ go build -o jtnctl main.go
```

## Usage
### Display available commands and options
```bash
~ $ jtnctl --help

jtnctl is a CLI tool for simplifying kubectl commands

Usage:
  jtnctl [command]

Available Commands:
  config       Manage configuration settings
  exec         Execute a command in an active service
  help         Help about any command
  logs         View logs of a Kubernetes pod
  port-forward Forward ports for a Kubernetes service

Flags:
  -h, --help   help for jtnctl

Use "jtnctl [command] --help" for more information about a command.
```

### Set the default namespace (e.g., tds-stage)
```bash
~ $ jtnctl config set namespace=tds-stage
```
![config-set](./docs/config-set.gif)

### Simplify 'kubectl port-forward' command
```bash
~ $ jtnctl port-forward [service-name] [destination:source]
```
![config-set](./docs/port-forward.gif)

### Simplify 'kubectl logs' command
```bash
~ $ jtnctl logs [pod-name] [--follow] [--tail=10]
```
![config-set](./docs/logs.gif)

### Simplify 'kubectl exec' command
```bash
~ $ jtnctl exec [pod-name] -it -- /bin/bash
```
![config-set](./docs/exec.gif)

## Configuration
`jtnctl` allows you to `config` a default namespace using the set command. For example:
```bash
~ $ jtnctl config set namespace=tds-stage
```

The configuration is stored in a `~/.config/jtnctl/config.yaml` file. You can modify it directly or use the set command to update values
