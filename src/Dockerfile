FROM nginx:alpine

COPY ./nginx.conf /etc/nginx/conf.d/default.conf
COPY nginx-server-configs /etc/nginx

COPY ./out /data/www

CMD sh -c "nginx -g 'daemon off;'"