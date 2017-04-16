FROM scratch
MAINTAINER Rafael Jesus <rafaelljesus86@gmail.com>
ADD cron-srv /cron-srv
ENV DATASTORE_URL="postgres://postgres:@docker/cron_srv?sslmode=disable"
ENV PORT="3000"
ENTRYPOINT ["/cron-srv"]
