<template>
    <nav class="navbar navbar-expand-lg navbar-light bg-light">
    <a class="navbar-brand mx-5" href="/login">BookingApp</a>
    <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
      <span class="navbar-toggler-icon"></span>
    </button>
    <div class="collapse navbar-collapse" id="navbarNav">
      <ul class="navbar-nav">
        <li class="nav-item">
          <a class="nav-link" href="/login">Login</a>
        </li>
        <li class="nav-item">
          <a class="nav-link active" href="/register">Register</a>
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
                
              <h5 class="card-title">Register your account</h5>
<form>
  <div class="form-group">
    <label for="input username">Username</label>
    <input type="username" v-model="username" class="form-control" id="inputusername" aria-describedby="username" placeholder="Enter username">
  </div>
  <div class="form-group">
    <label for="input name">Name</label>
    <input type="name" v-model="name" class="form-control" id="name" placeholder="Name">
  </div>
  <div class="form-group">
    <label for="exampleInputPassword1">Password</label>
    <input type="password" v-model="password" class="form-control" id="exampleInputPassword1" placeholder="Password">
  </div>
  <button type="button" @click="Register()" class="btn btn-primary my-3">Register Now</button>
</form>
            </div>
        </div>
</div>
  </div>
  
    </template>
    <script>
     import axios from 'axios'
    export default {
  name: 'Register',
  data(){
    return{
      username : null,
      name: null,
      password: null,
      errorMsg : "",
    }
  },
  methods:{
     Register(){
      const data= {
        "username" : this.username,
        "password" : this.password,
        "name" : this.name,
      }
      console.log(this.username)
      axios
      .post('http://127.0.0.1:8080/user', data)
      .then((response)=>{
                console.log(response.data)
                    this.errorMsg = "Register Success"
            })
            .catch((error)=>{
                console.log(error)
                this.errorMsg = 'Error registering user'
                this.$router.push({name: 'homepage.register'})
            })
      this.$router.push({name: 'login.login'})
      }
  },
}
    </script>
    
    <style>
    
    </style>