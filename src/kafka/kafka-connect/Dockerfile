FROM confluentinc/cp-kafka-connect:latest

RUN confluent-hub install --no-prompt confluentinc/kafka-connect-mqtt:1.4.0

# ....

COPY ./my-connector-config.properties /etc/kafka-connect/

ENV CONNECT_BOOTSTRAP_SERVERS="kafka:9092"
ENV CONNECT_REST_PORT=8083
ENV CONNECT_GROUP_ID="kafka-connect"
ENV CONNECT_CONFIG_STORAGE_TOPIC="kafka-connect-configs"
ENV CONNECT_OFFSET_STORAGE_TOPIC="kafka-connect-offsets"
ENV CONNECT_STATUS_STORAGE_TOPIC="kafka-connect-status"
ENV CONNECT_KEY_CONVERTER="org.apache.kafka.connect.json.JsonConverter"
ENV CONNECT_VALUE_CONVERTER="org.apache.kafka.connect.json.JsonConverter"
ENV CONNECT_CONFIG_STORAGE_REPLICATION_FACTOR="1"
ENV CONNECT_OFFSET_STORAGE_REPLICATION_FACTOR="1"
ENV CONNECT_STATUS_STORAGE_REPLICATION_FACTOR="1"
ENV CONNECT_PLUGIN_PATH="/usr/share/java,/usr/share/confluent-hub-components"

CMD ["/etc/confluent/docker/run"]
