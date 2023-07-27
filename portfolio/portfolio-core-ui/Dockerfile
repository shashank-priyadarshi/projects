#stage 1
FROM node:latest as node
WORKDIR /app
#COPY . .
#RUN npm install
#RUN npm run build --prod
COPY package.json package-lock.json ./
RUN npm i --f --legacy-peer-deps
COPY . .
EXPOSE 4200

#stage 2
#FROM nginx:alpine
#COPY --from=node /app/dist/resume /usr/share/nginx/html
