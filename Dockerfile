FROM alpine

COPY hello hello

EXPOSE 8080
CMD ["./hello"]
