# distroless-image-for-multi-stage-docker-builds

Both multi-stage docker builds and distroless images are efforts to reduce the size of docker image. 

Especially useful in scenarios where security, efficiency and minimalism are top priorities.

However, lack of an interactive shell makes debugging difficult. Also, since dependencies are fixed, we cannot add more on the fly. It's immutable.

So, need to choose only in scenarios where we are confident that the app won't break and is well tested.