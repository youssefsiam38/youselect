FROM node:14.6.0-alpine

WORKDIR /app

COPY ./package.json ./package-lock.json ./

RUN npm cache clean -f

RUN npm install -g

RUN npm i serve -g

COPY . .

# RUN npm run build

EXPOSE 80

# CMD ["npm", "start"]

CMD ["sh", "./run.sh"]