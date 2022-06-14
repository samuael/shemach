/// <reference types="cypress" />



describe('Agri Info Testing', () => {
    // beforeEach(() => {
    //   cy.visit('http://localhost:3000/')
    // })


    it('Login loads with correct user credentials', () => {
        cy.visit('http://localhost:3000/')
       // cy.contains('h1', 'todos')
       cy.get('#email').type('fantahunfekadu1@gmail.com')
       cy.get('#password').type('admin')
       //cy.get('.loginbutton').click()
    //    cy.get('.loginbutton', {timeout:10000}).click()
    
    // cy.wait(5000);
     //cy.contains('Login').click()
     cy.get('[data-cy="submit"]').click()
     cy.get('.list-group')
     //cy.wait(5000)
    //  cy.get('.list-group > :nth-child(1)').click()
    //  cy.get('.editbutton').click()
    //  cy.get('#currentprice').type(5000)
    //  cy.get('.updatebutton').click()
    //  cy.wait(3000)
    //  cy.get('.returnback').click()

   })


   it('Login fails with Incorrect user credentials', () => {
      cy.visit('/')
      // cy.contains('h1', 'todos')
      cy.get('#email').type('fantahunfekadu@gmail.com')
      cy.get('#password').type('admin')
      //cy.get('.loginbutton').click()
    //    cy.get('.loginbutton', {timeout:10000}).click()

      // cy.wait(5000);
      //cy.contains('Login').click()
      cy.get('[data-cy="submit"]').click()
      // cy.wait(5000)
      // cy.get('.list-group')
      // cy.get('.list-group > :nth-child(1)').click()
      // cy.get('.editbutton').click()
      // cy.get('#currentprice').type(5000)
      // cy.get('.updatebutton').click()
      // cy.wait(3000)
      // cy.get('.returnback').click()

})
   

it('Spy and Stub POST/login with (fixture)', () => {
  cy.visit('/')
  cy.intercept('POST', '/api/login', {fixture: 'login.json'}).as('login')

  cy.get('[data-cy="email"]').type('fantahunfekadu1@gmail.com')
  cy.get('[data-cy="password"]').type('admin')
  // cy.get('[data-cy=submit]').click()
  cy.get('[data-cy="submit"]').click()
  cy.wait('@login').then((inter) =>{
    console.log(inter)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    

    // console.log(inter.response.body.products)
    const response = inter.response.body.user.firstname;
    const resmessage = inter.response.body.msg;
    const resstatuscode = inter.response.body.status_code;
    console.log(response);
    expect(resmessage).to.equal('authenticated')
    expect(response).to.not.equal('Fani')

    // expect(resstatuscode).should('eq', 200);

    // expect(response).to.have.length(1);
    
  })
  cy.get('@login').should('not.be.null')
  // cy.get('@login').should('contain', 'user')
  // cy.get('@login').its('response.status_code').should('eq', 200)

  // cy.get('li.selected').should('have.length', 3)  

})


      

})
