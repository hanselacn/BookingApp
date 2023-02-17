<template>
   <nav class="navbar navbar-expand-lg navbar-light bg-light">
    <a class="navbar-brand mx-5" href="/bookings">BookingApp</a>
    <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
      <span class="navbar-toggler-icon"></span>
    </button>
    <div class="collapse navbar-collapse" id="navbarNav">
      <ul class="navbar-nav">
        <li class="nav-item">
          <a class="nav-link active" href="/bookings">Room Booking</a>
        </li>
        <li class="nav-item">
          <a class="nav-link" href="/mybooking">My Bookings</a>
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
    <h1 class="display-4">Check Room Availability</h1>
    <p class="lead">Select one day and check sessions availability based on selected room</p>
  </div>
    </div>
    <div class="d-flex justify-content-around mx-5">
        <div class="btn-group btn-group-toggle" data-toggle="buttons">
            <label class="btn btn-dark">
                <input ref= "monday" type="radio" name="options" id="option1" autocomplete="off" checked> Monday
            </label>
            <label class="btn btn-dark">
                <input ref= "tuesday" type="radio" name="options" id="option2" autocomplete="off"> Tuesday
            </label>
            <label class="btn btn-dark">
                <input ref= "wednesday" type="radio" name="options" id="option3" autocomplete="off"> Wednesday
            </label>
            <label class="btn btn-dark">
                <input ref= "thursday" type="radio" name="options" id="option4" autocomplete="off"> Thursday
            </label>
            <label class="btn btn-dark">
                <input ref= "friday" type="radio" name="options" id="option5" autocomplete="off"> Friday
            </label>
            <label class="btn btn-outline-dark">
                <input ref= "saturday" type="radio" name="options" id="option6" autocomplete="off"> Saturday
            </label>
            <label class="btn btn-outline-dark">
                <input ref= "sunday" type="radio" name="options" id="option7" autocomplete="off"> Sunday
            </label>
        </div>
    </div>
    <div class="d-flex flex-wrap justify-content-around">
      <h3 v-if="errorMsg">{{ errorMsg }}</h3>
            <div v-for="room in rooms" :key="room.id">
        <div class="d-flex justify-content-around my-5">
            
            <div class="card mx-2" style="width: 15rem; text-align: center;">
              
            <img class="card-img-top" src="https://picsum.photos/1000/200" alt="Card image cap">
                    <div class="card-body">
                        <h5 class="card-title" style="text-transform: capitalize;">{{room.room_name}}</h5>
                        <button @click="getRoomInfo(room.room_name)" class="btn btn-primary">Check Availability</button>
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
    name: 'RoomList',
    created(){
        this.getRooms()
    },
    data(){
        return{
            rooms:[],
            errorMsg:'',
            role:'',
          username : localStorage.getItem("username")
        }
    },
    methods:{
        getRooms(){
          this.role = localStorage.getItem("role")
          if(!this.role){
            this.$router.push({name: 'login.login'})
          }
            axios
            .get('http://127.0.0.1:8080/room')
            .then((response)=>{
                console.log(response.data.data)
                this.rooms = response.data.data
            })
            .catch((error)=>{
                console.log(error)
                this.errorMsg = 'Error retrieving rooms'
            })
        },
        getRoomInfo(room_name){
            if (this.$refs.monday.checked){
              localStorage.setItem('room_day',"monday")  
            }
            else if (this.$refs.tuesday.checked){
              localStorage.setItem('room_day',"tuesday")  
            }
            else if (this.$refs.wednesday.checked){
              localStorage.setItem('room_day',"wednesday")  
            }
            else if (this.$refs.thursday.checked){
              localStorage.setItem('room_day',"thursday")  
            }
            else if (this.$refs.friday.checked){
              localStorage.setItem('room_day',"friday")  
            }
            else if (this.$refs.saturday.checked){
              localStorage.setItem('room_day',"saturday")  
            }
            else{
              localStorage.setItem('room_day',"sunday")  
            }
            localStorage.setItem('room_name', room_name)
            this.$router.push({name: 'bookings.bookbyroomday'})
        }   
    },
  }
  </script>
  
  <style>
  
  </style>