FROM golang:1.12.7
RUN apt-get update && apt-get install -y --no-install-recommends libgtk-3-dev mingw-w64
RUN GO111MODULE=on go get \
	github.com/golangci/golangci-lint/cmd/golangci-lint@v1.24.0 \
	github.com/IndioInc/go-autoupdate \
	golang.org/x/tools/cmd/goimports
WORKDIR /app
ADD go.mod go.sum /app/
RUN go mod download
ADD . .