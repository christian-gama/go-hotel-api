FROM ubuntu:20.04 AS base
ARG GOVERSION=1.18.4
ARG GOARCH=amd64
ARG NODEVERSION=16

# Install build essentials
RUN apt-get update && apt-get install -y \
    build-essential \
    bash \
    curl \
    git \
    lsb-release

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

# Install golang-migrate
RUN curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add - && \
    echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ $(lsb_release -sc) main" > /etc/apt/sources.list.d/migrate.list && \
    apt-get update && apt-get install -y migrate 

FROM base AS gobooking
ARG WORKDIR
WORKDIR ${WORKDIR} 
COPY ./go.mod ./go.mod ./
COPY . ./