# demo api in go

Demo

## Usage

```sh
# GET this app version
$ curl -X GET localhost:3004/version | jq

# PUT this app NEW version
$ curl -X PUT localhost:3004/version --form "version=v4.4.4" --form "sha=dhjkashdjkashdjk"| jq
$ curl -X PUT https://sie-demo-go-api.xgeeks.tech/version --form "version=v4.4.4" --form "sha=dhjkashdjkashdjk"| jq

## -----------

# GET github demo app version
$ curl -X GET localhost:3004/github | jq
$ curl -X GET https://sie-demo-go-api.xgeeks.tech/github | jq

# PUT github app NEW version
$ curl -X PUT localhost:3004/github --form "version=v4.4.4" --form "sha=dhjkashdjkashdjk"| jq
$ curl -X PUT https://sie-demo-go-api.xgeeks.tech/github --form "version=v2.0.0" --form "sha=dhjkashdjkashdjk"| jq
$ curl -X PUT https://sie-demo-go-api.xgeeks.tech/github --form "version=v2.0.0" --form "sha=35e53aaf69af59d42c06e84bb835484c12ee8225"| jq
$ curl -X PUT https://sie-demo-go-api.xgeeks.tech/github --form "version=v2.0.0" --form "sha=f1c2a0d4b27ab71fd754fbfa406d1a73e6756c7e"| jq

## -----------

# GET gitlab demo app version
$ curl -X GET localhost:3004/gitlab | jq
$ curl -X GET https://sie-demo-go-api.xgeeks.tech/gitlab | jq

# PUT gitlab app NEW version
$ curl -X PUT localhost:3004/gitlab --form "version=v4.4.4" --form "sha=dhjkashdjkashdjk"| jq
$ curl -X PUT https://sie-demo-go-api.xgeeks.tech/gitlab --form "version=v4.4.4" --form "sha=dhjkashdjkashdjk"| jq

$ curl -X PUT "https://sie-demo-go-api.xgeeks.tech/gitlab" --form "version=v1.0.1" --form "sha=6e2e3ab452d4a87da2e5dacc1d2a5fb3c91e5c00"| jq
$ curl -X PUT "https://sie-demo-go-api.xgeeks.tech/gitlab" --form "version=v1.0.0" --form "sha=a628f07bad4fb0db5a2e92fd4718a7467232b867"| jq
```
