FROM qmcgaw/godevcontainer

# Need curl to install atlas
RUN apk --no-cache add curl

# Install sqlc to use command line within container for testing
RUN go mod init tmp && go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

# Install atlas to use command line within container for testing
RUN curl -sSf https://atlasgo.sh | sh
