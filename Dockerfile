FROM golang:1.22-alpine

WORKDIR /Users/eugenevorobiov/Desktop/work/studing/SES4Case

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /SES4Case
CMD /SES4Case

EXPOSE 8080