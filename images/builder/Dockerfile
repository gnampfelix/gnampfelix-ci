FROM alpine:3.6

ENV GOPATH /go

RUN apk add --no-cache bash wget tar ca-certificates git

# Install glibc for golang to work
RUN wget -q -O /etc/apk/keys/sgerrand.rsa.pub https://raw.githubusercontent.com/sgerrand/alpine-pkg-glibc/master/sgerrand.rsa.pub \
    && wget -q https://github.com/sgerrand/alpine-pkg-glibc/releases/download/2.25-r0/glibc-2.25-r0.apk \
    && apk add glibc-2.25-r0.apk

# Install golang
RUN mkdir $GOPATH \
    && mkdir $GOPATH/src && mkdir $GOPATH/pkg && mkdir $GOPATH/bin \
    && wget -q https://redirector.gvt1.com/edgedl/go/go1.9.2.linux-amd64.tar.gz \
    && tar -C /usr/local -xzfgo1.9.2.linux-amd64.tar.gz

ENV PATH=${PATH}:/usr/local/go/bin:${GOPATH}/bin

# Install testing utils
RUN go get github.com/onsi/ginkgo/ginkgo

COPY build.sh /bin/build.sh

ENTRYPOINT ["/bin/build.sh"]
