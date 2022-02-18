# Specify the base image for the go app.
FROM golang:1.17
# Specify that we now need to execute any commnds in this directory.
WORKDIR /go/src/github.com/postgres-go
# Copy everything from this project into the filesystem of the container.
COPY . .
# Obtain the package needed to run code. Alternatively use GO Modules.
RUN go get -u github.com/lib/pq
RUN go get -u github.com/gorilla/mux
# Compile our app
RUN go build -o main .
# Start the application
CMD ["./main"]
