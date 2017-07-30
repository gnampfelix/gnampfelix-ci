FROM alpine:3.6

COPY build.sh /bin/build.sh

RUN apk add --no-cache bash

ENTRYPOINT ["/bin/build.sh"]
