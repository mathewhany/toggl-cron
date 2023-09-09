FROM golang:1.21 as build
WORKDIR /app
# Copy dependencies list
COPY go.mod go.sum ./
# Build with optional lambda.norpc tag
COPY main.go .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -tags lambda.norpc -o main main.go
# Copy artifacts to a clean image
FROM public.ecr.aws/lambda/provided:al2
COPY --from=build /app/main ./main
ADD aws-lambda-rie /usr/local/bin/aws-lambda-rie
ADD entry.sh ./entry.sh

ENTRYPOINT [ "/bin/sh", "./entry.sh" ]