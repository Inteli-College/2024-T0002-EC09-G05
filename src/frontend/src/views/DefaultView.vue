<script setup>

import HeaderBar from '../components/HeaderBar.vue';

import HeaderBarComun from '../components/Comum/HeaderBar.vue'
import DashboardContainerComum from '../components/Comum/DashboardContainer.vue'

import { useCookies } from 'vue3-cookies'


import { defineAsyncComponent } from 'vue'

const DashboardContainer = defineAsyncComponent(() =>
  import('../components/DashboardContainer.vue')
)

const authTokenKey = 'authToken';


function apiData(){
  if (useCookies().cookies.get(authTokenKey) == null){
    return  false
  } 


  return true
}

</script>

<template >
  <div class="default" v-if="apiData() == true">
    <!-- Renderiza os componentes com base na apiData -->
    <HeaderBar  />
    <DashboardContainer />
  </div>
  <div class="commun"  v-else>
    <HeaderBarComun />
    <DashboardContainerComum />
  </div>
</template>

