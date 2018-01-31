ARG arch
FROM multiarch/alpine:${arch}-edge

COPY ./expino-grafana-proxy /expino-grafana-proxy

ENV GRAFANAURL=""
ENV USERNAME=""
ENV PASSWORD=""

CMD /expino-grafana-proxy