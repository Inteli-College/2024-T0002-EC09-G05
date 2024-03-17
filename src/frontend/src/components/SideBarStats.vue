<script setup>
import { CChart } from '@coreui/vue-chartjs'
</script>

<template>
<CChart
  type="polarArea"
  :options="{
    // backgroundColor: rgba(0, 1, 0, 0.1)
  }"
  :data="{
    labels: ['Butantã', 'Pinheiros', 'Liberdade', 'Morumbi'],
    datasets: [
      {
        data: [7, 12, 6, 8],
        backgroundColor: ['#FF6384', '#4BC0C0', '#FFCE56', '#E7E9ED'],
      },
    ],
  }"
/>
</template>