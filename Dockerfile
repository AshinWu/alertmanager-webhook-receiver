FROM scratch
MAINTAINER AshinWoo <https://github.com/AshinWu>

WORKDIR /
ADD build/alertmanager-webhook-receiver /

ENTRYPOINT [ "./alertmanager-webhook-receiver" ]
EXPOSE 8090
