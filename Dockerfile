FROM golang

RUN apt-get update && apt-get install -y curl build-essential \
&& apt-get install -y redis-server

WORKDIR /bookstudy
COPY . .

RUN chmod +x run.sh

COPY go.mod go.sum ./
RUN go mod download

RUN go build

CMD ["./run.sh"]