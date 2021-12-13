FROM node:14.15.4-alpine3.12

WORKDIR /usr/src/app

COPY package.json .
COPY package-lock.json .

RUN npm install

COPY . .

CMD npm run start:dev