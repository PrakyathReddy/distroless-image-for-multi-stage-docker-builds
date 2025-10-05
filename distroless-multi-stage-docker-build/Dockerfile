############## BASE IMAGE ##############

FROM ubuntu AS build

RUN apt-get update && apt-get install -y golang-go

ENV GO111MODULE=off

COPY . . 

RUN CGO_ENABLED=0 go build -o /app .

################ SECOND STAGE ##############

FROM scratch
# minimalist distroless image
# works for Go because Go language applications do not need runtime but won't work for python apps since python needs a python runtime to run

COPY --from=build /app /app
# Copy the compiled binary from build stage - from /app dir in build stage to /app in this stage

ENTRYPOINT ["/app"]
# setting the entrypoint for the container to run the binary


 