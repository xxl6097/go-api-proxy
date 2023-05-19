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
