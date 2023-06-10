# Stage 1
FROM node:18.10

WORKDIR /app

COPY ["package.json", "pnpm-lock.yaml", "./"]

RUN npm install

COPY . .

EXPOSE 4200

CMD [ "npm", "run", "ng", "serve"]

