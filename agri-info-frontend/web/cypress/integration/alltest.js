/// <reference types="cypress" />



describe('Agri Info Testing', () => {

  //  const infoemail= "fantahunfekadu1@gmail.com"
  // const superemail= "samuaeladadnew@outlook.com"
  // const password = "admin"

    // beforeEach(() => {
    //   cy.visit('http://localhost:3000/')
    // })


    it('Login loads with correct user credentials', () => {
      // cy.logininfo('infoemail', 'password')
      cy.logininfo()
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

      // cy.contains('h3', 'Sign into your account')
      cy.get('h3').should('contain', 'Sign into your account')
      cy.get('[data-cy="email"]').type('fantahunfekadu1@gmail.com')
      cy.get('[data-cy="password"]').type('adminn')
      cy.get('[data-cy="submit"]').click()
      // cy.get('.list-group')

      //  cy.logininfo('infoemail', 'password')
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



//Search


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
cy.intercept('GET', '/api/product/search?text=', {fixture: 'emptysearch.json'}).as('productsearch')

// cy.get('#email').type('fantahunfekadu1@gmail.com')
// cy.get('#password').type('admin')
cy.get('[data-cy="email"]').type('fantahunfekadu1@gmail.com')
cy.get('[data-cy="password"]').type('admin')
cy.get('[data-cy=submit]').click()
// cy.get('[data-cy="searchprod"]').type('be')
cy.get('[data-cy="onclicksearch"]').click()
cy.wait('@productsearch').then((intercepteddata) =>{
console.log(intercepteddata)
// console.log(intercepteddata.response.body)
const response = intercepteddata.response.body;
const resprod = intercepteddata.response.body[0].name;
expect(response).to.have.length(11);
expect(resprod).to.equal('Gebss');
 })
cy.get('@productsearch').should('not.be.null')


})

//Infoadmin load product
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
  // expect(response).to.have.length(11);
  // cy.log(JSON.stringify(inter))
  // console.log(JSON.stringify(inter.body))
  })
  
  
  })


it('mocking with intercept test with dynamic fixture for products load', ()=>{
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



//Update Product 

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
  // cy.wait(5000)
  // cy.get('.returnback').click()

})


//Super Admin Add Pro

it('Add Product', () => {
    cy.visit('/')
    // cy.contains('h1', 'todos')
    cy.get('[data-cy="email"]').type('samuaeladadnew@outlook.com')
    cy.get('[data-cy="password"]').type('admin')
    //cy.get('.loginbutton').click()
    //    cy.get('.loginbutton', {timeout:10000}).click()

    // cy.wait(5000);
    //cy.contains('Login').click()
    cy.get('[data-cy="submit"]').click()
    cy.get('[data-cy="superaddpro"]').click()
    cy.get('[data-cy="supernav"]').should('be.visible')
    cy.get('[data-cy="supernav"]').should('contain', 'Super Admin Page To Add New Product')
    cy.get(':nth-child(1) > label').should('contain', 'Name')
    cy.get(':nth-child(2) > label').should('contain', 'Unit ID')
    cy.get(':nth-child(3) > label').should('contain', 'Prod Area')
    cy.get(':nth-child(4) > label').should('contain', 'Price')
    cy.get('[data-cy="prodname"]').type('Sorghum')
    cy.get('[data-cy="produnit"]').select('Kg')
    cy.get('[data-cy="prodarea"]').type('Addis')
    cy.get('[data-cy="prodprice"]').type(2000)
    cy.get('[data-cy="submitprod"]').click()
    // cy.wait(20000)
    // cy.get('[data-cy="returnback"]').click()
    // cy.get('.list-group').should('have.length', 1)
    // cy.get('.list-group-item').should('have.length', 12)
    // cy.get('.list-group > :nth-child(1)').click()
    //cy.wait(5000)
    // cy.get('.list-group > :nth-child(1)').click()
    // cy.get('.editbutton').click()
    // cy.get('#currentprice').type(5000)
    // cy.get('.updatebutton').click()
    // cy.wait(3000)
    // cy.get('.returnback').click()

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




//View Product Info
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


  

//Manage Infodmins
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
    cy.get(':nth-child(1) > .row > .col-sm-7 > .Name').should('contain', 'Name :Fantahun')
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



//Add Infoadmins

it('Register Infoadmin', () => {
  cy.visit('/')
  // cy.contains('h1', 'todos')
  cy.get('[data-cy="email"]').type('samuaeladadnew@outlook.com')
  cy.get('[data-cy="password"]').type('admin')
  cy.get('[data-cy="submit"]').click()
  cy.get('[data-cy="adminlink"]').click()
  cy.get('[data-cy="regadminbtn"]').click()
  cy.get('[data-cy="superadminheader"]').should('contain', 'Super Admin Page To Register New Admin')
  // cy.get('[data-cy="superaddpro"]').click()
  cy.get('[data-cy="superadminheader"]').should('be.visible')
  // cy.get('[data-cy="supernav"]').should('contain', 'Super Admin Page To Add New Product')
  // cy.get(':nth-child(1) > label').should('contain', 'Name')
  // cy.get(':nth-child(2) > label').should('contain', 'Unit ID')
  // cy.get(':nth-child(3) > label').should('contain', 'Prod Area')
  // cy.get(':nth-child(4) > label').should('contain', 'Price')
  cy.get('[data-cy="firstname"]').type('Natnael')
  // cy.get('[data-cy="produnit"]').select('Kg')
  cy.get('[data-cy="lastname"]').type('Bahiruu')
  cy.get('[data-cy="email"]').type('natialemnali@gmail.com')
  cy.get('[data-cy="reginfo"]').click()
  // cy.wait(20000)
  // cy.get('[data-cy="returnback"]').click()
  // cy.get('.list-group').should('have.length', 1)
  // cy.get('.list-group-item').should('have.length', 4)
  // cy.get('.list-group > :nth-child(1)').click()
  //cy.wait(5000)
  // cy.get('.list-group > :nth-child(1)').click()
  // cy.get('.editbutton').click()
  // cy.get('#currentprice').type(5000)
  // cy.get('.updatebutton').click()
  // cy.wait(3000)
  // cy.get('.returnback').click()

})




//Manages Message

it('Super admin messages', ()=>{
    cy.visit('/')
    // cy.intercept('GET', '/api/products', {fixture: 'product.json'}).as('products')
  
    cy.get('[data-cy="email"]').type('samuaeladadnew@outlook.com')
    cy.get('[data-cy="password"]').type('admin')
    cy.get('[data-cy="submit"]').click()
    cy.get('[data-cy="messagelink"]').click()
    cy.get('h4').should('have.text', 'Inboxes')
    cy.get('h4').should('be.visible')
    cy.get(':nth-child(1) > .nav-link').should('have.text', 'Received')
    cy.get(':nth-child(2) > .nav-link').should('have.text', 'Sent')
    cy.get('[data-cy="sendmessagebtn"]').should('contain', 'Send').and('be.visible')
    cy.get('[data-cy="showadminlabel"]').should('contain', 'Show to Admins :').and('be.visible')
    cy.get('[data-cy="showachooseprod"]').should('have.text', 'Choose a Product :').and('be.visible')
    cy.get('[data-cy="showchooselang"]').should('have.text', 'Choose a Language :').and('be.visible')

    // cy.get('.description').should('have.text', 'Please click on a Admin to view full profile...')
    // cy.get(':nth-child(1) > .row > .col-sm-7 > .Name').should('contain', 'Name :Samuael')
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



  //Super admin login and add dict word 

  it.only('Login and add dict word', () => {
    cy.visit('/')
    // cy.contains('h1', 'todos')
    cy.get('[data-cy="email"]').type('samuaeladadnew@outlook.com')
    cy.get('[data-cy="password"]').type('admin')
  cy.get('[data-cy="submit"]').click()
  //cy.wait(5000)
  cy.get('[data-cy="dictlink"]').click()
  cy.get('[data-cy="selectlang"]').select('Amharic').should('have.value', 'amh')
  cy.get('[data-cy="inputkey"]').type('Englishword')
  cy.get('[data-cy="inputvalue"]').type('Amharicword')
  cy.get('[data-cy="addictbtn"]').click()
  // cy.get(':nth-child(4) > p').should('contain', 'created')
 //And tHIS 

//   cy.get('.list-group > :nth-child(2)').click()
//   cy.get('[data-cy="editpricebutton"]').click()
//   cy.get('[data-cy="currentproprice"]').type(5000)
//   cy.get('[data-cy="updatepricebutton"]').click()
//   cy.wait(5000)
//   cy.get('.returnback').click()
})




  it('spy and stub response for super admin searching recent dictionaries using fixture', ()=>{

    cy.visit('/')
    cy.intercept('GET', 'api/dictionaries?offset=0&limit=9', {fixture: 'recentdict.json'}).as('recentdictsearch')
    
  
    cy.get('[data-cy="email"]').type('samuaeladadnew@outlook.com')
    cy.get('[data-cy="password"]').type('admin')
    cy.get('[data-cy="submit"]').click()
     cy.get('[data-cy="dictlink"]').click()
     cy.wait('@recentdictsearch').then((intercepteddata) =>{
    console.log(intercepteddata)
    // console.log(intercepteddata.response.body)
    const response = intercepteddata.response.body.dictionaries;
    const resprod = intercepteddata.response.body.dictionaries[0].translation;
    const statuscode = intercepteddata.response.body.status_code;
    expect(response).to.have.length(3);
    expect(resprod).to.equal('ጤፍጤፍ');
    expect(statuscode).to.equal(200);
     })
    cy.get('@recentdictsearch').should('not.be.null')
    
    
    })

    it('Login and edit dict word', () => {
  
      cy.loginsuper('superemail', 'password')
      cy.get('[data-cy="dictlink"]').click()
      cy.get('.list-group > :nth-child(2)').click()
      cy.get('.linkeditiword').click()
      cy.get('[data-cy="naveditdict"]').should('contain', 'Super Admin currentDictionary Edit Page')
      cy.get('h4').should('have.text', 'Word')
      cy.get('[data-cy="langlabel"]').should('be.visible')
      cy.get('[data-cy="translabel"]').should('contain', 'Translation')
      cy.get('[data-cy="wordtranslation"]').type('newtranslation')
      cy.get('[data-cy="updatetranslation"]').click()
      
      })



      it.skip('Spy and Stub translation edit', () => {
        cy.visit('/')
        cy.intercept('GET', 'api/dictionaries?offset=0&limit=9', {fixture: 'recentdict.json'}).as('recentdictsearch')
        
  
        cy.loginsuper('superemail', 'password')
        cy.get('[data-cy="dictlink"]').click()
        cy.get('.list-group > :nth-child(2)').click()
        cy.get('.linkeditiword').click()
        cy.get('[data-cy="naveditdict"]').should('contain', 'Super Admin currentDictionary Edit Page')
        cy.get('h4').should('have.text', 'Word')
        cy.get('[data-cy="langlabel"]').should('be.visible')
        cy.get('[data-cy="translabel"]').should('contain', 'Translation')
        cy.get('[data-cy="wordtranslation"]').type('newtranslation')
        cy.get('[data-cy="updatetranslation"]').click()
        
        })


  

})


