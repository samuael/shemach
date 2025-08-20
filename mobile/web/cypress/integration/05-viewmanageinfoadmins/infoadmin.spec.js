/// <reference types="cypress" />



describe('Agri Info View Info Info Admins', () => {
    // beforeEach(() => {
    //   cy.visit('http://localhost:3000/')
    // })

    it('View  info Admins', ()=>{
        cy.visit('/')
        // cy.intercept('GET', '/api/products', {fixture: 'product.json'}).as('products')
      
        cy.get('[data-cy="email"]').type('samuaeladadnew@outlook.com')
        cy.get('[data-cy="password"]').type('admin')
        cy.get('[data-cy="submit"]').click()
        cy.get('[data-cy="adminlink"]').click()
        cy.get('h4').should('have.text', 'Admin List')
        cy.get('h4').should('be.visible')
        cy.get('.description').should('have.text', 'Please click on a Admin to view full profile...')
        cy.get(':nth-child(1) > .row > .col-sm-7 > .Name').should('contain', 'Name :Samuael')
        // cy.get('.list-group')
        // cy.get('.list-group > :nth-child(1)').click()
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



it('Spy and stub view info admins with fixture', ()=>{
  cy.visit('/')
  cy.intercept('GET', '/api/infoadmins', {fixture: 'infoadmins.json'}).as('infoadmins')

  cy.get('[data-cy="email"]').type('samuaeladadnew@outlook.com')
  cy.get('[data-cy="password"]').type('admin')
  cy.get('[data-cy=submit]').click()
  cy.get('[data-cy="adminlink"]').click()
  cy.wait('@infoadmins').then((inter) =>{
    console.log(inter)

    // console.log(inter.response.body.products)
    const response = inter.response.body;

    expect(response).to.have.length(2);
    
  })
  cy.get('@infoadmins').should('not.be.null')
//   cy.get('.list-group > :nth-child(1)').click()
//   cy.get('.list-group').should('have.length', 1)
//   cy.get('.list-group-item').should('have.length', 2)

  // cy.wait('@products').its('response.statusCode').should('eq', 200)

  // cy.get('li.selected').should('have.length', 3)

 
})
 
   
      

})
