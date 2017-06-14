FROM scratch
ADD present8 /
WORKDIR /data
ENTRYPOINT ["/present8"]
