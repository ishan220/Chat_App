#Build Stage
##base golangimage:tag
FROM --platform=linux/amd64 golang:1.21-alpine3.18 AS Builder

###all files copied to WORKDIR
WORKDIR /app

####copyinh all file from current directory to working directory first is current and 
###second is work dir after previous command dir is changed to work directory
COPY . .
#COPY start.sh .
#COPY migrate.linux-amd64 .

### -o <executable> <entrypoint file>
RUN tar -xvf migrate.linux-amd64.tar
RUN go build -o main main.go
####Run Stage
FROM alpine:3.18
WORKDIR /app

### . represents the work dire set from above command ,and app/main is the path from the builder stage
COPY --from=Builder /app/main .

###copying migrate binary fro, builder app dir to /app/migrate
COPY --from=Builder /app/migrate .
#COPY app.env .
#COPY start.sh .
COPY wait-for .
#COPY --from=Builder /app/start.sh .
###copying db migration scripts to "/app/migration" directory

COPY db/migration ./db/migration
##this id for readme to inform about the port exposed by the service
EXPOSE 8080

###command at last will run to execute the executable
CMD [ "/app/main" ]

##if entry point given CMD will become parameters
ENTRYPOINT [ "/app/start.sh" ]
