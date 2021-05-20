FROM golang:latest as builder
ADD ./app /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /main .

FROM scratch
ARG VERSION
ARG COLOUR
ENV VERSION=${VERSION}
ENV COLOUR=${COLOUR}
# RUN apk --no-cache add ca-certificates
COPY --from=builder /main ./
COPY --from=builder /app/html ./html 
COPY --from=builder /app/css ./css
# RUN chmod +x ./main
# ENTRYPOINT ["/bin/ash", "-c", "./main"]
ENTRYPOINT ["./main"]
EXPOSE 80