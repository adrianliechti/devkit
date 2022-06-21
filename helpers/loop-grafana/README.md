```bash
docker build . --tag grafana

docker run -it --rm \
    -p 3000:3000 \
    -p 3100:3100 \
    -p 3200:3200 \
    -p 9090:9090 \
    grafana
```