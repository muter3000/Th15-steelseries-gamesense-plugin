# Th15-steelseries-gamesense-plugin
Repo containing code for injectable Touhou 15 steelseries gamesense plugin

## Building

The plugin should be build with the following command
```sh
GOOS=windows GOARCH=386 CGO_ENABLED=1 go build -o helloworld.dll -buildmode=c-shared
```

## Running the plugin

The plugin can be run by injecting it into th15 process, e.g. with [System informer](https://github.com/winsiderss/systeminformer)
