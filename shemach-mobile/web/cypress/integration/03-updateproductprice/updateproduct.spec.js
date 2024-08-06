/// <reference types="cypress" />


describe('Agri Info Price Update Testing', () => {
    // beforeEach(() => {
    //   cy.visit('http://localhost:3000/')
    // })


        it('Login and update product price', () => {
            cy.visit('/')
            // cy.contains('h1', 'todos')
            cy.get('[data-cy="email"]').type('fantahunfekadu1@gmail.com')
            cy.get('[data-cy="password"]').type('admin')
            //cy.get('.loginbutton').click()
          //    cy.get('.loginbutton', {timeout:10000}).click()
          
          // cy.wait(5000);
          //cy.contains('Login').click()
          cy.get('[data-cy="submit"]').click()
          //cy.wait(5000)
          cy.get('.list-group > :nth-child(2)').click()
          cy.get('[data-cy="editpricebutton"]').click()
          cy.get('[data-cy="currentproprice"]').type(5000)
          cy.get('[data-cy="updatepricebutton"]').click()
          cy.wait(5000)
          cy.get('.returnback').click()

    })

 





it('test /api/products with simple intercept', ()=>{

    cy.visit('/')
  
    // cy.intercept({
    //     path : '/login'
  
    // }).as('products')
    cy.intercept('GET', '/api/products').as('products')
  
    // cy.get('#email').type('fantahunfekadu1@gmail.com')
    // cy.get('#password').type('admin')
    cy.get('[data-cy="email"]').type('fantahunfekadu1@gmail.com')
    cy.get('[data-cy="password"]').type('admin')
    cy.get('[data-cy="submit"]').click()
    cy.wait('@products').then((inter) =>{
      console.log(inter)
      console.log(inter.response.body.products)
      const response = inter.response.body.products;
      expect(response).to.have.length(11);
        // cy.log(JSON.stringify(inter))
        // console.log(JSON.stringify(inter.body))
       })
  
  
  })
  
  
  it('spy and stub (mocking) price update with intercept test with dynamic fixture (stubbing)', ()=>{
    cy.visit('/')
    cy.intercept('PUT', '/api/infoadmin/product', {fixture: 'propriceupdate.json'}).as('productprice')
  
    cy.get('[data-cy="email"]').type('fantahunfekadu1@gmail.com')
    cy.get('[data-cy="password"]').type('admin')
    cy.get('[data-cy="submit"]').click()
    cy.get('.list-group > :nth-child(1)').click()
    cy.get('[data-cy="editpricebutton"]').click()
    cy.get('[data-cy="updatepricebutton"]').click()
    // cy.wait(5000)
    cy.wait('@productprice').then((inter) =>{
      console.log(inter)
  
      // console.log(inter.response.body.products)
      // const response = inter.response.body.products;
  
      // expect(response).to.have.length(9);
      
    })
    // cy.get('@productprice').should('not.be.null')
    // cy.wait('@products').its('response.statusCode').should('eq', 200)
  
    // cy.get('li.selected').should('have.length', 3)
  
   
  })
  
   
})