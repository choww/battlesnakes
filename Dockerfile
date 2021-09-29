# Pick a good base image for your language stack
FROM golang:1.16-alpine3.14

# set the app directory that your stack expects
WORKDIR /usr/src/app

# copy in your source files
COPY src .

# run any compilation steps your stack needs
RUN go install \
  && go build 

# port 8080 is set by default in your snake's vars.yaml
EXPOSE 8080

# set the command to start your API server
CMD ["./starter-snake-go"]
