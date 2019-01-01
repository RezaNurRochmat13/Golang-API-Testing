# Alpine base images
FROM alpine:3.8

# Maintainer images
LABEL maintainer="Reja Nur Rochmat rezanurrochmat3@gmail.com"

# Add main executable files in container
ADD main /

# Expose port outside container
EXPOSE 8080

EXPOSE 3306

# Exceute binary go
CMD ["/main"]