FROM alpine:3

RUN apk add --no-cache poppler-utils poppler-data

COPY ./bin/pdf2text /pdf2text

RUN chmod +x /pdf2text

ENV HOST "0.0.0.0"
ENV PORT 80

EXPOSE 80

CMD ["/pdf2text"]
