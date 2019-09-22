# build this by using `docker build -t sorogbesi-backend .`
# run the image with `docker run -it -p 8080:8080 sorogbesi-backend`

FROM golang:1.12.7-buster

# create application directory in the container
RUN mkdir -p /app

# copy the code into application directory in the container
ADD . /app

# designate the working directory within the container
WORKDIR /app

# copies the go.mod and go.sum to WORKDIR
COPY go.mod go.sum ./

# Download and install dependencies e.g. github.com/gorilla/mux etc.
RUN go mod download

# note that presence of multiple packages means cannot use `-o main` like `go build -o main ./...`
RUN CGO_ENABLED=0 go build ./...

# Run the compiled binary.
ENTRYPOINT ["go", "run",  "main.go"]