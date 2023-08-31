FROM golang:alpine@sha256:445f34008a77b0b98bf1821bf7ef5e37bb63cc42d22ee7c21cc17041070d134f

RUN mkdir /src
WORKDIR /src

COPY ./go.mod /src/go.mod
COPY ./go.sum /src/go.sum
RUN go mod download

COPY ./ /src/
RUN go build -o github.com/amren1254/gin-docker ./

FROM alpine:3.18@sha256:7144f7bab3d4c2648d7e59409f15ec52a18006a128c733fcff20d3a4a54ba44a
RUN apk --no-cache add \
    ca-certificates \
    git \
    bash

# Allow git to run on mounted directories
RUN git config --global --add safe.directory '*'

WORKDIR /root/
COPY --from=0 /src/github.com/amren1254/gin-docker ./
COPY --from=0 /src/github.com/amren1254/gin-docker ./

ENV PATH="${PATH}:/root"

ENTRYPOINT ["bash", "/root/github.com/amren1254/gin-docker"]