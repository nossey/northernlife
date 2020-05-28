'use strict';

const { Nuxt } = require('nuxt');
const serverless = require('serverless-http');
const express = require('express');
const nuxtConfig = require('./nuxt.config');

const config = {
    dev: false,
    ...nuxtConfig,
};

const nuxt = new Nuxt(config);
const app = express();

exports.render = serverless(app);
