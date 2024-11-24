FROM golang:1.23.3-bullseye
RUN apt-get update && apt-get install -y curl && rm -rf /var/lib/apt/lists/*
WORKDIR  /app 
COPY . .
RUN go mod tidy
RUN  go build -o stock-service .
EXPOSE 7001
ENTRYPOINT [ "./stock-service" ]
