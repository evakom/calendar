# run from project root
FROM postgres:9.6.16-alpine
LABEL "Author"="Eugene Klimov"
RUN apk add coreutils
ADD ./sql/ /docker-entrypoint-initdb.d/
