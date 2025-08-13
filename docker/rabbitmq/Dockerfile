FROM rabbitmq:3.12-management

# Copy plugin file into RabbitMQ plugins folder
COPY plugins/rabbitmq_delayed_message_exchange-3.12.0.ez /plugins/

# Enable the delayed message plugin
RUN rabbitmq-plugins enable --offline rabbitmq_delayed_message_exchange
