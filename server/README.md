# Tanxz Server
This is the server code for the Tanxz multiplayer game

## Refresh Go Imports

`go mod download`

## Deploy
Locally, run `./deploy.sh`

When that command has completed, `ssh ec2-34-217-206-186.us-west-2.compute.amazonaws.com`

Then run the command `nohup ./main &`

## Build for Ubuntu
```
env GOOS=linux GOARCH=amd64 go build main.go
```

## Install Go For Mac using Brew
```
brew install go
```
or
```
brew upgrade go
```