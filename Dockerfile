FROM golang
FROM node:12
WORKDIR /webspeedtest
COPY . .
RUN cd ./expressjs && npm i && cd ..
RUN npm i
EXPOSE 3000 3001
CMD ["npm", "start"]