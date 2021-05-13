# gomesh

## Alpha state

This software is yet in alpha state, please use it with caution.

## Introduction

gomesh generates peer configuration files for WireGuard mesh networks.

## Installation 

Just copy a released binary to a directory and run it.


## Detailed Usages

You may refer to the program's help page for usages. Use the `-h` switch or the `--help` switch to print the help page.

```shell
$ gomesh -h
This little tool will generate and manage configuration files for Wireguard Mesh VPNs.

Usage:
  gomesh [command]

Available Commands:
  add         Add or update a peer to/in the registry
  del         Delete a peer from registry
  generate    Generate configs
  help        Help about any command
  show        Print a table with the peers

Flags:
  -d, --database string   database file (default is database.json) (default "database.json")
  -h, --help              help for gomesh
  -v, --version           version for gomesh

Use "gomesh [command] --help" for more information about a command.
```

You can get help with the commands

```shell
$./gomesh add -h
Wil add a peer to the registry or update

Usage:
  gomesh add [flags]

Flags:
  -a, --address strings        Address of the server. (Required)
      --allowedips strings     additional allowed IP addresses
      --dns string             DNS server
  -e, --endpoint string        The peer's endpoint
  -f, --fwmark int             Mark the outgoing packets with
  -h, --help                   help for add
  -l, --listenport int         Port to listen on, default 51820 (default 51820)
  -m, --mtu int                Server interface MTU
  -n, --name string            endpoint. (Required)
      --postDown string        Command to run after bringing the interface DOWN
      --postUP string          Command to run after bringing the interface UP
      --preDown string         Command to run before bringing the interface DOWN
      --preUP string           Command to run before bringing the interface UP
  -p, --privatekey string      private key of server interface (if none given one will be generated
  -r, --routing_table string   Server routing table
  -s, --saveconfig             Save config between reboots
  -u, --update                 Update Peer if existing.

Global Flags:
  -d, --database string   database file (default "database.json")
```

## License

Licensed under the MIT license

<https://opensource.org/licenses/MIT>

<img src="https://opensource.org/files/OSIApproved_1.png" width="120">

(C) 2021 Nagy Károly Gábriel

## Credits

This project is inspired by [wg-meshconf](https://github.com/k4yt3x/wg-meshconf) by k4yt3x
