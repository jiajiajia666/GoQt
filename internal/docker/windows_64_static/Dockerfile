FROM ubuntu:16.04
LABEL maintainer jiajiajia666

ENV USER user
ENV HOME /home/$USER
ENV GOPATH $HOME/work
ENV PATH /usr/lib/mxe/usr/bin:/usr/local/go/bin:$PATH
ENV QT_DOCKER true
ENV QT_MXE true
ENV QT_MXE_ARCH amd64
ENV QT_MXE_STATIC true

COPY --from=jiajiajia666/GoQt:linux /usr/local/go /usr/local/go
COPY --from=jiajiajia666/GoQt:linux $GOPATH/bin $GOPATH/bin
COPY --from=jiajiajia666/GoQt:linux $GOPATH/src/github.com/jiajiajia666/GoQt $GOPATH/src/github.com/jiajiajia666/GoQt
COPY --from=jiajiajia666/GoQt:windows_64_static_base /usr/lib/mxe /usr/lib/mxe

RUN $GOPATH/bin/qtsetup prep
RUN $GOPATH/bin/qtsetup check windows
RUN $GOPATH/bin/qtsetup generate windows
RUN $GOPATH/bin/qtsetup install windows
RUN cd $GOPATH/src/github.com/jiajiajia666/GoQt/internal/examples/widgets/line_edits && $GOPATH/bin/qtdeploy build windows && rm -rf ./deploy

RUN apt-get -qq update && apt-get --no-install-recommends -qq -y install ca-certificates git pkg-config