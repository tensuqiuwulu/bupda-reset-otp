# builder image
FROM golang:1.18.6-alpine as builder
WORKDIR /build
COPY . .
RUN apk add git && CGO_ENABLED=0 GOOS=linux go build -o bupda-reset-otp .
# RUN go build -o be-service-teman-bunda .

# generate clean, final image for end users
FROM alpine
RUN apk update && apk add ca-certificates && apk add tzdata && apk add git
COPY --from=builder /build .
ENV TZ="Asia/Makassar"
EXPOSE 9300

CMD ./bupda-reset-otp