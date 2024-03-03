FROM golang:1.22-alpine as builder
ARG CGO_ENABLED=0
WORKDIR /app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify


COPY . .
# RUN CGO_ENABLED=0 GOOS=linux go build -o /server
RUN go build -v -o server

FROM scratch
COPY --from=builder /app/server /server

EXPOSE 5000

# This should be the path to binary we run to start the app.
CMD ["/server"]
