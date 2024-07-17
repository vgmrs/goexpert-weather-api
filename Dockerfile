FROM golang:1.22-alpine AS binary

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -C cmd/ -ldflags="-w -s" -o server .

FROM scratch

COPY --from=binary /app/cmd /

EXPOSE 8000

CMD [ "./server" ]
