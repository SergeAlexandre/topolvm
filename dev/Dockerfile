FROM ubuntu:18.04

RUN apt-get update && apt-get -y upgrade && apt-get install -y sudo make git wget curl

RUN cd /tmp && wget https://dl.google.com/go/go1.13.8.linux-amd64.tar.gz && sudo tar -xvf go1.*.linux-amd64.tar.gz && mv go /usr/local && mkdir /gopath/


ENV GOROOT=/usr/local/go
ENV GOPATH=/gopath
ENV PATH=$GOPATH/bin:$GOROOT/bin:$PATH

CMD ["sh", "-c", "tail -f /dev/null"]
