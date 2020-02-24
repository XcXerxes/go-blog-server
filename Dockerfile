FROM scratch
WORKDIR $GOPATH/src/github.com/XcXerxes/go-blog-server
COPY . $GOPATH/src/github.com/XcXerxes/go-blog-server

EXPOSE 8000
CMD ["./go-blog-server"]
