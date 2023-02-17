<br></br>
# Welcome to BookingApp

 **BookingApp** allows you to manage room bookings at your office, bussiness, or any places that demands room management.
 


## Room Management
Booking App offers you a broad flexibility on room management, you can add and delete rooms that is not necessarily active.

### Create and Delete Room

To create a new room you simply can login into the website and access Room Management Menu. Same as deleting room feature, you can access through that menu. Just in-case there are human-error factors, only Admin roles are allowed to create and delete rooms.

**Page Link**
{base-url}/room


**API Contracts**
Create Room
***POST*** localhost:8080/room 
*HTTP Response required* 

Read All Rooms
***GET*** localhost:8080/room

Delete Room
***DELETE*** localhost:8080/room/{room_id}




## Booking Room Management
Manage your bookings with your desired time.

### Create and Delete Booking

Access Create Booking from Room Booking Menu that shows room availability based on sessions. Each room are default to have ten sessions per-day, bookings can be reserved within the range of 6 days ahead. You can also summarize and delete your bookings through MyBookings menu. Super Admin can delete all booking to reset the sessions availability.

**Page Link**
{base-url}/bookings
{base-url}/mybookings

**API Contracts**
Create Room
***POST*** localhost:8080/booking 
*HTTP Response required* 

Read Bookings by Username
***GET*** localhost:8080/booking/{username}
Read Bookings by selected Day and Room
***GET*** localhost:8080/booking?booked_day={day}&booked_room={room_name}

Delete Booking by User ID
***DELETE*** localhost:8080/booking/{book_id}
Delete All Booking
***DELETE*** localhost:8080/bookingreset

## User Management
Register and manage users to grant and restrict menu accesses.

### Create and Delete Booking

New user can register to access booking menu, user accesses can be manage by Super Admin to grant/restrict accesses

**Page Link**
{base-url}/login
{base-url}/register
{base-url}/edit

**API Contracts**
Create User
***POST*** localhost:8080/user 
*HTTP Response required* 

Read All User
***GET*** localhost:8080/user

Grant Admin Role
***PUT*** localhost:8080/user/grant/{user_id}
Demote to User
***PUT*** localhost:8080/user/demote/{user_id}


Delete Booking by User ID
***DELETE*** localhost:8080/booking/{book_id}
Delete All Booking
***DELETE*** localhost:8080/bookingreset