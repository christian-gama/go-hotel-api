FROM ubuntu:20.04 AS base
RUN apt-get update
RUN apt-get install -y \
    curl \
    lsb-release \
    build-essential

ARG GOVERSION=1.19
ARG GOARCH=amd64

# Install Go
RUN curl -O -L "https://golang.org/dl/go${GOVERSION}.linux-${GOARCH}.tar.gz" && \
    curl -sL https://golang.org/dl/ | grep -A 5 -w "go${GOVERSION}.linux-${GOARCH}.tar.gz" && \
    tar -xf "go${GOVERSION}.linux-${GOARCH}.tar.gz" && \
    mv go /usr/local/go && \
    rm "go${GOVERSION}.linux-${GOARCH}.tar.gz" && \
    mkdir ~/go

# Set paths
ENV PATH $PATH:/usr/local/go/bin
ENV GOPATH $HOME/go

RUN curl -sSL "https://github.com/gotestyourself/gotestsum/releases/download/v1.8.1/gotestsum_1.8.1_linux_amd64.tar.gz" \
    | tar -xz -C /usr/local/go/bin gotestsum 


FROM base AS gobooking
ARG WORKDIR
WORKDIR ${WORKDIR} 
COPY ./go.mod ./go.mod ./
COPY . ./
