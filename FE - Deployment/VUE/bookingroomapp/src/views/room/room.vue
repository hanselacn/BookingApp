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
          <a class="nav-link" href="/mybooking">My Bookings</a>
        </li>
        <li v-if="role == 1 || role == 2" class="nav-item">
          <a class="nav-link active" href="/room">Room Management</a>
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
  <div class="jumbotron jumbotron-fluid">
  <div class="container">
    <h1 class="display-4">Manage Available Rooms</h1>
    <p class="lead">Add or Delete Room that can be accessed through BookingApp</p>
  </div>
    </div>
  <form>
  <div class="form-group">
    <label for="inputroom">Add New Room</label>
    <input id="newroom" type="room" class="form-control" v-model="rmname" aria-describedby="room" placeholder="Enter Room Name">
  </div>
  <button @click="Addroom()" type="button" class="btn btn-primary my-1">Submit</button>
</form>
<div class="d-flex flex-wrap justify-content-around">
      <div v-for="room in rooms" :key="room.id">
        <div class="d-flex justify-content-around my-5">
          <div class="card" style="width: 14rem;">
            <img class="card-img-top" src="https://picsum.photos/1000/700" alt="Card image cap">
            <div class="card-body">
              <h5 class="card-title" style="text-transform: capitalize;">{{ room.room_name }}</h5>
              <p class="card-text mb-2 text-muted">Created by : {{ room.created_by }}</p>
              <button @click="DeleteRoom(room.id)" type="button" class="btn btn-danger btn-sm my-1">Delete Room</button>
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
    name: 'MyRoomList',
  created(){
      this.MyRoomList()
  },
  data(){
      return{
          rmname :null,
          rooms:[],
          errorMsg:'',
          role:'',
          username : localStorage.getItem("username")
      }
  },
  methods:{
      MyRoomList(){
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
      DeleteRoom(id){
        const us = localStorage.getItem("username")
        axios
        .delete(`http://127.0.0.1:8080/deleteroom/${us}/${id}`)
        .then((response)=>{
              console.log(response.data.data)
              if (response.data.status == "SUCCESS"){
          window.location.reload()
                    }
          })
        .catch((error)=>{
              console.log(error)
              this.errorMsg = 'Error delete room'
          })
            this.$router.push({name: 'room.room'})
      },
      Addroom(){
        const us = localStorage.getItem("username")
        const data= {
        "room_name" : this.rmname,
        "created_by" : us,
      }
      axios
        .post(`http://127.0.0.1:8080/room`, data)
        .then((response)=>{
              console.log(response.data.data)
              if (response.data.status == "SUCCESS"){
          window.location.reload()
                    }
          })
        .catch((error)=>{
              console.log(error)
              this.errorMsg = 'Error delete room'
          })
            this.$router.push({name: 'room.room'})
      },
},
  }
  </script>
  
  <style>
  
  </style>