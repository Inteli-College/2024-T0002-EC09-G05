<template>
  <div class="w-full h-auto">
    <select v-model="selectedTime">
      <option value="15">Últimos 15 minutos</option>
      <option value="60">Última 1h</option>
      <option value="1440">Últimas 24h</option>
    </select>
    <canvas id="chart"></canvas>
  </div>
</template>

<script setup>
import { Chart, registerables } from 'chart.js';
import { onMounted, onUnmounted, ref, watch } from 'vue';
import axios from 'axios';

Chart.register(...registerables);

let isFetching = false;
let fetchDataInterval = null;
const selectedTime = ref(15);
let lastFetchTime = null; 

async function fetchData(selectedMinutes, updateOnly = false) {
  if (isFetching) {
    return;
  }

  isFetching = true;
  const now = new Date();
  let startTimeToUse = selectedMinutes;

  if (updateOnly && lastFetchTime) {
    const timeSinceLastFetch = Math.round((now - lastFetchTime) / 60000);
    startTimeToUse = Math.min(timeSinceLastFetch, selectedMinutes);
  }

  const url = '/api/data/getDataByRelativeTime';
  const requestBody = {
    start_time: Number(startTimeToUse),
    aggregator: 30,  
    fields: ['temp'],  
    sensors: ['Pinheiros'],  
  };

  let labels = [];
  let values = [];

  try {
    const response = await axios.post(url, requestBody);

    console.log('Response:', response.data);
    if (response.data && response.data.length > 0) {
      const sensorData = response.data[0].data;
      labels = sensorData.map(entry => {
        const correctedTime = entry.time.replace(" +0000 UTC", ""); 
        const date = new Date(correctedTime);
        return date.toLocaleTimeString();
      });
      values = sensorData.map(entry => entry.value);
    } else {
      console.warn('Received data is empty or in an unexpected format:', response.data);
    }
    lastFetchTime = now;
    return { labels, values };
  } catch (error) {
    console.error('Error:', error);
  } finally {
    isFetching = false;
  }
}

async function updateChart(chart, selectedMinutes, updateOnly = false) {
  const data = await fetchData(selectedMinutes, updateOnly);
  if (data && data.labels && data.values) {
    let maxPoints = 30;
    if (selectedMinutes === 60) {
      maxPoints = 60;
    } else if (selectedMinutes === 1440) {
      maxPoints = 100;
    }

    if (updateOnly) {
      data.labels.forEach((label, index) => {
        chart.data.labels.push(label);
        chart.data.datasets.forEach((dataset) => {
          dataset.data.push(data.values[index]);
        });
      });

      // Only shift elements if the array is full
      while (chart.data.labels.length > maxPoints) {
        chart.data.labels.shift(); // Remove the oldest label
        chart.data.datasets.forEach((dataset) => {
          dataset.data.shift(); // Remove the oldest data point
        });
      }
    } else {
      // For initial load or when changing time frames, replace all data
      chart.data.labels = data.labels.slice(-maxPoints);
      chart.data.datasets.forEach((dataset) => {
        dataset.data = data.values.slice(-maxPoints);
      });
    }

    chart.update();
  }
}

let chartInstance = null;

onMounted(() => {
  const ctx = document.getElementById('chart');
  chartInstance = new Chart(ctx, {
    type: 'line',
    data: {
      labels: [],
      datasets: [{
        label: 'Temperatura',
        backgroundColor: 'rgba(255, 99, 132, 0.2)',
        borderColor: 'rgba(255, 99, 132, 1)',
        data: [],
      }],
    },
    options: {
      scales: {
        y: {
          display: true,
        },
        x: {
          display: true,
        }
      },
    },
  });

  updateChart(chartInstance, selectedTime.value);

  fetchDataInterval = setInterval(() => {
    updateChart(chartInstance, selectedTime.value, true);
  }, 30000);
});

watch(selectedTime, (newVal) => {
  console.log('Selected time:', newVal);
  if (chartInstance) {
    updateChart(chartInstance, newVal);
    lastFetchTime = null;
  }
});

onUnmounted(() => {
  if (fetchDataInterval) clearInterval(fetchDataInterval);
});
</script>
