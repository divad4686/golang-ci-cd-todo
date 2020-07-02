FROM flyway/flyway

# Set the Current Working Directory inside the container

COPY *.sql /flyway/sql/

ENTRYPOINT ["flyway"]
CMD ["migrate"]