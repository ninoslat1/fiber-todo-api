# Use the official Go image as the build stage.
FROM golang:latest AS firststage

WORKDIR /build/

COPY . /build

ENV CGO_ENABLED=0

# Install swag jika Anda menggunakannya untuk generate swagger docs
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Generate swagger docs (uncomment jika Anda menggunakan swag)
# RUN swag init

RUN go get
# Build the Go application with optimizations.
RUN go build -o todo-api

# Create a lightweight final stage to run the application.
FROM alpine:latest

WORKDIR /app/

RUN apk update && apk upgrade && apk add --no-cache tzdata gcompat

ENV TZ=Asia/Jakarta

COPY --from=firststage /build/todo-api .
COPY --from=firststage /build/docs /app/docs

CMD ["./todo-api"]