# Alpine base images
FROM alpine:3.7

# Maintainer images
LABEL maintainer="Reja Nur Rochmat rezanurrochmat3@gmail.com"

# Add main executable files in container
ADD main /

# Expose port outside container
EXPOSE 8000

# Exceute binary go
CMD ["/main"]