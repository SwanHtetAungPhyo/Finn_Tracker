FROM golang:1.23.2
WORKDIR /app
COPY . .
RUN go build -o expense-service main.go
EXPOSE 5002
CMD [ "./expense-service" ]