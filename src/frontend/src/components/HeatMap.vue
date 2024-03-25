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
  { location: { lat: -23.58758, lng:  -46.65567 }, weight: 2},

]
</script>

<template>
  <GoogleMap
    api-key="AIzaSyAm7xZj1d0fQRv0sYl_Urq4wggecZPaLlE"
    :libraries="['visualization']"
    style="width: 100%; height: 500px"
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