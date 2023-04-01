
FROM alpine:latest

WORKDIR /root
# ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.9.0/wait /wait

# RUN chmod +x /wait

COPY bin/main .
EXPOSE 8086
CMD ["./authApp"]