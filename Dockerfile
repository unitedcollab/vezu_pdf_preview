FROM mteeeur/go-vips 

WORKDIR /go/src/app

COPY . .

RUN make build

CMD ["./main"]
