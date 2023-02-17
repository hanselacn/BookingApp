
describe('Test task C3', ()=>{
    beforeEach(()=>{
        cy.viewport('macbook-16');
        cy.visit('http://localhost:5173/login'); 
        Cypress.config('chromeWebSecurity',false);
    });
    it('Test Login page', () =>{
        //test after opening login page
        cy.contains('BookingApp').should('exist');
        cy.contains('Login').should('exist');
        cy.contains('Username').should('exist');
        cy.contains('Password').should('exist');
        cy.contains('Register').should('exist');
    });
    it('Test Register page', ()=>{
                //register page
        cy.contains('Register').click();
        cy.contains('BookingApp').should('exist');
        cy.contains('Login').should('exist');
        cy.contains('Username').should('exist');
        cy.contains('Password').should('exist');
        cy.contains('Register').should('exist');
        cy.contains('Name').should('exist');
        cy.url().should('include', '/register');
    });
});
