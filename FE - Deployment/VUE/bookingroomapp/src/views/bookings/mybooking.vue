<template>
   <nav class="navbar navbar-expand-lg navbar-light bg-light">
    <a class="navbar-brand mx-5" href="/bookings">BookingApp</a>
    <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
      <span class="navbar-toggler-icon"></span>
    </button>
    <div class="collapse navbar-collapse" id="navbarNav">
      <ul class="navbar-nav">
        <li class="nav-item">
          <a class="nav-link" href="/bookings">Room Booking</a>
        </li>
        <li class="nav-item">
          <a class="nav-link active" href="/mybooking">My Bookings</a>
        </li>
        <li v-if="role == 1 || role == 2" class="nav-item">
          <a class="nav-link" href="/room">Room Management</a>
        </li>
        <li v-if="role == 2" class="nav-item">
          <a class="nav-link" href="/edit">User Management</a>
        </li>
        <li class="nav-item">
          <a class="nav-link" href="/logout">Logout</a>
        </li>
        <!-- <li class="nav-item">
          <a class="nav-link disabled" href="#">Disabled</a>
        </li> -->
      </ul>
    </div>
    <a class="nav-link mx-sm-5 disabled" href="#">Hello, {{username}}</a>
  </nav>

<div class="container my-5">
  <div class="jumbotron jumbotron-fluid" style="text-align: center;">
  <div class="container">
    <h1 class="display-4">Manage your Bookings</h1>
    <p class="lead">Change of plans? Delete your booking and re-order to adjust your needs</p>
    <a class="btn btn-primary" style="text-align: center;" data-toggle="collapse" href="/bookings" role="button" aria-expanded="false" aria-controls="collapseExample">
    Order Booking
  </a>
  </div>
    </div>
       <div class="d-flex flex-wrap justify-content-around">
        <div v-for ="book in bookings" :key="book.id">
        <div class="d-flex justify-content-around my-5">
          <div class="card" style="width: 14rem;">
            <img class="card-img-top" src="http://placekitten.com/1000/700" alt="Card image cap">
            <div class="card-body">
              <h5 class="card-title" style="text-transform: capitalize;">{{book.booked_day}}</h5>
              <h6 class="card-subtitle mb-2 text-muted">{{ book.booked_room }}</h6>
              <p v-if ="book.sesa == true" class="card-text">08.00 - 09.00</p>
              <p v-if ="book.sesb == true" class="card-text">09.00 - 10.00</p>
              <p v-if ="book.sesc == true" class="card-text">10.00 - 11.00</p>
              <p v-if ="book.sesd == true" class="card-text">11.00 - 12.00</p>
              <p v-if ="book.sese == true" class="card-text">12.00 - 13.00</p>
              <p v-if ="book.sesf == true" class="card-text">13.00 - 14.00</p>
              <p v-if ="book.sesg == true" class="card-text">14.00 - 15.00</p>
              <p v-if ="book.sesh == true" class="card-text">15.00 - 16.00</p>
              <p v-if ="book.sesi == true" class="card-text">16.00 - 17.00</p>
              <p v-if ="book.sesj == true" class="card-text">17.00 - 18.00</p>
              <button type="button"  @click="DeleteBookingList(book.id)" class="btn btn-warning btn-sm">Delete Booking</button>
            </div>
          </div>
        </div>
      </div>
    </div>
</div>
  </template>
  <script>
  import axios from 'axios'

export default {
  name: 'MyBookingList',
  created(){
      this.MyBookingList()
  },
  data(){
      return{
          bookings:[],
          errorMsg:'',
          role:'',
          username : localStorage.getItem("username")
      }
  },
  methods:{
      MyBookingList(){
        var user = localStorage.getItem("username")
        this.role = localStorage.getItem("role")
          if(!this.role){
            this.$router.push({name: 'login.login'})
          }
          axios
          .get(`http://127.0.0.1:8080/booking/${user}`)
          .then((response)=>{
              console.log(response.data.data)
              this.bookings = response.data.data
          })
          localStorage.getItem("")
      },DeleteBookingList(id){
        localStorage.setItem('book_id', id)
        axios
        .delete(`http://127.0.0.1:8080/booking/${id}`)
        .then((response)=>{
        console.log(response.data)
        if (response.data.status == "SUCCESS"){
          window.location.reload()
                    }
            }).catch((error)=>{
              console.log(error)
              this.errorMsg = 'Error delete booking'
          })
             
      }
  },
}
  </script>
  
  <style>
  
  </style>