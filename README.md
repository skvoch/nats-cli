# Command line interface for NATS Streaming
# Features
#### Templates
Templates this is useful tool for saving your NATS settings, and reuse them with other commands.

# Usage
```bash
  nats-cli [command]

Available Commands:
  help        Help about any command
  publish     Publish message to subject
  subscribe   Subscribe to subject
  template    CRUD operations with templates

Flags:
      --config string   config file (default is $HOME/.nats-cli.yaml)
  -h, --help            help for nats-cli
  -t, --toggle          Help message for toggle

Use "nats-cli [command] --help" for more information about a command.
```

# Subscribe exmaple
```bash
nats-cli sub -a your-nats-server -c nats-cluster-id -s subject -d 24h
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
