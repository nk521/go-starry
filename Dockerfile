# # Build Stage
# FROM alpine:latest:1.13 AS build-stage

# LABEL app="build-go-starry"
# LABEL REPO="https://github.com/nk521/go-starry"

# ENV PROJPATH=/go/src/github.com/nk521/go-starry

# # Because of https://github.com/docker/docker/issues/14914
# ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin

# ADD . /go/src/github.com/nk521/go-starry
# WORKDIR /go/src/github.com/nk521/go-starry

# RUN make build-alpine

# # Final Stage
# FROM alpine:latest

# ARG GIT_COMMIT
# ARG VERSION
# LABEL REPO="https://github.com/nk521/go-starry"
# LABEL GIT_COMMIT=$GIT_COMMIT
# LABEL VERSION=$VERSION

# # Because of https://github.com/docker/docker/issues/14914
# ENV PATH=$PATH:/opt/go-starry/bin

# WORKDIR /opt/go-starry/bin

# COPY --from=build-stage /go/src/github.com/nk521/go-starry/bin/go-starry /opt/go-starry/bin/
# RUN chmod +x /opt/go-starry/bin/go-starry

# # Create appuser
# RUN adduser -D -g '' go-starry
# USER go-starry

# ENTRYPOINT ["/usr/bin/dumb-init", "--"]

# CMD ["/opt/go-starry/bin/go-starry"]
