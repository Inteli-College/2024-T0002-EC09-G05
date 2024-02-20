FROM rabbitmq:3.12-management

ENV RABBITMQ_SERVER_ADDITIONAL_ERL_ARGS="-rabbitmq_management load_definitions \"/etc/rabbitmq/definitions.json\""

ADD --chown=rabbitmq definitions.json /etc/rabbitmq/
ADD --chown=rabbitmq rabbitmq.conf /etc/rabbitmq/
ADD --chown=rabbitmq  ./certs/ /etc/rabbitmq/

RUN rabbitmq-plugins enable rabbitmq_mqtt rabbitmq_top

# RUN rabbitmqctl add_user publisher_user publisher_password
# RUN rabbitmqctl add_user publisher_user publisher_password
# RUN rabbitmqctl add_user publisher_user publisher_password


EXPOSE 5671 15672 1883 8883 15671 61613 61614 15674

CMD ["rabbitmq-server"]