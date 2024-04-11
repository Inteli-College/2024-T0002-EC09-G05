<template>
    <div class="w-full h-auto">
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
      start_time: 60,
      aggregator: 360,  
      fields: ['hum'],  
      sensors: ['Liberdade'],  
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
          const correctedTime = entry.time.replace(":00 +0000 UTC", ""); 
          console.log("Corrected Time: ", correctedTime)
          const date = new Date(correctedTime);
          return date.toLocaleTimeString([], {hour: '2-digit', minute:'2-digit'});
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
          label: 'Porcentagem de Chuva',
          backgroundColor: 'rgba(37, 174, 186, 0.3)',
          borderColor: 'rgba(37, 174, 186, 1)',
          pointBackgroundColor: 'rgba(11, 99, 107, 0.8)',
          data: [],
          fill: true,
          cubicInterpolationMode: 'monotone',
          tension: 0.4,
        }],
      },
      options: {
        plugins: {
            legend: {
                display: false,
            },
            title: {
                display: true,
                text: 'Porcentagem de Precipitação',
            },
        },
        scales: {
          y: {
            display: true,
            grid: {
                display: true,
            },
            ticks: {
                // Include a dollar sign in the ticks
                callback: function(value, index, ticks) {
                    return value + '%';
                },
            },
          },
          x: {
            display: true,
            grid: {
                display: false,
            },
          },
          
        },
      },
    });
    console.log("Chart Ref: ", chartRef)
    updateChart(chartRef);
    fetchDataInterval = setInterval(fetchData, 10000);
  });
  
  onUnmounted(() => {
    clearInterval(fetchDataInterval);
  });
  
  </script>
  