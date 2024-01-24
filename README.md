# demo api in go

Demo

## Usage

```sh
# GET this app version
$ curl -X GET localhost:3004/version | jq

# PUT this app NEW version
$ curl -X PUT localhost:3004/version --form "version=v4.4.4" --form "sha=dhjkashdjkashdjk"| jq

## -----------

# GET github demo app version
$ curl -X GET localhost:3004/github | jq

# PUT github app NEW version
$ curl -X PUT localhost:3004/github --form "version=v4.4.4" --form "sha=dhjkashdjkashdjk"| jq

## -----------

# GET gitlab demo app version
$ curl -X GET localhost:3004/gitlab | jq

# PUT gitlab app NEW version
$ curl -X PUT localhost:3004/gitlab --form "version=v4.4.4" --form "sha=dhjkashdjkashdjk"| jq
```
