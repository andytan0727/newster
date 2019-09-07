# Newster

Simple web app that catches fresh news everyday! Its backend API is inspired by [tophubs/TopList](https://github.com/tophubs/TopList).

## Development
Before moving forward, please ensure that you have the following executable locally:

```bash
go >= v1.12
node >= 10.12.0
yarn >= 1.15.2
```

To get started in development, first execute the following commands in order:

Firstly, install all frontend dependencies:
```bash
cd frontend
yarn install
```

Then, tidy up `go module`:
```bash
cd api
go mod tidy
```
The `go mod tidy` will install the required `go modules` accordingly.

Next, we can start our development server locally:
```bash
yarn start:client # for starting up client
yarn start:server # for starting up api server
```

To test production ready `Go` server hosting both front and backend, executes the following commands accordingly:
```bash
yarn build:client  # build client and output dist dir
yarn build:server  # build api server and output an executable
yarn serve         # execute the Go executable built
```
