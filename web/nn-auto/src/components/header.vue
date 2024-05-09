<template>
  <header class="site-header">
    <div class="site-identity">
      <h1>
        <RouterLink to="/">nn-auto</RouterLink>
      </h1>
      <h2 class="tigranFTP">by tigranFTP</h2>
    </div>
    <div class="login-sign-up">
      <RouterLink to="/login" v-if="userName===''">
        <div>Login</div>
      </RouterLink>
      <RouterLink to="/sign_up" v-if="userName===''">
        <div>Sign Up</div>
      </RouterLink>



      <RouterLink to="/create_model" v-if="userName!==''">
          Create Model
      </RouterLink>
      <h1 v-if="userName!==''" class="username">
        {{userName}}
      </h1>
      <button v-if="userName!==''" class="sign_out" @click.prevent="log_out()">
        Sign Out
      </button>
    </div>
  </header>
</template>

<script lang="ts" setup>
import axios from 'axios';
import {eraseCookie, getCookie} from "@/utils";
import { jwtDecode } from "jwt-decode";
import {ref} from "vue";
import {useRouter} from "vue-router";

let token = getCookie("token")
let userName = ref("")
if (token !== null) {
  let payload = jwtDecode(token)
  userName.value = payload.username
}

axios.get('/ping')
    .then(response => {
      // Обработка успешного ответа
      console.log(response.data);
    })
    .catch(error => {
      // Обработка ошибок
      console.error('Ошибка при запросе:', error);
    });


const  router =  useRouter()
function log_out() {
  eraseCookie('token')
  userName.value = ""
  router.push({name: "Main"})
}

</script>

<style>
body {
  font-family: Helvetica;
  margin: 0;
}

a {
  text-decoration: none;
  color: #000;
}


.sign_out {
  border: none;
  background: none;
  color:red;
  font-size: 1em;
  display: flex;
  align-items: center;
  box-shadow: inset 0 0 0 0 black;
  margin: 0em 0em 0em 1em;
  transition: color .3s ease-in-out, box-shadow .3s ease-in-out;
}

.sign_out:hover {
  box-shadow: inset 5em 0 0 0 black;
  color: white;
}

.username {
  font-size: 1em;
  display: flex;
  align-items: center;
  box-shadow: inset 0 0 0 0 black;
  color: black;
  margin: 0em 0em 0em 1em;
  transition: color .3s ease-in-out, box-shadow .3s ease-in-out;
}

.username:hover {
  box-shadow: inset 5em 0 0 0 black;
  color: white;
}

.login-sign-up {
  font-size: 3em;
  display: flex;
}

.login-sign-up a {
  display: flex;
  align-items: center;
  box-shadow: inset 0 0 0 0 black;
  color: black;
  margin: 0em 0em 0em 1em;
  transition: color .3s ease-in-out, box-shadow .3s ease-in-out;
}

.login-sign-up a:hover {
  box-shadow: inset 5em 0 0 0 black;
  color: white;
}


.site-header {
  border-bottom: 1px solid #ccc;
  padding: .0em 0.4em;
  display: flex;
  justify-content: space-between;
}

.site-identity h1 {
  font-size: 5em;
  margin: 0 0;
  display: inline-block;
}

.tigranFTP {
  font-size: 1.5em;
  margin: .4em 0;
  display: inline-block;
}

</style>
