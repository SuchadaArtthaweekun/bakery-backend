FROM postgres

ENV POSTGRES_DB my_db
ENV POSTGRES_USER postgres
ENV POSTGRES_PASSWORD password

ADD product.sql /docker-entrypoint-initdb.d
ADD user.sql /docker-entrypoint-initdb.d