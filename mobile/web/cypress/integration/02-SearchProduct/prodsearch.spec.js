/// <reference types="cypress" />



describe('Agri Info Product Search Testing', () => {
    // beforeEach(() => {
    //   cy.visit('http://localhost:3000/')
    // })


   it('Search Product', () => {
        cy.visit('/')
        // cy.contains('h1', 'todos')
        cy.get('[data-cy="email"]').type('fantahunfekadu1@gmail.com')
        cy.get('[data-cy="password"]').type('admin')
        //cy.get('.loginbutton').click()
        //    cy.get('.loginbutton', {timeout:10000}).click()

        // cy.wait(5000);
        //cy.contains('Login').click()
        cy.get('[data-cy=submit]').click()
        cy.get('[data-cy="searchprod"]').type('be')
        cy.get('[data-cy="onclicksearch"]').click()
        cy.get('.list-group').should('have.length', 1)
        cy.get('.list-group-item').should('have.length', 4)
        // cy.get('.list-group > :nth-child(1)').click()
        //cy.wait(5000)
        // cy.get('.list-group > :nth-child(1)').click()
        // cy.get('.editbutton').click()
        // cy.get('#currentprice').type(5000)
        // cy.get('.updatebutton').click()
        // cy.wait(3000)
        // cy.get('.returnback').click()

})


it('test search api with simple spy and stub response (fixture)', ()=>{

  cy.visit('/')
  cy.intercept('GET', '/api/product/search?text=be', {fixture: 'search.json'}).as('productsearch')

  // cy.get('#email').type('fantahunfekadu1@gmail.com')
  // cy.get('#password').type('admin')
  cy.get('[data-cy="email"]').type('fantahunfekadu1@gmail.com')
  cy.get('[data-cy="password"]').type('admin')
  cy.get('[data-cy=submit]').click()
  cy.get('[data-cy="searchprod"]').type('be')
  cy.get('[data-cy="onclicksearch"]').click()
  cy.wait('@productsearch').then((intercepteddata) =>{
    console.log(intercepteddata)
    // console.log(intercepteddata.response.body)
    const response = intercepteddata.response.body;
    const resprod = intercepteddata.response.body[0].name;
    expect(response).to.have.length(2);
    expect(resprod).to.equal('Bekolo');
     })
  // cy.get('@productsearch').should('not.be.null')


})



it('test search  with simple spy and stub response (empty string)', ()=>{

  cy.visit('/')
  cy.intercept('GET', '/api/product/search?text=be', {fixture: 'search.json'}).as('productsearch')

  // cy.get('#email').type('fantahunfekadu1@gmail.com')
  // cy.get('#password').type('admin')
  cy.get('[data-cy="email"]').type('fantahunfekadu1@gmail.com')
  cy.get('[data-cy="password"]').type('admin')
  cy.get('[data-cy=submit]').click()
  cy.get('[data-cy="searchprod"]').type('be')
  cy.get('[data-cy="onclicksearch"]').click()
  cy.wait('@productsearch').then((intercepteddata) =>{
    console.log(intercepteddata)
    // console.log(intercepteddata.response.body)
    const response = intercepteddata.response.body;
    const resprod = intercepteddata.response.body[0].name;
    expect(response).to.have.length(2);
    expect(resprod).to.equal('Bekolo');
     })
  cy.get('@productsearch').should('not.be.null')


})


it.skip('mocking with intercept test with dynamic fixture', ()=>{
  cy.visit('/')
  cy.intercept('GET', '/api/products', {fixture: 'product.json'}).as('products')

  cy.get('[data-cy="email"]').type('fantahunfekadu1@gmail.com')
  cy.get('[data-cy="password"]').type('admin')
  cy.get('[data-cy=submit]').click()
  cy.wait('@products').then((inter) =>{
    console.log(inter)

    console.log(inter.response.body.products)
    const response = inter.response.body.products;

    expect(response).to.have.length(9);
    
  })
  cy.get('@products').should('not.be.null')
  // cy.wait('@products').its('response.statusCode').should('eq', 200)

  // cy.get('li.selected').should('have.length', 3)

 
})



//Update Product Price

 

   
   
      
   
      

})
