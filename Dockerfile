FROM golang:latest

WORKDIR /go/src/github.com/rajeshreddyt/ParserSchedulerService

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

CMD ["ParserSchedulerService"]


# docker build -t parser-and-scheduler-service .
# docker run -p 50051:50051 parserschedulerservice
# push to docker hub
# aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin 626995068279.dkr.ecr.us-east-1.amazonaws.com
# docker build -t parser-and-scheduler-service .
# docker tag parser-and-scheduler-service:latest 626995068279.dkr.ecr.us-east-1.amazonaws.com/parser-and-scheduler-service:latest
# docker push 626995068279.dkr.ecr.us-east-1.amazonaws.com/parser-and-scheduler-service:latest
# docker manifest create  626995068279.dkr.ecr.us-east-1.amazonaws.com/parser-and-scheduler-service:latest  626995068279.dkr.ecr.us-east-1.amazonaws.com/parser-and-scheduler-service:latest-linux/amd64
