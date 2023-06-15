# Fetching the latest node image on alpine linux
FROM node:alpine AS development

WORKDIR /react-app

COPY ./package.json  /react-app
COPY ./package-lock.json  /react-app
RUN npm install

# Copying all the files in our project
COPY . .
