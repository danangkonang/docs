# FROM php:7.2.12-apache
# COPY www/ /var/www/html/
# COPY php.ini /usr/local/etc/php/

# EXPOSE 80

FROM php:8.2.3-apache
WORKDIR /var/www/html
RUN mkdir -p /var/www/html/uploads
COPY index.php index.php
RUN chown -R www-data:www-data /var/www/html/
# COPY src/ src
EXPOSE 80