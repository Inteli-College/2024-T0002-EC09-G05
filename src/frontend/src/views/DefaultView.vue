<script setup>

import HeaderBar from '../components/HeaderBar.vue';
import DashboardContainer from '../components/DashboardContainer.vue';

//http://10.254.17.50:3000/app

function apiData(){
  return false
}

</script>

<template>
  <div class="default">
    <!-- Renderiza os componentes somente se apiData for definido -->
    <HeaderBar v-if="apiData" />
    <DashboardContainer v-if="apiData" />
  </div>
</template>
