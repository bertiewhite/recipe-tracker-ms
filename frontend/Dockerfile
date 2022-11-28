## STAGE 1: Build
FROM node:slim as build

WORKDIR /app

COPY package.json ./
COPY package-lock.json ./
RUN npm install
COPY . ./
RUN npm run build


### STAGE 2: Run
FROM nginx:stable
COPY --from=build /app/dist /usr/share/nginx/html