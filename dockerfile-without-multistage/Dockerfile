FROM ubuntu AS build 
# build is an alias for the image

RUN apt-get update && apt-get install -y golang-go

ENV GO111MODULE=off 
# disable Go modules 

COPY . . 
# first . is the source, second . is the destination. Bascially means copy everything from current dir to current dir in the image

RUN CGO_ENABLED=0 go build -o /app .
# disabling CGO forces Go compiler to produce a fully static binary without any dynamic linking to C libraries. 
# this means the binary includes everything it needs to run independently, making it portable across systems without requiring OS level C dependencies
# -o /app tells the compiler to output the compiled binary to /app path inside the image

ENTRYPOINT ["/app"]
# execute the binary as part of entrypoint

 