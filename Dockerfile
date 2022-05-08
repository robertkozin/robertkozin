FROM alpine

RUN apk add golang

RUN go install

CMD ['robertkozin-website']
