FROM golang:1.12
WORKDIR /go/src/app
#ENV GO111MODULE=on
#COPY . .

RUN go get github.com/pilu/fresh
#RUN go get github.com/jinzhu/gorm
#RUN go get github.com/jinzhu/gorm/dialects/mysql
#RUN go get github.com/gorilla/mux


CMD ["go", "run", "/go/src/app/main.go"]
CMD ["fresh"]