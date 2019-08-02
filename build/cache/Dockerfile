ARG GOPATH=/go
ARG ROOT=github.com/lichuan0620/nirvana-practice
ARG WORKSPACE=${GOPATH}/src/${ROOT}
ARG IMAGE_PREFIX=practice-
ARG TARGET=cache

FROM golang:1.12.7 as builder

ARG VERSION
ARG COMMIT
ARG BRANCH

ENV GOPATH $GOPATH
ARG ROOT
ARG WORKSPACE
ARG IMAGE_PREFIX
ARG TARGET

WORKDIR ${WORKSPACE}

COPY client ${WORKSPACE}/client
COPY cmd ${WORKSPACE}/cmd
COPY pkg ${WORKSPACE}/pkg
COPY vendor ${WORKSPACE}/vendor

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64									\
	go build -v -o ./bin/${IMAGE_PREFIX}${TARGET}							\
  	-ldflags "-s -w										               		\
            -X  "${ROOT}"/pkg/info.version="${VERSION}"						\
            -X  "${ROOT}"/pkg/info.commit="${COMMIT}"						\
            -X  "${ROOT}"/pkg/info.branch="${BRANCH}						\
  	./cmd/${TARGET}    

FROM cargo.caicloudprivatetest.com/caicloud/debian:jessie

ARG WORKSPACE
ARG IMAGE_PREFIX
ARG TARGET

COPY --from=0 									\
	${WORKSPACE}/bin/${IMAGE_PREFIX}${TARGET} 	\
	/usr/local/bin

ENTRYPOINT ["practice-cache"]
