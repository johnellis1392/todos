# ############# #
# ### Build ### #
# ############# #
FROM go:alpine AS build

ENV GOPATH=/tmp/go
WORKDIR /tmp/go/src/app

RUN apk add --update git
RUN go get -v -u \
    github.com/gorilla/mux \
    github.com/aws/aws-sdk-go \
    github.com/aws/aws-lambda-go

COPY . .
RUN go build -o main


# ########### #
# ### Run ### #
# ########### #
FROM alpine:latest AS run

EXPOSE 3000/tcp
WORKDIR /usr/src/app
COPY --from=build /tmp/go/src/app/main ./main

ENTRYPOINT [ "/usr/src/app/main" ]
