<template>
  <div>
    <canvas id="chart"></canvas>
  </div>
</template>

<script setup>
import { Chart, registerables } from 'chart.js';
import { onMounted, ref } from 'vue';

Chart.register(...registerables);

const chartRef = ref(null);

async function fetchData() {
  const url = 'http://localhost:8080/data/getDataByRelativeTime';
  const requestBody = {
    start_time: 60, // Puxa os dados da última hora
    aggregator: 30, // Média dos dados a cada 30 segundos
    fields: ['temp'], // Dados de temperatura e umidade
    sensors: ['Pinheiros'],
  };

  try {
    const response = await fetch(url, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        //'Token': "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MTI0MDQ1NDAsInJvbGUiOjEsInVzZXJfaWQiOjF9.jndAxkhb-6OZloqVyi6ZyA3pnHRjFJ'
      },
      body: JSON.stringify(requestBody),
    });

    if (!response.ok) throw new Error('Network response was not ok');

    const data = await response.json();

    const labels = data.map(entry => new Date(entry.time).toLocaleTimeString());
    const values = data.map(entry => entry.value);

    if (chartRef.value) {
      chartRef.value.data.labels = labels;
      chartRef.value.data.datasets[0].data = values;
      chartRef.value.update();
    }
  } catch (error) {
    console.error('There was a problem with the fetch operation:', error);
  }
}

onMounted(() => {
  const ctx = document.getElementById('chart').getContext('2d');
  chartRef.value = new Chart(ctx, {
    type: 'line',
    data: {
      labels: [],
      datasets: [{
        label: 'Dados do Sensor',
        backgroundColor: 'rgba(255, 99, 132, 0.2)',
        borderColor: 'rgba(255, 99, 132, 1)',
        data: [],
      }],
    },
    options: {
    },
  });

  fetchData();
});
</script>
