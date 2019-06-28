FROM alpine:3.4

RUN apk -U add ca-certificates
RUN apk add --no-cache bash

EXPOSE 8080

ADD . /testApp
Add ./templates /templates
CMD ["/testApp/app"]

