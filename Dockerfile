FROM alpine:3

RUN set -eux && sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk add --no-cache poppler-utils
COPY pdf2text /pdf2text
RUN chmod +x /pdf2text

CMD ["/pdf2text"]
