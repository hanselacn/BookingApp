
function Room_Alpha() {
    var text = "";
    var possible = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz";

    for (var i = 0; i < 10; i++)
      text += possible.charAt(Math.floor(Math.random() * possible.length));

    return text;
  }
  
  describe('Test task C3', ()=>{
    beforeEach(()=>{
        cy.viewport('macbook-16');
        cy.visit('http://localhost:5173/login');
    });
    it('Booking Page', () =>{
        //test after opening login page
        cy.contains('BookingApp').should('exist');
        cy.contains('Login').should('exist');
        cy.contains('Username').should('exist');
        cy.contains('Password').should('exist');
        cy.contains('Register').should('exist');
        cy.get("#exampleInputUsername1").type("superadmin")
        cy.get("#exampleInputPassword1",{timeout: 2000}).type("123")
        cy.get("#buttonlogin",{timeout: 2000}).click();
        cy.location('pathname',{timeout: 5000}).should('include', '/bookings');
        cy.contains('BookingApp').should('exist');
        cy.contains('Room Booking').should('exist');
        cy.contains('My Booking').should('exist');
        cy.contains('Room Management').should('exist');
        cy.contains('User Management').should('exist');
        cy.contains('Logout').should('exist');
        cy.contains("Check Room Availability").should('exist');
        cy.contains("Monday").should('exist');
        cy.contains("Tuesday").should('exist');
        cy.contains("Wednesday").should('exist');
        cy.contains("Thursday").should('exist');
        cy.contains("Friday").should('exist');
        cy.contains("Saturday").should('exist');
        cy.contains("Sunday").should('exist');
        cy.contains('Check Availability',{timeout: 2000}).should('exist').click({timeout: 2000});
        cy.location('pathname',{timeout: 5000}).should('include', '/bookings/roomday');
        cy.contains("Select Available Session").should('exist');
        cy.contains("If you can't find your desired session, try to check another room.").should('exist');
        cy.contains("08.00-09.00").should('exist');
        cy.contains("09.00-10.00").should('exist');
        cy.contains("10.00-11.00").should('exist');
        cy.contains("11.00-12.00").should('exist');
        cy.contains("12.00-13.00").should('exist');
        cy.contains("13.00-14.00").should('exist');
        cy.contains("14.00-15.00").should('exist');
        cy.contains("15.00-16.00").should('exist');
        cy.contains("16.00-17.00").should('exist');
        cy.contains("17.00-18.00").should('exist');
        cy.contains('My Bookings',{timeout: 2000}).should('exist').click({timeout: 2000});
        cy.location('pathname',{timeout: 5000}).should('include', '/mybooking');
        cy.contains("Manage your Bookings").should('exist');
        cy.contains("Change of plans? Delete your booking and re-order to adjust your needs").should('exist');
        cy.contains('Delete Booking').should('exist');
        cy.contains('Order Booking',{timeout: 2000}).should('exist').click({timeout: 2000});
        cy.location('pathname',{timeout: 5000}).should('include', '/bookings');
        cy.contains("Check Room Availability").should('exist');
        cy.contains('Room Management',{timeout: 2000}).should('exist').click({timeout: 2000});
        cy.location('pathname',{timeout: 5000}).should('include', '/room');
        cy.contains('Manage Available Rooms').should('exist');
        cy.contains('Delete Room').should('exist');
        cy.get("newroom").type(Room_Alpha())
        cy.contains('Submit',{timeout: 2000}).should('exist').click({timeout: 2000});
        cy.location('pathname',{timeout: 5000}).should('include', '/room');
        cy.contains(Room_Alpha()).should('exist')
    });
});
