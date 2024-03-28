<template>
  <div>
    <canvas id="polarAreaChart"></canvas>
  </div>
</template>

<script setup>
import { Chart, registerables } from 'chart.js';
import { onMounted } from 'vue';

Chart.register(...registerables);

onMounted(() => {
  const ctx = document.getElementById('polarAreaChart').getContext('2d');
  const polarAreaChart = new Chart(ctx, {
    type: 'polarArea',
    data: {
      labels: ['Butantã', 'Pinheiros', 'Liberdade', 'Morumbi'],
      datasets: [{
        data: [7, 12, 6, 8],
        backgroundColor: ['#FF6384', '#4BC0C0', '#FFCE56', '#E7E9ED'],
      }],
    },
    options: {
    }
  });
});
</script>
