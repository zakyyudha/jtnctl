# jtnctl - Simplify kubectl Commands

`jtnctl` is a command-line interface (CLI) tool that simplifies the execution of `kubectl` commands, such as 'logs', 'port-forward', and 'exec'  It allows you to streamline Kubernetes operations with ease.

## Installation

You can install `jtnctl` using the following methods:

### Option 1: Binary Releases (Cross-platform)

Download the latest release for your platform from the [Releases](https://github.com/zakyyudha/jtnctl/releases) page.

```bash
# Replace with the actual release version and your OS/architecture
wget https://github.com/zakyyudha/jtnctl/releases/download/v1.0.0/jtnctl_v1.0.0_linux_amd64 -O jtnctl
chmod +x jtnctl
sudo mv jtnctl /usr/local/bin/
```

### Option 2: Building from Source
Clone the repository and build jtnctl:
```bash
git clone https://github.com/zakyyudha/jtnctl.git
cd jtnctl
go build -o jtnctl main.go
```

## Usage
### Display available commands and options
```bash
jtnctl --help
```

### Set the default namespace (e.g., tds-stage)
```bash
jtnctl set namespace=tds-stage
```

### Simplify 'kubectl logs' command
```bash
jtnctl logs [pod-name] [--follow] [--tail=10]
```

### Simplify 'kubectl exec' command
```bash
jtnctl exec [pod-name] -it -- /bin/bash
```

### Simplify 'kubectl port-forward' command
```bash
jtnctl port-forward [service-name] [destination:source]
```

## Configuration
`jtnctl` allows you to `config` a default namespace using the set command. For example:
```bash
jtnctl config namespace=tds-stage
```

The configuration is stored in a `~/.config/jtnctl/config.yaml` file. You can modify it directly or use the set command to update values