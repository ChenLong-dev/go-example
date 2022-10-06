# Prometheus监控NATS

## docker执行方式
### grafana的docker执行
```azure
docker run -d -p 3000:3000 --name grafana grafana/grafana:latest
```

### prometheus的docker执行
```azure
docker run -p 9090:9090 -v /d/ChenLong/00_study/02_code/go-example/nats/monitor/prometheus.yml:/etc/prometheus/prometheus.yml bitnami/prometheus:latest
```
http://127.0.0.1:7777/metrics

### prometheus-nats-exporter的docker执行
```azure
docker run -p 7777:7777 natsio/prometheus-nats-exporter:latest -varz "http://192.168.0.110:8222"
```
