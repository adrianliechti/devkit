```bash
docker build . --tag elk

docker run -it --rm \
    -p 9200:9200 \
    -p 5000:5000 \
    -p 5601:5601 \
    --ulimit nofile=65535:65535 \
    --ulimit memlock=-1:-1 \
    elk
```