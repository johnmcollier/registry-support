FROM nginx:stable-alpine

COPY nginx.conf.template /etc/nginx/nginx.conf.template

COPY entrypoint.sh /

RUN set -x ; \
  adduser -u 82 -D -S -G root www-data && exit 0 ; exit 1

RUN touch /var/run/nginx.pid

# Modify the permissions on the necessary files to allow the container to properly run as a non-root UID
RUN chown -R www-data:root /var/run/nginx.pid && \
  chown -R www-data:root /var/cache/nginx && \
  chown www-data:root /etc/nginx/conf.d /etc/nginx/nginx.conf /etc/nginx/nginx.conf.template
RUN chmod g+rwx /var/run/nginx.pid && \
  chmod -R g+rwx /var/cache/nginx && chmod -R g+rwx /etc/nginx
  
USER www-data

EXPOSE 5001    
ENTRYPOINT ["/entrypoint.sh"]
CMD ["nginx", "-g", "daemon off;"]