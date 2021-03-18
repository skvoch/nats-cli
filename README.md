# Command line interface for NATS Streaming
> :warning: **All versions have MacOS binaries only**: If you want to use nats-cli on other systems, you should to build the binaries manually

# Features
#### Templates
* Publishing 
* Subscribing
* Templates this is a useful tool for saving your NATS settings, and reuse them with other commands.
# Usage
```bash
Usage:
  nats-cli [command]

Available Commands:
  help        Help about any command
  publish     Publish to subject
  subscribe   Subscribe to subject
  template    Manage templates

Flags:
      --config string   config file (default is $HOME/.nats-cli.yaml)
  -h, --help            help for nats-cli
  -t, --toggle          Help message for toggle

Use "nats-cli [command] --help" for more information about a command.
```

# Subscribe example
```bash
nats-cli sub -a your-nats-server -c nats-cluster-id -s subject -d 24h
```

# Publish example
```bash
nats-cli pub -a your-nats-server -c nats-cluster-id -s subject -m '{"json":"here"}'

```
# Templates example
#### List
```bash
nats-cli template list
```
#### Create
```bash
nats-cli template create -a your-nats-server -c nats-cluster-id  -s subject -n template-name
```
#### Remove
```bash
nats-cli tpl remove -n template-name
```
#### Usage
```bash
nats-cli sub tpl -n your-template-name -d 2h
```
```bash
nats-cli pub tpl -n your-template-name -m '{"json":"here"}'
```
