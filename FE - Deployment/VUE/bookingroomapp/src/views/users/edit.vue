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
          <a class="nav-link" href="/room">Room Management</a>
        </li>
        <li v-if="role == 2" class="nav-item">
          <a class="nav-link active" href="/edit">User Management</a>
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
    <h1 class="display-4">Manage Existing User</h1>
    <p class="lead">Promote to Admin to Grant Room Management Access.</p>
  </div>
    </div>
    <div class="d-flex flex-wrap justify-content-around">
      <div v-for="user in users" :key="user.id">
        <div v-if="user.role == 0 || user.role == 1" class="d-flex justify-content-around my-5">
          <div class="card" style="width: 16rem;">
            <img class="card-img-top" src="https://i.pravatar.cc/1000" alt="Card image cap">
            <div class="card-body">
              <h5 class="card-title" style="text-transform: capitalize;">{{ user.name }}</h5>
              <h6 class="card-subtitle mb-2 text-muted"></h6>
              <h6 v-if="user.role == 0" class="card-subtitle mb-2 text-muted">Role : User</h6>
              <h6 v-if="user.role == 1" class="card-subtitle mb-2 text-muted">Role : Admin</h6>
              <h6 v-if="user.role == 2" class="card-subtitle mb-2 text-muted">Role : Super Admin</h6>
              <p class="card-text">Username : {{ user.username }}</p>
              <button @click="GrantAdmin(user.id)" type="button" class="btn btn-primary btn-sm mx-1 my-1">Grant Admin</button> 
              <button @click="DemoteUser(user.id)" type="button" class="btn btn-danger btn-sm my-1">Demote to User</button>
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
    name: 'MyUsersList',
  created(){
      this.MyUsersList()
  },
  data(){
      return{
          users:[],
          errorMsg:'',
          role:'',
          username : localStorage.getItem("username")
      }
  },
  methods:{
      MyUsersList(){
        this.role = localStorage.getItem("role")
          if(!this.role){
            this.$router.push({name: 'login.login'})
          }
          else if (this.role != 2){
            this.$router.push({name: 'login.login'})
          }
          axios
          .get('http://127.0.0.1:8080/users')
          .then((response)=>{
              console.log(response.data.data)
              this.users = response.data.data
          })
          .catch((error)=>{
              console.log(error)
              this.errorMsg = 'Error retrieving users'
          })
          localStorage.getItem("")
      },GrantAdmin(id){
        this.uname = localStorage.getItem("username")  
        console.log(this.uname)
        const sup = this.uname
        axios
        .put(`http://127.0.0.1:8080/user/grant/${id}/${sup}`)
        .then((response)=>{
              console.log(response.data.data)
              if (response.data.status == "SUCCESS"){
          window.location.reload()
                    }
          })
        .catch((error)=>{
              console.log(error)
              this.errorMsg = 'Error grant user'
          })
            this.$router.push({name: 'users.edit'})
      },DemoteUser(id){
        this.uname = localStorage.getItem("username")
        console.log(this.uname)
        const sup = this.uname
        axios
        .put(`http://127.0.0.1:8080/user/demote/${id}/${sup}`)
        .then((response)=>{
              console.log(response.data.data)
              if (response.data.status == "SUCCESS"){
          window.location.reload()
                    }
          })
        .catch((error)=>{
              console.log(error)
              this.errorMsg = 'Error demote admin'
          })
            this.$router.push({name: 'users.edit'})
      }
  },
  }
  </script>
  
  <style>
  
  </style>