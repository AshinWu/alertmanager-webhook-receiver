# alertmanger-webhook-receiver

Is a simple demo of a webhook receiver for the Prometheus Alertmanager. It provides a route `/alerts` to fetch the POST request from the AlertManager and get the request JSON body.
for more information, please see: [webhook_config](https://prometheus.io/docs/alerting/latest/configuration/#webhook_config)

This demo has no external dependencies, you could handle the request body by yourself.

## Usage

Alertmanger config example:

```yaml
global:
      resolve_timeout: 5m
    route:
      group_by: ['alertname']
      group_wait: 10s
      group_interval: 1m
      repeat_interval: 2m
      receiver: 'web.hook'
    receivers: 
    - name: 'web.hook'
      webhook_configs:
      - url: 'http://ip:8090/alerts'
```

Run in local:

```shell script
go run main.go
```

check connection: `http://localhost:8090/health`

## Building

Clone this repo into your GOPATH and run `go build .` by yourself.

Or could use ./build/alertmanager-webhook-receiver (build for linux-amd64)

Ensure executable permissions:

```shell script
chmod +x build/alertmanager-webhook-receiver
```

then in the root directory of project:

```yaml
docker build -t alertmanager-webhook-receiver:v1 .
```

Running:

```shell script
docker run -p 8090:8090 -d alertmanager-webhook-receiver:v1 .
```
