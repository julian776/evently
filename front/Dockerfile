# Stage 1
FROM node

WORKDIR /app

COPY ["package.json", "pnpm-lock.yaml", "./"]

RUN npm install

COPY . .

EXPOSE 4200

CMD [ "npm", "run", "ng", "serve"]

