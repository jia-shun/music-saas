FROM scratch

WORKDIR $GOPATH/src/music-saas
COPY . $GOPATH/src/music-saas

EXPOSE 8888
CMD ["./music-saas"]