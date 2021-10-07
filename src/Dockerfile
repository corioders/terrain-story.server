FROM nginx:alpine

COPY ./nginx.conf /etc/nginx/conf.d/default.conf
COPY ./out /data/www

CMD sh -c "nginx -g 'daemon off;'"