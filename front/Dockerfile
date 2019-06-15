FROM node:11.15.0-alpine as builder

WORKDIR /home/app

COPY yarn.lock ./
COPY package.json ./
RUN yarn install

COPY .eslintrc.js .browserslistrc babel.config.js postcss.config.js ./
COPY src ./src
COPY public ./public
RUN yarn run build

FROM nginx:1.16-alpine

COPY --from=builder /home/app/dist/ /usr/share/nginx/html
