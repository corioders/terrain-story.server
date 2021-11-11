FROM nginx:alpine

COPY ./nginx.conf /etc/nginx/conf.d/default.conf
COPY ./nginx-server-configs /etc/nginx/conf.d/nginx-server-configs

COPY ./out /data/www

CMD sh -c "nginx -g 'daemon off;'"