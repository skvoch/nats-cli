# Command line interface fot NATS Streaming
# Usage
```bash
nats-cli [command]

Available Commands:
  help        Help about any command
  publish     Publish message to subject
  subscribe   Subscribe to subject

Flags:
      --config string   config file (default is $HOME/.nats-cli.yaml)
  -h, --help            help for nats-cli
  -t, --toggle          Help message for toggle

Use "nats-cli [command] --help" for more information about a command.
```

# Subscribe exmaple
```bash
nats-cli sub -a your-nats -c nats-cluster-id -s subject -d 24h
```
