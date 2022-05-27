FROM node:10.13-alpine as build
WORKDIR /root
COPY ./package*.json ./
RUN npm ci --cache .npm
COPY . . 
RUN npm run build -o app

# multistage build...
# this runs a lil nginx instance to serve the app
FROM nginx:alpine
EXPOSE 3000
COPY --from=build /root/build /usr/share/nginx/html
COPY nginx.conf /etc/nginx/nginx.conf