FROM golang:1.22.5 as build

WORKDIR /app

COPY go.mod go.sum ./

RUN go version
RUN go mod tidy

COPY . .

RUN CGO_ENABLED=0 go build


FROM scratch

WORKDIR /app

COPY --from=build /app .

EXPOSE 8080
CMD ["/app/book-library"]

