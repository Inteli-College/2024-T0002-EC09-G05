<template>
  <div>
    <canvas id="chart"></canvas>
  </div>
</template>

<script setup>
import { Chart, registerables } from 'chart.js';
import { onMounted, onUnmounted, ref } from 'vue';
import axios from 'axios';

Chart.register(...registerables);

const chartRef = ref(null);
let isFetching = false;
let fetchDataInterval = null;

async function fetchData() {
  if (isFetching) {
    return;
  }

  isFetching = true;
  const url = '/api/data/getDataByRelativeTime';
  const requestBody = {
    start_time: 5,
    aggregator: 30,  
    fields: ['temp'],  
    sensors: ['Pinheiros'],  
  };
  let labels;
  let values;

  try {
    const response = await axios.post(url, requestBody, {
      headers: {
        'Content-Type': 'application/json',
        'Token': 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MTI0NDc5NTUsInJvbGUiOjEsInVzZXJfaWQiOjF9.Uhlb9g6aMonA78FUZyw06Y4VftpKglY-VxiXb_ycxTw', 
      },
    });

    if (response.data && response.data.length > 0) {
      const sensorData = response.data[0].data;
      labels = sensorData.map(entry => {
        const correctedTime = entry.time.replace(" +0000 UTC", ""); 
        console.log("Corrected Time: ", correctedTime)
        const date = new Date(correctedTime);
        return date.toLocaleTimeString();
      });
      values = sensorData.map(entry => entry.value);

      console.log('Data: ', sensorData)
      console.log('Labels:', labels)
      console.log('Values:', values)
      
      if (chartRef.value) {
        console.log("Data received")
      }
      } else {
        console.warn('Received data is empty or in an unexpected format:', response.data);
      }
    return { labels, values };
  } catch (error) {
    console.error('Error:', error);
  } finally {
    isFetching = false;
}}

async function updateChart(chart){
    const data = await fetchData()
    console.log("Flush: ", data)
    data.labels.forEach((label, index) => {
        chart.data.labels.push(label);
        chart.data.datasets.forEach((dataset) => {
            if(data.values.length > index){
                dataset.data.push(data.values[index]);
            }
        });
    });
    chart.update();
}

onMounted(() => {
  const ctx = document.getElementById('chart');
  let chartRef = new Chart(ctx, {
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
  console.log("Chart Ref: ", chartRef)
  updateChart(chartRef);
  fetchDataInterval = setInterval(fetchData, 1000);
});

onUnmounted(() => {
  clearInterval(fetchDataInterval);
});

</script>
