/// <reference types="cypress" />



describe('Agri Info View Product Info Testing', () => {
    // beforeEach(() => {
    //   cy.visit('http://localhost:3000/')
    // })

    it('View product info', ()=>{
        cy.visit('/')
        // cy.intercept('GET', '/api/products', {fixture: 'product.json'}).as('products')
      
        cy.get('[data-cy="email"]').type('fantahunfekadu1@gmail.com')
        cy.get('[data-cy="password"]').type('admin')
        cy.get('[data-cy="submit"]').click()
        cy.get('.list-group > :nth-child(1)').click()
        // cy.wait('@products').then((inter) =>{
        //   console.log(inter)
      
        //   console.log(inter.response.body.products)
        //   const response = inter.response.body.products;
      
        //   expect(response).to.have.length(9);
          
        // })
        // cy.get('@products').should('not.be.null')
        // cy.wait('@products').its('response.statusCode').should('eq', 200)
      
        // cy.get('li.selected').should('have.length', 3)
      
       
      })



it('Spy and stub view product information with fixture', ()=>{
  cy.visit('/')
  cy.intercept('GET', '/api/products', {fixture: 'viewprodinfo.json'}).as('productview')

  cy.get('[data-cy="email"]').type('fantahunfekadu1@gmail.com')
  cy.get('[data-cy="password"]').type('admin')
  cy.get('[data-cy=submit]').click()
  cy.wait('@productview').then((inter) =>{
    console.log(inter)

    console.log(inter.response.body.products)
    const response = inter.response.body.products;

    expect(response).to.have.length(2);
    
  })
  cy.get('@productview').should('not.be.null')
  cy.get('.list-group > :nth-child(1)').click()
  cy.get('.list-group').should('have.length', 1)
  cy.get('.list-group-item').should('have.length', 2)

  // cy.wait('@products').its('response.statusCode').should('eq', 200)

  // cy.get('li.selected').should('have.length', 3)

 
})
 
   
      

})
