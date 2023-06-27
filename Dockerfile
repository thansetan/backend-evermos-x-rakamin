FROM golang:1.20.2


WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
COPY .env /.env
RUN go build -o ./app/main-app ./app/

EXPOSE ${HTTP_PORT}

CMD "./app/main-app"