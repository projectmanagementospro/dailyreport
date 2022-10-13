FROM golang:alpine
RUN mkdir /app

WORKDIR /golang-dev

ADD go.mod .
ADD go.sum .

RUN go mod download
ADD . .

# RUN go install -mod=mod github.com/jinzhu/gorm
# RUN go install -mod=mod github.com/dgrijalva/jwt-go
# RUN go install -mod=mod github.com/joho/godotenv
# RUN go get golang.org/x/crypto
# RUN go install -mod=mod github.com/githubnemo/CompileDaemon

EXPOSE 8000
EXPOSE 8080

# ENTRYPOINT ./start.sh
