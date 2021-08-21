FROM alpine:3.12
RUN mkdir /app
COPY ./config /app/
COPY ./main /app/
RUN chmod +x /app/main
WORKDIR "/app"
CMD ["./main"]