<script setup>
import { GoogleMap, HeatmapLayer } from 'vue3-google-map'
const inteli = { lat: -23.555626840909746, lng: -46.73381383445283 }
const heatmapData = [
  //Ambas as formas adicionam um ponto ao heatmap. O 1° adiciona um ponto ponderado, ex -> 1 ponto equivalente a 2
  { location: { lat: -23.5556, lng: -46.7338 }, weight: 0.4 },
  { location: { lat: -23.555, lng: -46.733 }, weight: 0.8 },
  { location: { lat: -23.56041, lng: -46.7252 }, weight: 0.4 },
  { location: { lat: -23.55802, lng:  -46.72048 }, weight: 0.3},
  { location: { lat: -23.57284, lng:  -46.70642 }, weight: 0.3},
  { location: { lat: -23.55901, lng:  -46.72010 }, weight: 0.3},
  { location: { lat: -23.55950, lng:  -46.71900 }, weight: 0.3},
  { location: { lat: -23.55998, lng:  -46.71850 }, weight: 0.3},
  { location: { lat: -23.58758, lng:  -46.65567 }, weight: 0.2},
  { location: { lat: -23.564895123795928, lng:  -46.725863933095766 }, weight: 0.1},
  { location: { lat: -23.564465723383172, lng:  -46.73500853493441 }, weight: 0.6},
  { location: { lat: -23.559697245496334, lng:  -46.735226289531944 }, weight: 0.3},
  { location: { lat: -23.56758913163716, lng:  -46.731692464102814 }, weight: 0.4},
  { location: { lat: -23.560095862679532, lng:  -46.71526953170796 }, weight: 0.1},
  { location: { lat: -23.56151030545915, lng:  -46.72939429302866 }, weight: 0.1},
  { location: { lat: -23.57232609, lng: -46.70090389 }, weight: 0.5 },
  { location: { lat: -23.57228831, lng: -46.70093386 }, weight: 0.5 },
  { location: { lat: -23.57223959, lng: -46.70097251 }, weight: 0.5 },
  { location: { lat: -23.57213969, lng: -46.70105175 }, weight: 0.5 },
  { location: { lat: -23.57212049, lng: -46.70106698 }, weight: 0.5 },
  { location: { lat: -23.57175243, lng: -46.70135894 }, weight: 0.5 },
  { location: { lat: -23.57171968, lng: -46.70138492 }, weight: 0.5 },
  { location: { lat: -23.57166285, lng: -46.70142055 }, weight: 0.5 },
  { location: { lat: -23.56883314, lng: -46.70297658 }, weight: 0.5 },
  { location: { lat: -23.5660097, lng: -46.70398516 }, weight: 0.1 },
  { location: { lat: -23.56288432, lng: -46.70432171 }, weight: 0.1 },
  { location: { lat: -23.56264154, lng: -46.7039744 }, weight: 0.1 },
  { location: { lat: -23.56194836, lng: -46.70315198 }, weight: 0.1 },
  { location: { lat: -23.559366617271948, lng: -46.712009802495594 }, weight: 0.4 },
  { location: { lat: -23.553875059484884, lng: -46.70219564095203 }, weight: 0.4 },
  { location: { lat: -23.552906215120533, lng: -46.692646863905566 }, weight: 0.4 },
  { location: { lat: -23.566611093941315, lng: -46.69486848049209 }, weight: 0.4 },
  { location: { lat: -23.586160225898023, lng: -46.68237097790382 }, weight: 0.4 },
  { location: { lat: -23.538548, lng: -46.6360091 }, weight: 0.4 },
  { location: { lat: -23.5430556, lng: -46.64388890000001 }, weight: 0.4 },
  { location: { lat: -23.555038, lng: -46.6621 }, weight: 0.4 },
  { location: { lat: -23.567605, lng: -46.69279299999999 }, weight: 0.4 },
  { location: { lat: -23.5659344, lng: -46.7010872 }, weight: 0.4 },
  { location: { lat: -23.565, lng: -46.7069444 }, weight: 0.4 },
  { location: { lat: -23.647066, lng: -46.63954 }, weight: 0.4 },
  { location: { lat: -23.636214, lng: -46.641141 }, weight: 0.4 },
  { location: { lat: -23.6244547, lng: -46.6379525 }, weight: 0.4 },
  { location: { lat: -23.617243, lng: -46.6343336 }, weight: 0.4 },
  { location: { lat: -23.61069, lng: -46.637852 }, weight: 0.4 },
  { location: { lat: -23.599225, lng: -46.63665899999999 }, weight: 0.4 },
  { location: { lat: -23.5870563, lng: -46.63574370000001 }, weight: 0.4 },
  { location: { lat: -23.581376, lng: -46.638388 }, weight: 0.4 },
  { location: { lat: -23.575922, lng: -46.64141 }, weight: 0.4 },
  { location: { lat: -23.568994, lng: -46.639858 }, weight: 0.4 },
  { location: { lat: -22.8, lng: -47.766667 }, weight: 0.4 },
  { location: { lat: -23.5555348, lng: -46.6359564 }, weight: 0.4 },
  { location: { lat: -23.53862032, lng: -46.70807534 }, weight: 0.4 },
]
const googleApi = import.meta.env.VITE_GOOGLE_API
</script>
<template >
  <GoogleMap
    :api-key="googleApi"
    :libraries="['visualization']"
    style="width: 100%; height: 600px;"
    :center="inteli"
    :zoom="16"
    :disableDefaultUi="true"
  >
    <HeatmapLayer :options="{
      data: heatmapData,
      radius: 5 ,
      dissipating: false,
      opacity: 0.4,
      maxIntensity: 0.2,
      // gradient: ['#C2C2C2', '#8FFF66', '#EBFF66']
       }" />
  </GoogleMap>
</template>