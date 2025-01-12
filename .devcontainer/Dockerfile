FROM ubuntu:22.10

ARG NODE_VERSION="18.12.1"
ARG GO_VERSION="1.20.1"
ARG PNPM_VERSION="8"

ARG TARGETARCH
ARG USERNAME=tsuwari
ARG USER_UID=1000
ARG USER_GID=$USER_UID
RUN groupadd --gid $USER_GID $USERNAME && \
    useradd --uid $USER_UID --gid $USER_GID -m $USERNAME && \
    usermod -s /bin/bash $USERNAME

RUN apt update && \
    apt install -y build-essential git protobuf-compiler curl wget sudo
RUN curl -fsSL https://deb.nodesource.com/setup_18.x | bash - && \
    apt install -y nodejs
RUN wget https://go.dev/dl/go$GO_VERSION.linux-$TARGETARCH.tar.gz && \
    tar -C /usr/local -xzf go$GO_VERSION.linux-$TARGETARCH.tar.gz && \
    rm go$GO_VERSION.linux-$TARGETARCH.tar.gz

RUN npm i -g pnpm@$PNPM_VERSION @go-task/cli

RUN echo "${USERNAME} ALL=(ALL) NOPASSWD:ALL" > /etc/sudoers

# switch user on build time
RUN echo "${USERNAME}:tsuwari" | chpasswd
USER ${USERNAME}
ENV PATH="$PATH:/usr/local/go/bin:$PATH:/home/$USERNAME/go/bin"
RUN SHELL=/bin/bash pnpm setup
# install neede go tools
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest && \
    go install github.com/twitchtv/twirp/protoc-gen-twirp@latest && \
    go install github.com/cespare/reflex@latest && \
    go install github.com/pressly/goose/v3/cmd/goose@latest
