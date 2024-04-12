<template>
      <HeaderBar  />
            <!-- Modal -->
            <div v-if="isModalOpen" class="modal" @click.self="closeModal">
                <div class="modal-content">
                    <span class="close" @click="closeModal">&times;</span>
                    <p class="font-extrabold">Usuário</p>
                    <form class="mt-8 flex items-center justify-center" action="" @submit.prevent="userUpdate">
                        <div class="flex  w-2/4 flex-col items-center justify-center">
                            <div class="flex  w-4/5 justify-between flex-wrap">
                                <div>
                            <label for="name" class="block text-sm font-medium leading-6 text-gray-900"
                            >Name</label
                            >
                            <div class="mt-2">
                            <input
                                id="name"
                                name="name"
                                type="text"
                                v-model="name"
                                class="block w-40 rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                            />
                            </div>
                        </div>
                        <div>
                            <label for="email" class="block text-sm font-medium leading-6 text-gray-900"
                            >Email</label
                            >
                            <div class="mt-2">
                            <input
                            
                                id="email"
                                email="email"
                                type="text"
                                v-model="email"
                                class="block w-40 rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                            />
                            </div>
                        </div>
                            </div>
                        <div class="flex  w-4/5 justify-between flex-wrap">
                            <div class="mt-4">
                            <label for="role" class="block text-sm font-medium leading-6 text-gray-900"
                            >Role</label
                            >
                            <div class="mt-2">
                                <select class="block w-40 rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6" name="pets" id="pet-select">
                                <option value="">--Please choose an option--</option>
                                <option value="0">0</option>
                                <option value="1">1</option>
                                <option value="2">2</option>
                                </select>
                            </div>
                        </div>
                        <div class="mt-4">
                            <label for="diretoria" class="block text-sm font-medium leading-6 text-gray-900"
                            >Diretoria</label
                            >
                            <div class="mt-2">
                                <select class="block w-40 rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6" name="pets" id="pet-select">
                                <option  value="">--Please choose an option--</option>
                                <option  v-for="item in raw_data.directorates" value="dog">{{item.Directorate}}</option>
                                </select>
                            </div>
                        </div>
                        <div class="mt-6 w-full">
                            <button type="submit"
                  class="flex w-full justify-center rounded-md bg-indigo-600 px-3 py-1.5 text-sm font-semibold leading-6 text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
                >
                  Salvar
                </button>
                        </div>
                        </div>

                        </div>
                    </form>
                </div>
            </div>
      <div class="w-screen h-screen flex items-center justify-center">
        <div class="painel rounded-md bg-white">

            <div class="flex flex-col w-full my-10 items-center">

                    <h1 class="font-extrabold">ADM</h1>
                    <CFormSelect
                aria-label="Default select example"
                :options="[
                    'Selecione a ação',
                    { label: 'Visão de edição', value: 'dited view' },
                    { label: 'Visão de criação', value: '2' },
                 
                ]">
                </CFormSelect>

            </div>
            <div class="flex w-full items-center justify-center">
                <table class="w-4/5">
                    <thead>
                        <th>ID</th>
                        <th>Name</th>
                        <th>Role</th>
                        <th>Diretoria</th>
                        <th>Email</th>
                    </thead>
                    <tr v-for="item in raw_data.users_raw" @click="modalConfig(item.Email)">
                        <td>{{item.ID}}</td>
                        <td>{{item.Name}}</td>
                        <td>{{item.Role}}</td>
                        <td>{{item.Directorate}}</td>
                        <td>{{item.Email}}</td>
                    </tr>
                </table>
            </div>
        </div>
       
      </div>
  </template>
  
<script setup>

import { CFormSelect } from '@coreui/vue';
import HeaderBar from '../components/HeaderBar.vue';
import { ref ,onMounted } from 'vue';
import axios from 'axios';


const raw_data = ref({'users_raw':["oi"]});
const get_raw_data = async () =>{
    try {
        const url = '/api/platform/get_all_sources';
        
        const response = await axios.get( url );

        console.log(response.data)
        if  (response.data != null){
            return response.data}

    } catch (error) {
        console.error('Ocorreu um erro:', error);
        
    }
    return [{"Name": "MainChart","Index": 1,"Value": "w-96"}]
} 


async function format_data() {
   var raw = await get_raw_data()
   raw_data.value = raw
   console.log("aaaaa ",raw_data.value.users_raw)
}



var email = 0
const userUpdate = async () =>{
      try {
        const url = '/api/auth/changeUser';
        const response = await axios.post( url,      {
		"email"       : "maria@gmail.com",
		"role"        : 2,
		"directorate" : 2,
		"name"       :  "Maria Carla"
});

        if (!response.data) {
          throw new Error('Erro ao fazer requisição para o servidor.');
        }

      } catch (error) {
        this.$toast.error(`Usuário ou senha estão incorretos.`);
        console.error('Ocorreu um erro:', error);
        
      
    }
    location.reload();


}

const isModalOpen = ref(false);


const modalConfig = (email) =>{
    email = email
    isModalOpen.value = true
}

const closeModal = () => {
  isModalOpen.value = false;
};

onMounted(() => {
format_data();
});

</script>
<style >


td, th {
  text-align: left;
  padding: 12px;
}

th {
    border-bottom: solid ;
}

td{
    border-bottom: 1px solid rgb(179, 184, 181);
}

tr:hover{
    cursor: pointer;
    background-color: rgb(240, 255, 247);
}
/* Modal Styles */
.modal {
  display: flex;
  position: fixed;
  z-index: 999;
  left: 0;
  top: 0;
  width: 100%;
  height: 100%;
  align-items: center;
  justify-content: center;
  background-color: rgba(0, 0, 0, 0.4);
}

/* Modal Content */
.modal-content {
  background-color: #fefefe;
  width: 60%;
  height: 40%;
  padding: 20px;
  border-radius: 5px;
}

/* Close Button */
.close {
  color: #aaa;
  float: right;
  font-size: 28px;
  font-weight: bold;
  cursor: pointer;
}
</style>