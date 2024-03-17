<template>
  <CChart type="line" :data="data" :options="options" ref="mainChartRef" />
</template>

<script>
import { onMounted, ref } from 'vue'
import { CChart } from '@coreui/vue-chartjs'

  const random = (min, max) => Math.floor(Math.random() * (max - min + 1) + min)

export default {
  name: 'MainChart',
  components: {
    CChart,
  },
  setup() {
    const mainChartRef = ref()
    const data = {
      labels: ['10:25:00', '10:26:00', '10:27:00', '10:28:00', '10:29:00', '10:30:00', '10:31:00'],
      datasets: [
        {
          label: 'My First dataset',
          backgroundColor: 'rgba(0, 123, 255, .1)', // Example RGB value for --cui-info
          borderColor: 'rgb(0, 123, 255)', // Example RGB value for --cui-info
          pointHoverBackgroundColor: 'rgb(0, 123, 255)', // Example RGB value for --cui-info
          borderWidth: 2,
          data: [
            random(50, 200),
            random(50, 200),
            random(50, 200),
            random(50, 200),
            random(50, 200),
            random(50, 200),
            random(50, 200),
          ],
          fill: true,
        },
        {
          label: 'My Second dataset',
          backgroundColor: 'transparent',
          borderColor: 'rgb(40, 167, 69)', // Example RGB value for --cui-success
          pointHoverBackgroundColor: 'rgb(40, 167, 69)', // Example RGB value for --cui-success
          borderWidth: 2,
          data: [
            random(50, 200),
            random(50, 200),
            random(50, 200),
            random(50, 200),
            random(50, 200),
            random(50, 200),
            random(50, 200),
          ],
        },
        {
          label: 'My Third dataset',
          backgroundColor: 'transparent',
          borderColor: 'rgb(220, 53, 69)', // Example RGB value for --cui-danger
          pointHoverBackgroundColor: 'rgb(220, 53, 69)', // Example RGB value for --cui-danger
          borderWidth: 1,
          borderDash: [8, 5],
          data: [65, 65, 65, 65, 65, 65, 65],
        },
      ],
    }

    const options = {
      maintainAspectRatio: false,
      plugins: {
        legend: {
          display: false,
        },
      },
      scales: {
        x: {
          grid: {
            color: 'rgb(233, 236, 239)', // Example RGB value for --cui-border-color-translucent
            drawOnChartArea: false,
          },
          ticks: {
            color: 'rgb(52, 58, 64)', // Example RGB value for --cui-body-color
          },
        },
        y: {
          beginAtZero: true,
          border: {
            color: 'rgb(233, 236, 239)', // Example RGB value for --cui-border-color-translucent
          },
          grid: {
            color: 'rgb(233, 236, 239)', // Example RGB value for --cui-border-color-translucent
          },
          max: 250,
          ticks: {
            color: 'rgb(52, 58, 64)', // Example RGB value for --cui-body-color
            maxTicksLimit: 5,
            stepSize: Math.ceil(250 / 5),
          },
        },
      },
      elements: {
        line: {
          tension: 0.4,
        },
        point: {
          radius: 0,
          hitRadius: 10,
          hoverRadius: 4,
          hoverBorderWidth: 3,
        },
      },
    }

    onMounted(() => {
    })

    return {
      data,
      mainChartRef,
      options,
    }
  },
}
</script>
