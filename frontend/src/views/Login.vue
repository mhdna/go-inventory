<script setup>
// import {reactive} from 'vue' 
import {Login,Print} from '../../wailsjs/go/main/App'
import { useRouter } from 'vue-router'
import {ref} from "vue"

// const data = reactive({
//   name: "",
//   password: "",
//   resultText: "Please Enter your Username and Password:",
// })

const password = ref('')
const name = ref('')

const router = useRouter()

async function handleLogin() {
const success = await Login(name.value, password.value);
    if (success) {
    router.push('/home')
    }
}
</script>

<template>
  <main>
    <!-- <div id="result" class="result">{{ data.resultText }}</div> -->
    <div id="input" class="input-box">
      <input id="name" v-model="name" autocomplete="off" class="input" type="text" />
    </div>
    <div id="input" class="input-box">
        <input id="password" v-model="password" autocomplete="off" class="input" type="password" />
    </div>
    <button class="btn" @click="handleLogin">Login</button>
  </main>
</template>

<style scoped>
.result {
  height: 20px;
  line-height: 20px;
  margin: 1.5rem auto;
}

.input-box .btn {
  width: 60px;
  height: 30px;
  line-height: 30px;
  border-radius: 3px;
  border: none;
  margin: 0 0 0 20px;
  padding: 0 8px;
  cursor: pointer;
}

.input-box .btn:hover {
  background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
  color: #333333;
}

.input-box .input {
  border: none;
  border-radius: 3px;
  outline: none;
  height: 30px;
  line-height: 30px;
  padding: 0 10px;
  background-color: rgba(240, 240, 240, 1);
  -webkit-font-smoothing: antialiased;
}

.input-box .input:hover {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}

.input-box .input:focus {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}
</style>
