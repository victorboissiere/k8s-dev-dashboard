FROM golang:1.12.6-alpine3.9

RUN apk --update add git
RUN go get k8s.io/api/core/v1 && \
    go get k8s.io/client-go/rest && \
    go get k8s.io/client-go/tools/clientcmd && \
    go get k8s.io/apimachinery/pkg/apis/meta/v1 && \
    go get github.com/dgrijalva/jwt-go && \
    go get github.com/labstack/echo && \
    go get github.com/googleapis/gnostic/OpenAPIv2 && \
    go get github.com/golang/protobuf/proto && \
    go get github.com/jasonlvhit/gocron

WORKDIR /gopath/src/app
COPY . /gopath/src/app/

RUN go build -ldflags "-s -w" -o /opt/app/devDashboard

FROM alpine:3.7

COPY --from=0 /opt/app/devDashboard /opt/app/

WORKDIR /opt/app/

ENV PORT 3000
CMD ["./devDashboard"]
