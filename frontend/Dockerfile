FROM node:14.4.0

# Install yarn
RUN apt-get update && apt-get install -y curl apt-transport-https wget && \
    curl -sS https://dl.yarnpkg.com/debian/pubkey.gpg | apt-key add - && \
    echo "deb https://dl.yarnpkg.com/debian/ stable main" | tee /etc/apt/sources.list.d/yarn.list && \
    apt-get update && apt-get install -y yarn

WORKDIR /usr/src/app
COPY . .

# Install nuxt
RUN yarn install

RUN apt-get update
RUN apt-get install -y vim

EXPOSE 80

RUN yarn build
CMD [ "yarn", "start" ]
