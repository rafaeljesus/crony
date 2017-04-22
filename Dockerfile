FROM scratch
MAINTAINER Rafael Jesus <rafaelljesus86@gmail.com>
ADD crony /crony
ENV DATASTORE_URL="postgres://postgres:@docker/crony?sslmode=disable"
ENV PORT="3000"
ENTRYPOINT ["/crony"]
