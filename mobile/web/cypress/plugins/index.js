/// <reference types="cypress" />
// ***********************************************************
// This example plugins/index.js can be used to load plugins
//
// You can change the location of this file or turn off loading
// the plugins file with the 'pluginsFile' configuration option.
//
// You can read more here:
// https://on.cypress.io/plugins-guide
// ***********************************************************

// This function is called when a project is opened or re-opened (e.g. due to
// the project's config changing)

/**
 * @type {Cypress.PluginConfig}
 */

// cypress/plugins/index.js
module.exports = (on, config) => {
  require('@cypress/code-coverage/task')(on, config)
  // tell Cypress to use .babelrc file
  // and instrument the specs files
  // only the extra application files will be instrumented
  // not the spec files themselves
  on('file:preprocessor', require('@cypress/code-coverage/use-babelrc'))

  return config
}
// eslint-disable-next-line no-unused-vars
// module.exports = (on, config) => {
//   require('@cypress/code-coverage/task')(on, config)
//   on('file:preprocessor', require('@cypress/code-coverage/use-browserify-istanbul'))
//   return config
// }

// module.exports = (on, config) => {
//   require('@cypress/code-coverage/task')(on, config)
//   on('file:preprocessor', require('@cypress/code-coverage/use-babelrc'))
//   return config
// }
