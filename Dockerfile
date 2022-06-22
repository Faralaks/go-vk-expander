FROM golang:1.16-alpine as builder

ENV GOFLAGS="-mod=vendor"
ENV CGO_ENABLED=0

ADD . /build
WORKDIR /build

RUN cd app && go build -o /build/vk-expander

FROM scratch

ARG port=80



COPY --from=builder /build/vk-expander /srv/vk-expander

WORKDIR /srv
EXPOSE $port
ENTRYPOINT ["/srv/vk-expander"]