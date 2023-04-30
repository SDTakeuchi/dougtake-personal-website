FROM node:lts-alpine3.16

WORKDIR /usr/app/

RUN apk update

COPY ./productlog-nextjs-theme/ .

RUN yarn install

# CMD [ "yarn", "run", "dev" ]
# CMD [ "yarn", "build", "&&", "yarn", "start"]
