FROM ubuntu:18.04

RUN apt-get update && \
    apt-get install software-properties-common -y

RUN add-apt-repository ppa:longsleep/golang-backports && \
    apt-get update && \
    apt-get install golang-go nodejs npm curl jq git wget unzip -y

RUN wget https://releases.hashicorp.com/terraform/0.11.14/terraform_0.11.14_linux_386.zip && \
    unzip terraform_0.11.14_linux_386.zip && \
    mv terraform /usr/local/bin

RUN download_url=$(curl -s https://api.github.com/repos/go-swagger/go-swagger/releases/latest | \
    jq -r '.assets[] | select(.name | contains("'"$(uname | tr '[:upper:]' '[:lower:]')"'_amd64")) | .browser_download_url') && \
    curl -o /usr/local/bin/swagger -L'#' "$download_url" && \
    chmod +x /usr/local/bin/swagger

WORKDIR /root/go/src/github.com/billtrust/meetup-terraform-provider
COPY ./ .

RUN go get -v -insecure ./...
RUN go get github.com/gruntwork-io/terratest/modules/terraform

WORKDIR /root/go/src/github.com/billtrust/meetup-terraform-provider/server
RUN npm i

COPY ./container-scripts /scripts
RUN chmod +x /scripts/*

ENTRYPOINT node index.js