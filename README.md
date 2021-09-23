# Fake Rest Server

A fake rest server written in GO to serve data from json file. Dynamically creates the routes based on the objects in the JSON. Get the contents of a resource, get a particular content using it's id if the resource is a json array, append or replace a resource and delete a resource. If a static directory is specified the server can serve the static files from the directory.n.

## Setup

Make sure you have `go` setup along with `build-essential`.

### Build & Start using Make
```bash
make build # builds for the current platform
make start_server #starts the server from dist dir
make kill_server #kills the server
```


## Usage

```bash
./dist/fake-rest-server -port $(PORT) -file $(DB_FILE) -serve-static $(STATIC_FOLDER_LOCATION)
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## TODO

 - [ ] Tests
