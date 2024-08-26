FROM alpine:latest

COPY interfaces /etc/network

RUN apk add --no-cache agetty openrc \
    && rc-update add local default \
    && echo rc_logger=\"NO\" >> /etc/rc.conf \
    && echo rc_verbose=\"NO\" >> /etc/rc.conf 

COPY img.start /etc/local.d 

COPY inittab /etc

