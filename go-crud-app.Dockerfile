FROM golang:latest

#setup working directory
WORKDIR /go/src/github.com/lakshanwd/go-crud

#copy files
COPY dao ./dao
COPY db ./db
COPY handler ./handler
COPY repository ./repository
COPY server.go ./server.go

#install dependancies and build executable
RUN go get -d -v ./... && go install -v ./...

#copy config files
COPY config.prod.json ./config.json

#run executable
CMD ["go-crud"]

#expose port
EXPOSE 8080