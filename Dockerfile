FROM golang:1.21 as build
WORKDIR /app
# Copy dependencies list
COPY go.mod go.sum ./
# Build with optional lambda.norpc tag
COPY *.go ./

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -tags lambda.norpc -o toggl-cron .
# Copy artifacts to a clean image
FROM public.ecr.aws/lambda/provided:al2
COPY --from=build /app/toggl-cron ./toggl-cron
ADD aws-lambda-rie /usr/local/bin/aws-lambda-rie
ADD entry.sh ./entry.sh

ENTRYPOINT [ "/bin/sh", "./entry.sh" ]