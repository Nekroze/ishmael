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
