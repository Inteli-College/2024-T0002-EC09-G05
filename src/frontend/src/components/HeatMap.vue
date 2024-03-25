<script setup>
import { GoogleMap, HeatmapLayer } from 'vue3-google-map'

const inteli = { lat: -23.555626840909746, lng: -46.73381383445283 }

const heatmapData = [
  //Ambas as formas adicionam um ponto ao heatmap. O 1° adiciona um ponto ponderado, ex -> 1 ponto equivalente a 2
  { location: { lat: -23.5556, lng: -46.7338 }, weight: 2 },
  { lat: -23.555, lng: -46.733 },
  { lat: -23.56041, lng: -46.7252 },
  { location: { lat: -23.55802, lng:  -46.72048 }, weight: 3},
  { location: { lat: -23.57284, lng:  -46.70642 }, weight: 1.8},
  { location: { lat: -23.55901, lng:  -46.72010 }, weight: 3},
  { location: { lat: -23.55950, lng:  -46.71900 }, weight: 2},
  { location: { lat: -23.55998, lng:  -46.71850 }, weight: 3},
  { location: { lat: -23.58758, lng:  -46.65567 }, weight: 2},

]

const googleApi = import.meta.env.VITE_GOOGLE_API

</script>

<template >
  {{ googleApi }}
  <GoogleMap
    :api-key="googleApi"
    :libraries="['visualization']"
    style="width: 100%; height: 600px; width: 100%"
    :center="inteli"
    :zoom="14"
    
    :disableDefaultUi="true"
  >
    <HeatmapLayer :options="{ 
      data: heatmapData,
      radius: 75,
      dissipating: true, 
      opacity: 0.4
       }" />
  </GoogleMap>
</template>
