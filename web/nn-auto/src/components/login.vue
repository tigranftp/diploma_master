<template>
  <div class="main">
    <h3>Enter your login credentials</h3>
    <form action="">
      <label>
        Username:
      </label>
      <input v-model="login"
             type="text"
             id="first"
             name="first"
             placeholder="Enter your Username" required>

      <label>
        Password:
      </label>
      <input v-model="password"
             type="password"
             id="password"
             name="password"
             placeholder="Enter your Password" required>

      <div class="wrap">
        <button type="submit"
                @click.prevent="submit_login">
          Submit
        </button>
      </div>
    </form>
    <p>Not registered?
      <RouterLink to="/sign_up" style="text-decoration: none;">Create an account
      </RouterLink>
    </p>
  </div>
</template>

<script lang="ts" setup>
import axios from "axios";
import {ref} from "vue"
import {getCookie, setCookie} from "@/utils"
import {useRouter} from "vue-router";
import token from "./header.vue"

const login = ref("")
const password = ref("")
const router = useRouter()

if (getCookie("token") != null) {
  router.push({
    name: "Main"
  })
}

function submit_login() {

  axios.get('/logging', {headers: {Password: password.value, Login: login.value}})
      .then(response => {
        // Обработка успешного ответа
        //const json = JSON.parse(response.data)
        setCookie("token", response.data.jwt_token, 5)
        setCookie("refresh_token", response.data.refresh_token, 5*24*60)
        router.go(0)
      })
      .catch(error => {
        // Обработка ошибок
        alert(error)
        console.error('Ошибка при запросе:', error);
      });
}


</script>

<style scoped>

.main {
  font-size: 1.5em;
  margin-top: 1em;
  margin-left: auto;
  margin-right: auto;
  background-color: #fff;
  border-radius: 1em;
  box-shadow: 0 0 1em rgba(0, 0, 0, 0.2);
  padding: 0.5em 1em;
  transition: transform 0.2s;
  width: 15em;
  text-align: center;
}

h1 {
  color: #4CAF50;
}

label {
  display: block;
  margin-top: 1em;
  margin-bottom: 0;
  text-align: left;
  color: #555;
  font-weight: bold;
}


input {
  display: block;
  width: 100%;
  margin-bottom: 15px;
  padding: 10px;
  box-sizing: border-box;
  border: 1px solid #ddd;
  font-size: 1em;
  border-radius: 5px;
}

button {
  padding: 15px;
  border-radius: 10px;
  margin-top: 15px;
  margin-bottom: 15px;
  border: none;
  color: white;
  cursor: pointer;
  background-color: #4CAF50;
  width: 100%;
  font-size: 1em;
}

.wrap {
  display: flex;
  justify-content: center;
  align-items: center;
}
</style>