FROM debian:10-slim AS builder

ARG VERSION="0.4.7"
ARG DOWNLOAD_URL="http://www.phontron.com/kytea/download/kytea-${VERSION}.tar.gz"
ARG EXTRACT_DIR="/var/tmp/kytea"
ARG INSTALL_DIR="/opt/kytea"

RUN set -x \
 && DEBIAN_FRONTEND=noninteractive apt-get -q -y update \
 && DEBIAN_FRONTEND=noninteractive apt-get -q -y install \
    build-essential \
    curl \
    tar \
    gzip \
 ;
RUN set -x \
 && mkdir -p "${EXTRACT_DIR}" \
 && curl -fsSL "${DOWNLOAD_URL}" \
  | tar -xzf - -C "${EXTRACT_DIR}" --strip-components=1 \
 ;

RUN set -x \
 && cd "${EXTRACT_DIR}" \
 && ./configure --prefix="${INSTALL_DIR}" \
 && make \
 && make install \
 && echo "${INSTALL_DIR}/lib" > /etc/ld.so.conf.d/kytea.conf \
 && rm -f /etc/ld.so.cache \
 && ldconfig \
 && ldconfig -p \
 ;


FROM golang:1.15-buster AS go-builder

COPY docker-entrypoint.go /var/tmp/

RUN set -x \
 && cd /var/tmp \
 && go build -o /docker-entrypoint docker-entrypoint.go


FROM gcr.io/distroless/cc-debian10

ARG INSTALL_DIR="/opt/kytea"
ENV KYTEA_DIR="${INSTALL_DIR}"

COPY --from=go-builder /docker-entrypoint /docker-entrypoint
COPY --from=builder "${INSTALL_DIR}" "${INSTALL_DIR}"
COPY --from=builder "/etc/ld.so.conf.d/kytea.conf" "/etc/ld.so.conf.d/kytea.conf"

ENTRYPOINT ["/docker-entrypoint"]
