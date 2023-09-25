FROM debian:bookworm-slim

COPY pipeLog_0_0_amd64 ./pipeLog_0_0_amd64

RUN dpkg-deb --build ./pipeLog_0_0_amd64/

RUN dpkg -i pipeLog_0_0_amd64.deb

ENTRYPOINT ["pipeLog"]