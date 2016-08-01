'use strict';

require('./elm-app/index.html');

const Elm = require('./elm-app/Main.elm');
const mountNode = document.getElementById('main');

const app = Elm.Main.embed(mountNode);

