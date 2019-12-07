FROM dockercore/golang-cross

LABEL maintainer="darrikonn@gmail.com"

ENV GORELEASER_VERSION=0.106.0
ENV GORELEASER_SHA=5828aa6837fef79df8952d5762ba7b58740a7503e254067a98a20360e75eca87

ENV GORELEASER_DOWNLOAD_FILE=goreleaser_Linux_x86_64.tar.gz
ENV GORELEASER_DOWNLOAD_URL=https://github.com/goreleaser/goreleaser/releases/download/v${GORELEASER_VERSION}/${GORELEASER_DOWNLOAD_FILE}

RUN  wget ${GORELEASER_DOWNLOAD_URL}; \
			echo "$GORELEASER_SHA $GORELEASER_DOWNLOAD_FILE" | sha256sum -c - || exit 1; \
			tar -xzf $GORELEASER_DOWNLOAD_FILE -C /usr/bin/ goreleaser; \
			rm $GORELEASER_DOWNLOAD_FILE;
			
CMD ["goreleaser", "-v"]
