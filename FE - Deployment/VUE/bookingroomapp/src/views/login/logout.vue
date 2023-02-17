<template>
    </template>
    <script>
import axios from 'axios'
export default {
  name: 'Logout',
  created(){
      this.Logout()
  },
  methods:{
     Logout(){const data= {
        "username" : localStorage.getItem("username"),
      }
      axios
      .put('http://127.0.0.1:8080/users/logout', data)
      .then((response)=>{
                console.log(response.data.data)
                    localStorage.removeItem("username")        
                    localStorage.removeItem("role")
            })
            .catch((error)=>{
                console.log(error)
            })
            this.$router.push({name: 'login.login'})
      }
  },
}
    </script>
    
    <style>
    
    </style>