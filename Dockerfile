FROM daocloud.io/golang

RUN mkdir /go/src/pwa-blog-server

ADD . /go/src/pwa-blog-server

WORKDIR  /go/src/pwa-blog-server

RUN wget http://golangtc.com/static/download/packages/golang.org.x.net.tar.gz
RUN tar xvf golang.org.x.net.tar.gz
RUN mv golang.org /go/src

RUN go get github.com/astaxie/beego
RUN go get github.com/PuerkitoBio/goquery

RUN go install pwa-blog-server

ENTRYPOINT /go/bin/pwa-blog-server

EXPOSE 8080
