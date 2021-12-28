FROM golang:1.17-alpine as builder

# first (build) stage

WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o /serverapi

# final (target) stage

FROM alpine
WORKDIR /
COPY --from=builder /serverapi /serverapi
EXPOSE 8000
CMD [ "/serverapi", "serve" ]