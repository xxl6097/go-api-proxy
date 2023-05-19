# openai-proxy


## Deploy

`docker-compose.yml` example:

```
   openai-proxy:
    image: ghcr.io/orvice/openai-proxy:main
    restart: always
    container_name:   openai-proxy
    ports:
      - 8080
    environment:
      - OPENAI_API_KEY=sk-xxxx
```


```docker build -t openai-proxy .```
```docker run --name openai-proxy -p 3456:3456 -d openai-proxy```

docker login -u prdsl-1683373983040 -p ffd28ef40d69e45f4e919e6b109d5a98601e3acd clife-devops-docker.pkg.coding.net