FROM golang:1.10

WORKDIR /go/src/github.com/nexus-uw/paisley
RUN set -x \
    && go get -v github.com/revel/revel \
    && go get -v github.com/revel/cmd/revel

COPY . .
RUN go get -d -v ./...

RUN revel build -a github.com/nexus-uw/paisley
EXPOSE 9000

# environment vars that you must provide
ENV REDDIT_ID=replaceMe
ENV REDDIT_SECRET=replaceMe
ENV REDDIT_USER=replaceMe
ENV REDDIT_PASSWORD=replaceMe

ENTRYPOINT revel run -a github.com/nexus-uw/paisley --run-mode=prod -p 9000
