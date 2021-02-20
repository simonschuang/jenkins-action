FROM golang:1.15.8-alpine

RUN apk add git && \
  go get -u -v github.com/simonschuang/jenkins-action && \
  go install -i -n github.com/simonschuang/jenkins-action && \
  go install -i github.com/simonschuang/jenkins-action

