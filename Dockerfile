# VERSION 0.2.0
# AUTHOR:         Nicolas Lamirault <nicolas.lamirault@gmail.com>
# DESCRIPTION:    zeiot/jarvis-bot

FROM resin/armv7hf-debian:jessie
MAINTAINER Nicolas Lamirault <nicolas.lamirault@gmail.com>

ENV JARVIS_BOT_VERSION 0.2.0

ADD https://bintray.com/artifact/download/zeiot/jarvis/jarvis-bot-${JARVIS_BOT_VERSION}_linux_arm /jarvis-bot

RUN chmod +x /jarvis-bot

VOLUME [ "/etc/jarvis" ]

ENTRYPOINT [ "/jarvis-bot" ]
