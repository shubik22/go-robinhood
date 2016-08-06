'use strict'

const path = require("path");

module.exports = {
  entry: {
    app: [
      './web/index.js'
    ]
  },

  output: {
    path: path.resolve(__dirname + '/static_files'),
    publicPath: "/assets/",
    filename: '[name].bundle.js',
  },

  module: {
    loaders: [
      {
        test:    /\.html$/,
        exclude: /node_modules/,
        loader:  'file?name=[name].[ext]',
      },

      {
        test:    /\.elm$/,
        exclude: [/elm-stuff/, /node_modules/],
        loader:  'elm-webpack',
      }
    ],

    noParse: /\.elm$/,
  },

  devServer: {
    stats: { colors: true },
    contentBase: './static_files',
    proxy: {
      "/leaderboard": {
        target: 'http://localhost:4000',
        ignorePath: false
      },
      "/assets/style.css": {
        target: 'http://localhost:4000/style.css',
        ignorePath: true
      },
      "/assets/milligram.min.css": {
        target: 'http://localhost:4000/milligram.min.css',
        ignorePath: true
      },
      "/assets/favicon.ico": {
        target: 'http://localhost:4000/favicon.ico',
        ignorePath: true
      },
      "/assets/logo.png": {
        target: 'http://localhost:4000/logo.png',
        ignorePath: true
      },
      "/assets/open-iconic.svg": {
        target: 'http://localhost:4000/open-iconic.svg',
        ignorePath: true
      },
      "/assets/open-iconic.woff": {
        target: 'http://localhost:4000/open-iconic.woff',
        ignorePath: true
      }
    }
  }
};
