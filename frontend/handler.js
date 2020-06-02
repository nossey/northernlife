const path = require('path');
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
  //app.use('/_nuxt', express.static(path.join(__dirname, '.nuxt', 'dist', 'client')));
  app.use('/static', express.static(path.join(__dirname, './static')));
  app.use((req, res) => 
    {
      //req.url = `${config.router.base}${req.url}`.replace('//', '/');
      nuxt.render(req, res)
    });
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
