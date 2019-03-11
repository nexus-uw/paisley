FROM golang:1.9

WORKDIR /go/src/paisley
RUN set -x \
    && go get -v github.com/revel/revel \
    && go get -v github.com/revel/cmd/revel

COPY . .
RUN go get -d -v ./...

RUN revel build -a paisley
EXPOSE 9000

# environment vars that you must provide
ENV REDDIT_ID=replaceMe
ENV REDDIT_SECRET=replaceMe
ENV REDDIT_USER=replaceMe
ENV REDDIT_PASSWORD=replaceMe

ENTRYPOINT revel run -a paisley --run-mode=prod -p 9000
