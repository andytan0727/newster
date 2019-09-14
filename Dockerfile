# building frontend
FROM node:12.10-alpine as frontend

WORKDIR /frontend

# install tools needed to build node-sass
RUN apk update && apk add python g++ make && rm -rf /var/cache/apk/*

COPY ./frontend/package.json ./frontend/yarn.lock /tmp/
RUN cd /tmp && \
    yarn install && \
    cd /frontend && \
    ln -s /tmp/node_modules

COPY ./frontend .
RUN yarn build


# building backend api
FROM golang:1.13.0-alpine as api

WORKDIR /backend

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY ./api ./api
COPY main.go .
RUN go build main.go

# start a new stage from scratch
# minimize the output image size
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /usr/src/newster

# Copy the frontend dist folder and pre-built api binary file from the previous stages
COPY --from=frontend /frontend/dist ./frontend/dist
COPY --from=api /backend/main .

# execute main api/server binary
# frontend spa is served at /, whereas api endpoint is served at /api
CMD ["./main"] 
