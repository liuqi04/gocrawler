# use go v.11 as base image
FROM golang:1.12

# set the working directory
WORKDIR /go/src/app

# copy current directory contents to working directory
COPY . .

# download all the dependencies
RUN go get -d -v ./...

# install the package
RUN go install -v ./...

# port 80 to outside world
EXPOSE 8080

# run crawler.go when the container launches
CMD ["app"]