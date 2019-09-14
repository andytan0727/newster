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

Then, download all the modules specified by `go module`:
```bash
cd api
go mod download
```

Next, we can start our development server locally:
```bash
yarn start:client # for starting up client
yarn start:server # for starting up api server
```

Or with one-shot command:
```bash
yarn start        # concurrently start dev client and server
```

To test production ready `Go` server hosting both front and backend, executes the following commands accordingly:
```bash
yarn build
./main             # or ./main.exe in Windows
```

## Docker
To containerize `newster`, firstly we build the `newster` image with `docker`. Make sure your local system has `docker` installed:
```bash
docker build -t newster .
```
The command above will produce a `newster` image which size is only as small as 10+ MB as of writing time.

Next, we can run the image with:
```bash
docker run -it --rm -p 8000:8000 --name newster-app newster
```
The command above will create automatically disposed container using `newster` image built previously.
