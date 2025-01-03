# goRelay

`goRelay` is a TCP-based **intranet penetration** tool written in `go`.

### Background:
This tool is used when there is a need to map a port on an internal network machine to the public internet for access. It provides end-to-end communication functionality.

### Requirements:
At least one machine with a public IP address is required to use this tool.

### Latest Feature:
Supports port reuse.

## Deployment Diagram

![Flowchart](img/d4.png)

## How to Obtain the Executable

### From Released Versions

You can obtain the following executable files from the released versions:

`goRelay` is an executable that distinguishes server types using the `--type` flag. The specific types are:

- `relayServer`: A relay server used for receiving and sending data. It needs to be deployed in the intranet environment where client requests are made. Multiple instances can be deployed to support port reuse.
- `pipeServer`: A pipe server used for data transmission. It needs to be deployed on a server with a public IP address.
- `pipeClient`: A pipe client used for user data transmission. It needs to be deployed in the real service's intranet environment.
- `relayClient`: A relay client used for receiving and sending data. It needs to be deployed in the real service's intranet environment. You need to configure the `relayServer` information to support port reuse.

### Compilation

Clone the repository, then compile by running `bash build.sh` with the version number, e.g.:

```bash
➜  goRelay git:(main) ✗ bash build.sh v0.0.1
build project
go build -ldflags "-X goRelay/pkg.Version=v0.0.1 -X goRelay/pkg.BuildAt=2025-01-04 -X goRelay/pkg.GitCommit=9002ae53fff26c433cdcde76abed1781c3cc218a" -o ./bin/goRelay 
tar zcvf pipeSourcev0.0.1.tar.gz ./bin
./bin/
./bin/goRelay
➜  goRelay git:(main) ✗ 
```

After compiling, the corresponding binary files will be generated in the `./bin/` directory.

```bash
➜  goRelay git:(main) ✗ ls bin/
goRelay
➜  goRelay git:(main) ✗ 
```

## How to Run

### `pipeServer`

`pipeServer` is the pipe server used for data transmission. It needs to be deployed on a server with a public IP address.

To start this service, first create a configuration file, such as `conf_pipeServer.json`:

```json
{
    "listen_pipe_server_addr":":8888",
    "black_ip_list":[
        "127.0.0.3"
    ],
    "white_ip_list":[
        "127.0.0.1"
    ],
    "debug_log":true
}
```

Where:
- `listen_pipe_server_addr` is the address to listen on externally. If you want to listen on a specific network interface, specify the interface address, e.g., `192.168.2.3:8888`.
- `white_ip_list` sets the whitelist, allowing only connections from listed IPs. If empty, no whitelist is applied, and all can connect.
- `black_ip_list` sets the blacklist, preventing connections from IPs in the list. If empty, no blacklist is applied. The blacklist takes precedence over the whitelist if they contain the same values.
- `debug_log` determines whether to output debug logs.

To start the `pipeServer`, simply specify the configuration file:

```bash
➜  bin git:(main) ✗ ./goRelay --type pipeServer --config conf/conf_pipeServer.json 
```

### `relayServer`

`relayServer` is the relay server used for receiving and sending user data. It needs to be deployed in the intranet environment requested by the client.

Multiple `relayServer` instances can be started to support port reuse. Each `relayServer` needs a unique `id`.

Create a configuration file, for example `conf_relayServer.json`:

```json
{
    "id": "client1",
    "pipe_server_addr":"127.0.0.1:8888",
    "listen_relay_server_addr":":10010",
    "white_ip_list":[
        "127.0.0.1"
    ],
    "debug_log":true
}
```

Where:
- `id` is the unique identifier for the `relayServer`. Ensure it is unique if running multiple instances, and it should be a string with a recommended complexity.
- `pipe_server_addr` is the address of the `pipeServer`.
- `listen_relay_server_addr` is the address the `relayServer` listens on.
- `white_ip_list` sets the whitelist.

You can create a second configuration file, `conf_relayServer2.json`:

```json
{
    "id": "client2",
    "pipe_server_addr":"127.0.0.1:8888",
    "listen_relay_server_addr":":10012",
    "white_ip_list":[
        "127.0.0.1"
    ],
    "debug_log":true
}
```

Ensure that `id` is unique, and if deploying on the same machine, `listen_relay_server_addr` must be unique.

To start the service, specify the configuration file:

```bash
➜  bin git:(main) ✗ ./goRelay --type relayServer --config conf/conf_relayServer.json 
```

### `relayClient`

`relayClient` is the relay client used for receiving/sending data from the pipe client and real services. It needs to be deployed in the real service's intranet environment.

To start this service, configure the `relayServer` with the corresponding `id` and specify the address of the real service and the port to listen on for connections from the pipe client.

Create a configuration file, such as `conf_relayClient.json`:

```json
{
    "listen_relay_client_addr": ":10011",
    "white_ip_list":[
        "127.0.0.1"
    ],
    "debug_log": true,
    "realServerInfo": [
        {
            "id": "client1",
            "real_Server_Addr": "127.0.0.1:80"
        },
        {
            "id": "client2",
            "real_Server_Addr": "127.0.0.1:22"
        }
    ]
}
```

Where:
- `listen_relay_client_addr` is the external address to listen on for connections from the `pipeClient`.
- `realServerInfo` is an array of real server information, mapping each `id` to a real service address. The `id` must match the `id` in the corresponding `relayServer`.

Start the service by specifying the configuration file:

```bash
➜  bin git:(main) ✗ ./goRelay --type relayClient --config conf/conf_relayClient.json 
```

### `pipeClient`

`pipeClient` is the pipe client used for transmitting data from the pipe server and relay client. It should be deployed in the real service's intranet environment.

To start this service, create a configuration file, such as `conf_pipeClient.json`:

```json
{
    "pipe_server_addr":"127.0.0.1:8888",
    "relay_client_addr":"127.0.0.1:10011",
    "debug_log": true
}
```

Where:
- `pipe_server_addr` is the address of the `pipeServer`.
- `relay_client_addr` is the address of the `relayClient`.

Start the service by specifying the configuration file:

```bash
➜  bin git:(main) ✗ ./goRelay --type pipeClient --config conf/conf_pipeClient.json 
```

## Data Encryption

If you are not the owner of the public server, you must consider the risk of man-in-the-middle attacks. To ensure data security, it is recommended to use data encryption, as shown in the green section of the deployment diagram. Data transmission can be encrypted. However, please note that the current version does not implement data encryption, but an interface for this purpose is reserved. You can modify the `Encode` and `Decode` functions in the `pipeProtocol/enDecode.go` file to implement data encryption.

The functions are as follows:

```go
func Encode(s []byte) []byte {
	return s
}

func Decode(s []byte) []byte {
	return s
}
```

Here, `Encode` is used for encryption, and `Decode` is used for decryption.

## Other

This project does not accept any feature requests.