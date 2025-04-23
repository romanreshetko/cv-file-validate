FROM ubuntu:latest
LABEL authors="rvres"

ENTRYPOINT ["top", "-b"]