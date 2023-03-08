FROM golang:1.19-bullseye

COPY ./build/dev/app/entrypoint /entrypoint
RUN sed -i 's/\r$//g' /entrypoint
RUN chmod +x /entrypoint

COPY ./build/dev/app/start /start
RUN sed -i 's/\r$//g' /start
RUN chmod +x /start

WORKDIR /app

COPY go.mod go.sum /app/
RUN go mod download
RUN go install github.com/cosmtrek/air@latest

ENTRYPOINT ["/entrypoint"]