<script setup>

import HeaderBar from '../components/HeaderBar.vue';
import DashboardContainer from '../components/DashboardContainer.vue';

function apiData(){
  return true
}

</script>

<template>
  <div class="default">
    <!-- Renderiza os componentes somente se apiData for definido -->
    <HeaderBar v-if="apiData" />
    <DashboardContainer v-if="apiData" />
  </div>
</template>
