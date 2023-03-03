FROM node:19

COPY entrypoint.sh /entrypoint.sh
COPY .prettierrc /default-prettier-config.json

ENTRYPOINT ["/entrypoint.sh"]