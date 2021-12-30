FROM golang:1.17-alpine as builder

# build stage
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o /serverapi

# target stage
FROM alpine
WORKDIR /

# Lets us set a timezone for logging purposes
RUN apk update && apk add --no-cache tzdata
ENV TZ Asia/Kuala_Lumpur

# Copy necessary files for execution
COPY --from=builder /serverapi /serverapi
COPY ./public /public
COPY ./startup.sh /startup.sh
RUN chmod +x /startup.sh

# Sets up a cronjob to run every 6 hours to ingest data
RUN echo '*  */6  *  *  * /serverapi ingest >> /ingest.log 2>&1' > /etc/crontabs/root

EXPOSE 8000

CMD [ "/startup.sh" ]