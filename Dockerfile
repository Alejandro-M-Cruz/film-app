FROM golang:1.24

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -v -o film-app

RUN apt-get update && apt-get install -y sqlite3
RUN make db-migrate && make db-seed

EXPOSE 8000

CMD ["/app/film-app"]
