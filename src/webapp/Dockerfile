# From golang docker to be able to build our golang application
# also choosen an specific version to be sure that this code will never broke
FROM golang:1.12.9

# Set the Current Working Directory inside the container
WORKDIR /hello

# Copy the source from the current directory to the Working Directory inside the container
COPY ./src/webapp/webapp.go .

# Build the Go app
RUN go build -o main webapp.go

# Expose desired port to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
