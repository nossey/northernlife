const awsServerlessExpress = require('aws-serverless-express');
const express = require('express');
const app = express();
const config = require('./nuxt.config');
const { Nuxt, Builder } = require('nuxt');

async function initApp () {
  const nuxt = new Nuxt(config);

  if (config.dev){
    const builder = new Nuxt(config);
    await builder.build();
  } else {
    await nuxt.ready();
  }
  app.use(nuxt.render);
  return app;
};

let server = undefined;
const binaryMimeTypes = [
  'application/javascript',
  'application/json',
  'application/octet-stream',
  'application/xml',
  'font/eot',
  'font/opentype',
  'font/otf',
  'image/jpeg',
  'image/png',
  'image/svg+xml',
  'text/comma-separated-values',
  'text/css',
  'text/html',
  'text/javascript',
  'text/plain',
  'text/text',
  'text/xml'
];

exports.render = (event, context) => {
  initApp().then((app) => {
    if (server === undefined) {
       server = awsServerlessExpress.createServer(app, null, binaryMimeTypes);
    }
    awsServerlessExpress.proxy(server, event, context);
  });
}
