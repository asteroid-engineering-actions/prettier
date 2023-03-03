FROM node:19

COPY entrypoint.sh /entrypoint.sh
COPY .prettierrc /default-prettier-config.json

RUN node --version && npm --version && npm install -g prettier

#ENTRYPOINT ["/entrypoint.sh"]