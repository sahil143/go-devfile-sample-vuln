FROM registry.access.redhat.com/ubi9/go-toolset:9.5-1737480393

COPY . .
RUN go mod download

RUN go build -o ./main

ENV PORT 8081
EXPOSE 8081

CMD [ "./main" ]
