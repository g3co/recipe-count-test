FROM scratch
COPY ./bin /bin
ENTRYPOINT ["/bin/calc"]
