ARG GCR_MIRROR=gcr.io/
FROM ${GCR_MIRROR}distroless/static:nonroot
LABEL org.opencontainers.image.source https://github.com/norskhelsenett/ror
WORKDIR /

COPY dist/ror-ms-auth /bin/ror-ms-auth
ENTRYPOINT ["/bin/ror-ms-auth"]
