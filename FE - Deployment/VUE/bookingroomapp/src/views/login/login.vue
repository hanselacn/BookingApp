<template>
    <nav class="navbar navbar-expand-lg navbar-light bg-light">
    <a class="navbar-brand mx-5" href="/login">BookingApp</a>
    <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
      <span class="navbar-toggler-icon"></span>
    </button>
    <div class="collapse navbar-collapse" id="navbarNav">
      <ul class="navbar-nav">
        <li class="nav-item">
          <a class="nav-link active" href="/login">Login</a>
        </li>
        <li class="nav-item">
          <a class="nav-link" href="/register">Register</a>
        </li>
        <!-- <li class="nav-item">
          <a class="nav-link disabled" href="#">Disabled</a>
        </li> -->
      </ul>
    </div>
  </nav>
  
  <div class="container my-5">
    <div class="d-flex flex-wrap justify-content-around">
        <div class="card" style="width: 20rem;">
            <div class="card-body">
              <h5 class="card-title">Login</h5>
<form>
  <div class="form-group">
    <h6 v-if="errorMsg">{{ errorMsg }}</h6>
    <label for="input username">Username</label>
    <input type="username" class="form-control" v-model="username" id="exampleInputUsername1" placeholder="Username">
  </div>
  <div class="form-group">
    <label for="exampleInputPassword1">Password</label>
    <input type="password" class="form-control" v-model="password" id="exampleInputPassword1" placeholder="Password">
  </div>
  <button type="button" @click="Login" id="buttonlogin" class="btn btn-primary my-3">Login</button>
</form>
            </div>
        </div>
</div>
  </div>
    </template>
    <script>
         import axios from 'axios'
    export default {
  name: 'Login',
  data(){
    return{
      username : null,
      password: null,
      errorMsg : "",
    }
  },
  methods:{
     Login(){
      const data= {
        "username" : this.username,
        "password" : this.password,
      }
      axios
      .put('http://127.0.0.1:8080/users/login', data)
      .then((response)=>{
                console.log(response.data)
                    this.errorMsg = "Login Success" 
                    var userdata = response.data.data
                    localStorage.setItem("username", this.username)        
                    localStorage.setItem("role", userdata.role)  
                    if (response.data.status == "SUCCESS"){
                      this.$router.push({name: 'bookings.index'})
                    }
            })
            .catch((error)=>{
                console.log(error)
                this.errorMsg="Invalid login credential!"
            })
      }
  },
}
    </script>
    
    <style>
    
    </style>