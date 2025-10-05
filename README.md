# distroless-image-for-multi-stage-docker-builds

Reduce docker image size to <1% of original docker image without compromising on functionality 

Both multi-stage docker builds and distroless images are efforts to reduce the size of docker image. 

Especially useful in scenarios where security, efficiency and minimalism are top priorities.

However, lack of an interactive shell makes debugging difficult. Also, since dependencies are fixed, we cannot add more on the fly. It's immutable.

So, need to choose only in scenarios where we are confident that the app won't break and is well tested.

We can create the binaries from Stage 1 with Ubuntu as the base image. Then, only use those binaries in stage 2 plus the executable - python/java,etc..

Step 1: Create calculator application in Go. $ go run calculator.go

Step 2: Containerise this and execute the container

Step 3: Create dockerfile without multi-stage

Step 4: Build docker image
% docker build -t simplecalculator .
. means use current directory as build context
% docker images    
REPOSITORY         TAG       IMAGE ID       CREATED         SIZE
simplecalculator   latest    1cbe04de39d9   2 minutes ago   954MB

Size of image: 954MB

Step 5: push image to docker hub
$ docker login
$ docker tag simplecalculator prakyathreddyk/sample-images:latest 

Step 6: Create multi-stage build with distroless image

Step 7: Build the multi-stage image
$ docker build -t simple-calculator-multi-stage-distroless .

$ docker images
REPOSITORY                                 TAG       IMAGE ID       CREATED          SIZE
simple-calculator-multi-stage-distroless   latest    a6c1aeaea036   32 seconds ago   3.19MB

IMAGE SIZE: 3.19 MB

