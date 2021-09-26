'use strict';

const path = require('path');
const escape = require('escape-string-regexp');

// Fix for emacs until this gets merged: 
// https://github.com/facebook/create-react-app/pull/10706/files
module.exports = function ignoredFiles(appSrc) {
  return [
    new RegExp(
      `^(?!${escape(
        path.normalize(appSrc + '/').replace(/[\\]+/g, '/'),
      )}).+/node_modules/`,
      'g',
    ),
    '**/.#*',
    '**/*~',
    '**/#*#',
  ];
};
