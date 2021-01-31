FROM golang
WORKDIR /go/src/gofiber
COPY . .
RUN go install
EXPOSE 3001
CMD /go/bin/gofiber