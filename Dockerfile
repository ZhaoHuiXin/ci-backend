FROM alpine:3.10
ENV TZ=Asia/Shanghai

WORKDIR /opt/srv/bin/
EXPOSE 9205

COPY ./bin/backend .

RUN chmod +x /opt/srv/bin/backend

CMD /opt/srv/bin/backend
