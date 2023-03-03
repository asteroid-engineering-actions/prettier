FROM node:19

COPY entrypoint.sh /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]