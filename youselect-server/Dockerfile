FROM golang:1.15.0-alpine3.12

WORKDIR /app

ENV GOBIN="$GOPATH/bin"

# Fetch dependencies first; they are less susceptible to change on every build
# and will therefore be cached for speeding up the next build
COPY ./go.mod ./
RUN go mod download -x

# Import the code from the context.
COPY . .

EXPOSE 5500
CMD [ "go", "run", "main.go" ]