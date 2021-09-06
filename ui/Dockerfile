FROM node:16-stretch-slim

ADD ./ /src

WORKDIR /src
RUN npm install
RUN npm run build

ENV NODE_ENV=production
ENV HOST=0.0.0.0
ENV PORT=4000

EXPOSE 4000

CMD ["node", "build/index.js"]
