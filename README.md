# Ishmael

This is a tool for poking docker in scripts and gaining specific information such as if a container is passing its health checks at present.

## Installing and Updating

You must have the Go compiler installed and correctly configured, then you can run:

```bash
go get -u github.com/Nekroze/ishmael/...
```

## Usage

All subcommands support the `--help` and `-h` switches to get more information about what can be done but here are some examples:

### Check if a container is alive:

This command will do a one off check to see if a container with ID `f31760063290` is running. Returns a non zero exit code if the container is not currently running.

```bash
ishmael alive f31760063290
```
### Wait for a container to be alive:

This command will do one check every second for 10 seconds to see if a container with ID `f31760063290` is running. Returns a non zero exit code if the container is not up and running in time.

```bash
ishmael alive --wait 10 f31760063290
```

Should the container be restarting or not running at all then it will stop immediately.

### Check if a container is healthy:

This command will do a one off check to see if a container with ID `f31760063290` is healthy. Returns a non zero exit code if the container is not currently healthy.

```bash
ishmael healthy f31760063290
```

### Wait for a container to be healthy:

This command will do one check every second for 10 seconds to see if a container with ID `f31760063290` is running. Returns a non zero exit code if the container is not healthy in time.

```bash
ishmael healthy --wait 10 f31760063290
```

Should the container be running but not have any health status or not even be running at all then it will stop immediately.

### Find the id for a docker compose project's service

Say we have a docker compose project called foo with a service called bar. If we wanted to get the ID of a running instance of the service, the `find` subcommand is your friend.

```bash
ishmael find foo bar
```

This will return 0 and print out the ID of the first running instance of the service it can find.

### Get the addresses a container exposes

If we had a running container called `my_nginx_1` and it had declared port `8080` as exposed, we could get the address to this container port using the `address` subcommand.

```bash
ishmael address my_nginx_1
```

Then we might get `172.16.0.2:8080` depending on the IP address docker assigns it. For containers that use host networking `127.0.0.1` will be given instead in an attempt to access the port on the loopback interface.
