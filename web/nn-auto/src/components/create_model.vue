<template>
  <div class="main">
    <h5> Нужные файлы (модель и т.д.):</h5>
    <input
        type="file"
        @change="onFilesChanged($event)"
        multiple
    />
    <div v-for="file in files" :key="file.name" class="files"> {{file.name}}</div>
    <h5 > Скрипт (.py расширение):</h5>
    <input
        type="file"
        @change="onScriptChanged($event)"
        accept=".py"
    />
    <div v-if="script !== undefined && script !== null" class="files"> {{script.name}} </div>
  <button v-on:click="submitForm()">Upload</button>
  </div>
</template>

<script lang="ts" setup>
  import {  ref } from "vue";
  import axios from "axios";
  let files = ref<FileList | null>();
  let script = ref<File | null> ()

  function onFilesChanged($event: Event) {
    const target = $event.target as HTMLInputElement;
    if (!target.files) {
      files.value = null
      return
    }
    files.value = target.files
  }
  function onScriptChanged($event: Event) {
    const target = $event.target as HTMLInputElement;
    if (!target.files) {
      script.value = null
      return
    }
    script.value = target.files[0]
  }


  function submitForm() {
    const formData = new FormData();
    if (files.value == null) {
      console.log("no files applied")
      return
    }
    if (script.value == null) {
      console.log("no script applied")
      return;
    }

    for(let i=0; i < files.value.length; i++) {
      formData.append("files", files.value[i])
    }

    formData.append("script", script.value)

    axios.post('create_model',
        formData,
        {
          headers: {
            'Content-Type': 'multipart/form-data'
          }
        }
    ).then(function (data) {
      console.log(data.data);
    })
    .catch(error => {

      console.error('Ошибка при запросе:', error);
    });
  }

</script>

<style scoped>

h5 {
  margin-top: 1em;
  margin-bottom: 0;
}


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

.files {
  margin-top: 0.5em;
  font-size: 0.5em;
  text-align: left;
  color: black;
}
</style>