// ***********************************************
// This example commands.js shows you how to
// create various custom commands and overwrite
// existing commands.
//
// For more comprehensive examples of custom
// commands please read more here:
// https://on.cypress.io/custom-commands
// ***********************************************
//
//
// -- This is a parent command --
// Cypress.Commands.add('login', (email, password) => { ... })
//
//
// -- This is a child command --
// Cypress.Commands.add('drag', { prevSubject: 'element'}, (subject, options) => { ... })
//
//
// -- This is a dual command --
// Cypress.Commands.add('dismiss', { prevSubject: 'optional'}, (subject, options) => { ... })
//
//
// -- This will overwrite an existing command --
// Cypress.Commands.overwrite('visit', (originalFn, url, options) => { ... })
Cypress.Commands.add('logininfo', (infoemail, password) => { 
    cy.visit('/')
    cy.get('[data-cy="email"]').type('fantahunfekadu1@gmail.com')
    cy.get('[data-cy="password"]').type('admin')
    // cy.get('[data-cy=submit]').click()
    cy.get('[data-cy="submit"]').click()

 })


 Cypress.Commands.add('loginsuper', (superemail, password) => { 
    cy.visit('/')
    cy.get('[data-cy="email"]').type('samuaeladadnew@outlook.com')
    cy.get('[data-cy="password"]').type('admin')
    // cy.get('[data-cy=submit]').click()
    cy.get('[data-cy="submit"]').click()

 })