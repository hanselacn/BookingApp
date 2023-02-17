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
    <div class="jumbotron jumbotron-fluid">
  <div class="container">
    <h1 class="display-4">Select Available Session</h1>
    <p class="lead">If you can't find your desired session, try to check another room.</p>
  </div>
  <div  v-if="errorMsg" class="alert alert-danger alert-dismissible fade show" style="text-align: center;" role="alert">
  <strong>{{ errorMsg }}</strong> Unfortunately, you have to pick another session that still available.
</div>
    </div>
    <div class="d-flex flex-wrap justify-content-around">
        <div class="d-flex justify-content-around my-5 mx-1">
          <div class="card" style="width: 12rem;">
            <div class="card-body">
              <h5 class="card-title">08.00-09.00</h5>
              <h6 class="card-subtitle mb-2 text-muted" style="text-transform: capitalize;">{{roomday}}</h6>
                <p class="card-text" style="text-transform: capitalize;">{{roomname}}</p>
                <div v-for="broomday in broomdays" :key="broomday.id">
                  <div v-if="broomday.sesa == true">
                    <button type="submit" @click="submitBooking()" class="btn btn-dark btn-sm disabled" id="button1">Already Booked
                    </button>
                  </div>
                </div>
                <button type="submit" @click="submitBookinga()" class="btn btn-primary btn-sm my-1" id="button1">Book Session
                  </button>                 
                </div>
            </div>
          </div>
        <div class="d-flex justify-content-around my-5 mx-1">
          <div class="card" style="width: 12rem;">
            <div class="card-body">
              <h5 class="card-title">09.00-10.00</h5>
              <h6 class="card-subtitle mb-2 text-muted" style="text-transform: capitalize;">{{roomday}}</h6>
                <p class="card-text" style="text-transform: capitalize;">{{roomname}}</p>
                <div v-for="broomday in broomdays" :key="broomday.id">
                  <div v-if="broomday.sesb == true">
                    <button v-bind:show=true type="button" @click="submitBooking()" class="btn btn-dark btn-sm disabled" id="button1">Already Booked
                    </button>
                  </div>
                </div>
                <button type="button" @click="submitBookingb()" class="btn btn-primary btn-sm my-1" id="button1">Book Session
                  </button>   
            </div>
          </div>
        </div>
        <div class="d-flex justify-content-around my-5 mx-1">
          <div class="card" style="width: 12rem;">
            <div class="card-body">
              <h5 class="card-title">10.00-11.00</h5>
              <h6 class="card-subtitle mb-2 text-muted" style="text-transform: capitalize;">{{roomday}}</h6>
                <p class="card-text" style="text-transform: capitalize;">{{roomname}}</p><div v-for="broomday in broomdays" :key="broomday.id">
                  <div v-if="broomday.sesc == true">
                    <button v-bind:show=true type="button" @click="submitBooking()" class="btn btn-dark btn-sm disabled" id="button1">Already Booked
                    </button>
                  </div>
                </div>
                <button type="button" @click="submitBookingc()" class="btn btn-primary btn-sm my-1" id="button1">Book Session
                  </button>  
            </div>
          </div>
        </div>
        <div class="d-flex justify-content-around my-5 mx-1">
          <div class="card" style="width: 12rem;">
            <div class="card-body">
              <h5 class="card-title">11.00-12.00</h5>
              <h6 class="card-subtitle mb-2 text-muted" style="text-transform: capitalize;">{{roomday}}</h6>
                <p class="card-text" style="text-transform: capitalize;">{{roomname}}</p><div v-for="broomday in broomdays" :key="broomday.id">
                  <div v-if="broomday.sesd == true">
                    <button v-bind:show=true type="button" @click="submitBooking()" class="btn btn-dark btn-sm disabled" id="button1">Already Booked
                    </button>
                  </div>
                </div>
                <button type="button" @click="submitBookingd()" class="btn btn-primary btn-sm my-1" id="button1">Book Session
                  </button>  
            </div>
          </div>
        </div>
        <div class="d-flex justify-content-around my-5 mx-1">
          <div class="card" style="width: 12rem;">
            <div class="card-body">
              <h5 class="card-title">12.00-13.00</h5>
              <h6 class="card-subtitle mb-2 text-muted" style="text-transform: capitalize;">{{roomday}}</h6>
                <p class="card-text" style="text-transform: capitalize;">{{roomname}}</p>
                <div v-for="broomday in broomdays" :key="broomday.id">
                  <div v-if="broomday.sese == true">
                    <button v-bind:show=true type="button" @click="submitBooking()" class="btn btn-dark btn-sm disabled" id="button1">Already Booked
                    </button>
                  </div>
                </div>
                <button type="button" @click="submitBookinge()" class="btn btn-primary btn-sm my-1" id="button1">Book Session
                  </button>  
            </div>
          </div>
        </div>
        <div class="d-flex justify-content-around my-5 mx-1">
          <div class="card" style="width: 12rem;">
            <div class="card-body">
              <h5 class="card-title">13.00-14.00</h5>
              <h6 class="card-subtitle mb-2 text-muted" style="text-transform: capitalize;">{{roomday}}</h6>
                <p class="card-text" style="text-transform: capitalize;">{{roomname}}</p>
              <div v-for="broomday in broomdays" :key="broomday.id">
                  <div v-if="broomday.sesf == true">
                    <button v-bind:show=true type="button" @click="submitBooking()" class="btn btn-dark btn-sm disabled" id="button1">Already Booked
                    </button>
                  </div>
                </div>
                <button type="button" @click="submitBookingf()" class="btn btn-primary btn-sm my-1" id="button1">Book Session
                  </button>  
            </div>
          </div>
        </div>
        <div class="d-flex justify-content-around my-5 mx-1">
          <div class="card" style="width: 12rem;">
            <div class="card-body">
              <h5 class="card-title">14.00-15.00</h5>
              <h6 class="card-subtitle mb-2 text-muted" style="text-transform: capitalize;">{{roomday}}</h6>
                <p class="card-text" style="text-transform: capitalize;">{{roomname}}</p>
              <div v-for="broomday in broomdays" :key="broomday.id">
                  <div v-if="broomday.sesg == true">
                    <button v-bind:show=true type="button" @click="submitBooking()" class="btn btn-dark btn-sm disabled" id="button1">Already Booked
                    </button>
                  </div>
                </div>
                <button type="button" @click="submitBookingg()" class="btn btn-primary btn-sm my-1" id="button1">Book Session
                  </button>  
            </div>
          </div>
        </div>
        <div class="d-flex justify-content-around my-5 mx-1">
          <div class="card" style="width: 12rem;">
            <div class="card-body">
              <h5 class="card-title">15.00-16.00</h5>
              <h6 class="card-subtitle mb-2 text-muted" style="text-transform: capitalize;">{{roomday}}</h6>
                <p class="card-text" style="text-transform: capitalize;">{{roomname}}</p>
              <div v-for="broomday in broomdays" :key="broomday.id">
                  <div v-if="broomday.sesh == true">
                    <button v-bind:show=true type="button" @click="submitBooking()" class="btn btn-dark btn-sm disabled" id="button1">Already Booked
                    </button>
                  </div>
                </div>
                <button type="button" @click="submitBookingh()" class="btn btn-primary btn-sm my-1" id="button1">Book Session
                  </button>  
            </div>
          </div>
        </div>
        <div class="d-flex justify-content-around my-5 mx-1">
          <div class="card" style="width: 12rem;">
            <div class="card-body">
              <h5 class="card-title">16.00-17.00</h5>
              <h6 class="card-subtitle mb-2 text-muted" style="text-transform: capitalize;">{{roomday}}</h6>
                <p class="card-text" style="text-transform: capitalize;">{{roomname}}</p>
              <div v-for="broomday in broomdays" :key="broomday.id">
                  <div v-if="broomday.sesi == true">
                    <button v-bind:show=true type="button" @click="submitBooking()" class="btn btn-dark btn-sm disabled" id="button1">Already Booked
                    </button>
                  </div>
                </div>
                <button type="button" @click="submitBookingi()" class="btn btn-primary btn-sm my-1" id="button1">Book Session
                  </button>  
            </div>
          </div>
        </div>
        <div class="d-flex justify-content-around my-5 mx-1">
          <div class="card" style="width: 12rem;">
            <div class="card-body">
              <h5 class="card-title">17.00-18.00</h5>
              <h6 class="card-subtitle mb-2 text-muted" style="text-transform: capitalize;">{{roomday}}</h6>
                <p class="card-text" style="text-transform: capitalize;">{{roomname}}</p>
              <div v-for="broomday in broomdays" :key="broomday.id">
                  <div v-if="broomday.sesj == true">
                    <button v-bind:show=true type="button" @click="submitBooking()" class="btn btn-dark btn-sm disabled" id="button1">Already Booked
                    </button>
                  </div>
                </div>
                <button type="button" @click="submitBookingj()" class="btn btn-primary btn-sm my-1" id="button1">Book Session
                  </button>  
            </div>
          </div>
        </div>
    </div>
    </div>
  
    </template>
    <script>
import axios from 'axios'

    export default {
    name: 'BookByRoomDayList',
    created(){
        this.getBookByRoomDay()
    },
    data(){
        return{
            broomdays:[],
            sessions:[],
            errorMsg:'',
            roomday: '',
            roomname : '',
            successMsg:'',
            role:"",
            username : localStorage.getItem("username")

        }
    },  
    methods:{
        getBookByRoomDay(){
          this.role = localStorage.getItem("role")
          if(!this.role){
            this.$router.push({name: 'login.login'})
          }
            var dataday = localStorage.getItem("room_day")
            var dataroom = localStorage.getItem("room_name")

            this.roomday = dataday
            this.roomname = dataroom

            axios
            .get(`http://127.0.0.1:8080/booking?booked_day=${dataday}&booked_room=${dataroom}`)
            .then((response)=>{
                console.log(response.data.data)
                var final = response.data.data
                    console.log(final);
                    this.broomdays = final             
            })
            .catch((error)=>{
                console.log(error)
                this.errorMsg = 'Error retrieving roomday'
            })

        },
        submitBookinga(){
            var dataday = localStorage.getItem("room_day")
            var dataroom = localStorage.getItem("room_name")
            var databy = localStorage.getItem("username")
            const data = {
                "booked_by" : databy,
                "booked_room" : dataroom,
                "booked_day" : dataday,
                "sesa" : true,
                "sesb" : false,
                "sesc" : false,
                "sesd" : false,
                "sese" : false,
                "sesf" : false,
                "sesg" : false,
                "sesh" : false,
                "sesi" : false,
                "sesj" : false,
              }
            console.log()
            axios
            .post("http://127.0.0.1:8080/booking",
              data
            ).then((response)=>{
                console.log(response.data.data)
                if (response.data.status == "SUCCESS"){
                    window.location.reload()

                    }
              })
                .catch((error)=>{
                console.log(error)       
                
                this.errorMsg = 'The session is already booked!'
            })
            
        },
        submitBookingb(){
            var dataday = localStorage.getItem("room_day")
            var dataroom = localStorage.getItem("room_name")
            var databy = localStorage.getItem("username")
            const data = {
                "booked_by" : databy,
                "booked_room" : dataroom,
                "booked_day" : dataday,
                "sesa" : false,
                "sesb" : true,
                "sesc" : false,
                "sesd" : false,
                "sese" : false,
                "sesf" : false,
                "sesg" : false,
                "sesh" : false,
                "sesi" : false,
                "sesj" : false,
              }
            axios
            .post("http://127.0.0.1:8080/booking",
              data
            ).then((response)=>{
                console.log(response.data.data)
                if (response.data.status == "SUCCESS"){
                    window.location.reload()
                    
                    }
              })
                .catch((error)=>{
                console.log(error)       
                
                this.errorMsg = 'The session is already booked!'
            })
            
        },
        submitBookingc(){
            var dataday = localStorage.getItem("room_day")
            var dataroom = localStorage.getItem("room_name")
            var databy = localStorage.getItem("username")
            const data = {
                "booked_by" : databy,
                "booked_room" : dataroom,
                "booked_day" : dataday,
                "sesa" : false,
                "sesb" : false,
                "sesc" : true,
                "sesd" : false,
                "sese" : false,
                "sesf" : false,
                "sesg" : false,
                "sesh" : false,
                "sesi" : false,
                "sesj" : false,
              }
            console.log()
            axios
            .post("http://127.0.0.1:8080/booking",
              data
            ).then((response)=>{
                console.log(response.data.data)
                if (response.data.status == "SUCCESS"){
                    window.location.reload()
                    
                    }
              })
                .catch((error)=>{
                console.log(error)       
                
                this.errorMsg = 'The session is already booked!'
            })
            
        },
        submitBookingd(){
            var dataday = localStorage.getItem("room_day")
            var dataroom = localStorage.getItem("room_name")
            var databy = localStorage.getItem("username")
            const data = {
                "booked_by" : databy,
                "booked_room" : dataroom,
                "booked_day" : dataday,
                "sesa" : false,
                "sesb" : false,
                "sesc" : false,
                "sesd" : true,
                "sese" : false,
                "sesf" : false,
                "sesg" : false,
                "sesh" : false,
                "sesi" : false,
                "sesj" : false,
              }
            console.log()
            axios
            .post("http://127.0.0.1:8080/booking",
              data
            ).then((response)=>{
                console.log(response.data.data)
                if (response.data.status == "SUCCESS"){
                    window.location.reload()
                    
                    }
              })
                .catch((error)=>{
                console.log(error)       
                
                this.errorMsg = 'The session is already booked!'
            })
            
        },
        submitBookinge(){
            var dataday = localStorage.getItem("room_day")
            var dataroom = localStorage.getItem("room_name")
            var databy = localStorage.getItem("username")
            const data = {
                "booked_by" : databy,
                "booked_room" : dataroom,
                "booked_day" : dataday,
                "sesa" : false,
                "sesb" : false,
                "sesc" : false,
                "sesd" : false,
                "sese" : true,
                "sesf" : false,
                "sesg" : false,
                "sesh" : false,
                "sesi" : false,
                "sesj" : false,
              }
            console.log()
            axios
            .post("http://127.0.0.1:8080/booking",
              data
            ).then((response)=>{
                console.log(response.data.data)
                if (response.data.status == "SUCCESS"){
                    window.location.reload()
                    
                    }
              })
                .catch((error)=>{
                console.log(error)       
                
                this.errorMsg = 'The session is already booked!'
            })
            
        },
        submitBookingf(){
            var dataday = localStorage.getItem("room_day")
            var dataroom = localStorage.getItem("room_name")
            var databy = localStorage.getItem("username")
            const data = {
                "booked_by" : databy,
                "booked_room" : dataroom,
                "booked_day" : dataday,
                "sesa" : false,
                "sesb" : false,
                "sesc" : false,
                "sesd" : false,
                "sese" : false,
                "sesf" : true,
                "sesg" : false,
                "sesh" : false,
                "sesi" : false,
                "sesj" : false,
              }
            console.log()
            axios
            .post("http://127.0.0.1:8080/booking",
              data
            ).then((response)=>{
                console.log(response.data.data)
                if (response.data.status == "SUCCESS"){
                    window.location.reload()
                    
                    }
              })
                .catch((error)=>{
                console.log(error)       
                
                this.errorMsg = 'The session is already booked!'
            })
            
        },
        submitBookingg(){
            var dataday = localStorage.getItem("room_day")
            var dataroom = localStorage.getItem("room_name")
            var databy = localStorage.getItem("username")
            const data = {
                "booked_by" : databy,
                "booked_room" : dataroom,
                "booked_day" : dataday,
                "sesa" : false,
                "sesb" : false,
                "sesc" : false,
                "sesd" : false,
                "sese" : false,
                "sesf" : false,
                "sesg" : true,
                "sesh" : false,
                "sesi" : false,
                "sesj" : false,
              }
            console.log()
            axios
            .post("http://127.0.0.1:8080/booking",
              data
            ).then((response)=>{
                console.log(response.data.data)
                if (response.data.status == "SUCCESS"){
                    window.location.reload()
                    
                    }
              })
                .catch((error)=>{
                console.log(error)       
                
                this.errorMsg = 'The session is already booked!'
            })
            
        },
        submitBookingh(){
            var dataday = localStorage.getItem("room_day")
            var dataroom = localStorage.getItem("room_name")
            var databy = localStorage.getItem("username")
            const data = {
                "booked_by" : databy,
                "booked_room" : dataroom,
                "booked_day" : dataday,
                "sesa" : false,
                "sesb" : false,
                "sesc" : false,
                "sesd" : false,
                "sese" : false,
                "sesf" : false,
                "sesg" : false,
                "sesh" : true,
                "sesi" : false,
                "sesj" : false,
              }
            console.log()
            axios
            .post("http://127.0.0.1:8080/booking",
              data
            ).then((response)=>{
                console.log(response.data.data)})
                
            .catch((error)=>{
                console.log(error)
                this.errorMsg = 'The session is already booked!'
            })
            
        },
        submitBookingi(){
            var dataday = localStorage.getItem("room_day")
            var dataroom = localStorage.getItem("room_name")
            var databy = localStorage.getItem("username")
            const data = {
                "booked_by" : databy,
                "booked_room" : dataroom,
                "booked_day" : dataday,
                "sesa" : false,
                "sesb" : false,
                "sesc" : false,
                "sesd" : false,
                "sese" : false,
                "sesf" : false,
                "sesg" : false,
                "sesh" : false,
                "sesi" : true,
                "sesj" : false,
              }
            console.log()
            axios
            .post("http://127.0.0.1:8080/booking",
              data
            ).then((response)=>{
                console.log(response.data.data)
                if (response.data.status == "SUCCESS"){
                    window.location.reload()
                    
                    }
              })
                .catch((error)=>{
                console.log(error)       
                
                this.errorMsg = 'The session is already booked!'
            })
            
        },
        submitBookingj(){
            var dataday = localStorage.getItem("room_day")
            var dataroom = localStorage.getItem("room_name")
            var databy = localStorage.getItem("username")
            const data = {
                "booked_by" : databy,
                "booked_room" : dataroom,
                "booked_day" : dataday,
                "sesa" : false,
                "sesb" : false,
                "sesc" : false,
                "sesd" : false,
                "sese" : false,
                "sesf" : false,
                "sesg" : false,
                "sesh" : false,
                "sesi" : false,
                "sesj" : true,
              }
            console.log()
            axios
            .post("http://127.0.0.1:8080/booking",
              data
            ).then((response)=>{
                console.log(response.data.data)})
                
            .catch((error)=>{
                console.log(error)
                this.errorMsg = 'The session is already booked!'
                this.$router.push({name: 'bookings.bookbyroomday'})
            })
            
        }
    },
  }
  </script>
  
  <style>
  
  </style>