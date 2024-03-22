<template>
  <div
    class="bg-[url('https://blog.123milhas.com/wp-content/uploads/2022/04/conheca-estado-sao-paulo-conexao123-3.jpg')] w-screen h-screen"
  >
    <div class="filtro flex items-center justify-center w-screen h-screen">
      <div class="bg-white max-w-xl min-w-96 rounded-2xl">
        <div class="flex min-h-full flex-1 flex-col justify-center px-6 py-12 lg:px-8">
          <div class="sm:mx-auto sm:w-full sm:max-w-sm">
            <img class="mx-auto" src="../assets/logo.png" alt="Logo" height="128" width="128" />
            <h2 class="mt-10 text-center text-2xl font-bold leading-9 tracking-tight text-gray-900">
              Sign in to your account
            </h2>
          </div>

          <div class="mt-10 sm:mx-auto sm:w-full sm:max-w-sm">
            <form class="space-y-6" @submit.prevent="login">
              <div>
                <label for="email" class="block text-sm font-medium leading-6 text-gray-900"
                  >Email address</label
                >
                <div class="mt-2">
                  <input
                    id="email"
                    name="email"
                    type="email"
                    autocomplete="email"
                    v-model="email"
                    class="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                  />
                </div>
              </div>

              <div>
                <div class="flex items-center justify-between">
                  <label for="password" class="block text-sm font-medium leading-6 text-gray-900"
                    >Password</label
                  >
                  <div class="text-sm">
                    <a href="#" class="font-semibold text-indigo-600 hover:text-indigo-500"
                      >Forgot password?</a
                    >
                  </div>
                </div>
                <div class="mt-2">
                  <input
                    id="password"
                    name="password"
                    type="password"
                    v-model="password"
                    autocomplete="current-password"
                    class="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                  />
                </div>
              </div>

              <div>
                <button
                  class="flex w-full justify-center rounded-md bg-indigo-600 px-3 py-1.5 text-sm font-semibold leading-6 text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
                >
                  Sign in
                </button>
              </div>
            </form>

            <p class="mt-10 text-center text-sm text-gray-500">
              Not have an account?
              {{ ' ' }}
              <a href="#" class="font-semibold leading-6 text-indigo-600 hover:text-indigo-500"
                >Sign up</a
              >
            </p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>

import axios from 'axios';
import { useCookies } from 'vue3-cookies'

export default {
  data() {
    return {
      email: '',
      password: ''
    }
  },
  methods: {
    async login() {
      try {
        const url = '/api/auth/login';
        

        const response = await axios.post( url,  {
          email: String(this.email),
          password: this.password
        });

        
        useCookies().cookies.set('authToken', response.data.token)
        if (!response.data) {
          throw new Error('Erro ao fazer requisição para o servidor.');
        }

        this.$toast.success(`Login successful! Welcome, ${response.data.name}`);
        window.user = response.data.id;
        this.$router.push('/');
      } catch (error) {
        this.$toast.error(`Usuário ou senha estão incorretos.`);
        console.error('Ocorreu um erro:', error);
        // Aqui você pode exibir uma mensagem de erro para o usuário, informando que houve um problema durante a autenticação
      }
    }
  }
}
</script>

