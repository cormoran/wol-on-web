FROM golang:1.22.4-alpine as builder

ENV CGO_ENABLED=0
ENV GOOS=linux
RUN apk add git --no-cache

COPY "./webapi" /appbuild
WORKDIR /appbuild
RUN go build -o output/bin/app

FROM scratch
COPY --from=builder /appbuild/output/bin/app /app
COPY "./webapi/static" /static
CMD ["/app"]
