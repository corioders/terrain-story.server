FROM nginx:alpine

COPY ./nginx.conf /etc/nginx/conf.d/default.conf
ADD nginx-server-configs /etc/nginx/nginx-server-configs

COPY ./out /data/www

CMD sh -c "nginx -g 'daemon off;'"