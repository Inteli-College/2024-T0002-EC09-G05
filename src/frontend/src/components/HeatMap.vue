<script setup>
import { GoogleMap, HeatmapLayer } from 'vue3-google-map'

const inteli = { lat: -23.555626840909746, lng: -46.73381383445283 }

const heatmapData = [
  //Ambas as formas adicionam um ponto ao heatmap. O 1° adiciona um ponto ponderado, ex -> 1 ponto equivalente a 2
  { location: { lat: -23.5556, lng: -46.7338 }, weight: 1 },
  { location: { lat: -23.555, lng: -46.733 }, weight: 2 },
  { location: { lat: -23.56041, lng: -46.7252 }, weight: 1.4 },
  { location: { lat: -23.55802, lng:  -46.72048 }, weight: 3},
  { location: { lat: -23.57284, lng:  -46.70642 }, weight: 1.8},
  { location: { lat: -23.55901, lng:  -46.72010 }, weight: 1.2},
  { location: { lat: -23.55950, lng:  -46.71900 }, weight: 1.8},
  { location: { lat: -23.55998, lng:  -46.71850 }, weight: 1.6},
  { location: { lat: -23.58758, lng:  -46.65567 }, weight: 2},
  { location: { lat: -23.564895123795928, lng:  -46.725863933095766 }, weight: 1},
  { location: { lat: -23.564465723383172, lng:  -46.73500853493441 }, weight: 1.6},
  { location: { lat: -23.559697245496334, lng:  -46.735226289531944 }, weight: 1.1},
  { location: { lat: -23.56758913163716, lng:  -46.731692464102814 }, weight: 1.4},
  { location: { lat: -23.560095862679532, lng:  -46.71526953170796 }, weight: 1},
  { location: { lat: -23.56151030545915, lng:  -46.72939429302866 }, weight: 1},

]

const googleApi = import.meta.env.VITE_GOOGLE_API

</script>

<template >
  <GoogleMap
    api-key="AIzaSyDUYQMh9jDe7mKbrfTBDjbnuqYsX3fUUFI"
    :libraries="['visualization']"
    style="width: 100%; height: 600px;"
    :center="inteli"
    :zoom="16"
    
    :disableDefaultUi="true"
  >
    <HeatmapLayer :options="{ 
      data: heatmapData,
      radius: 50,
      dissipating: false, 
      opacity: 0.4
       }" />
  </GoogleMap>
</template>
