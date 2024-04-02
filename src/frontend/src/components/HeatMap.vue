<script setup>
import { GoogleMap, HeatmapLayer } from 'vue3-google-map'
// import api from "../../services/api"
import { defineComponent, onMounted, ref } from 'vue';
import axios from 'axios'

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

const url = '/api/data/getDataByRelativeTime'
const sensors = ref([])
const fetchData = async () =>
  axios
    .post(url, {
      "start_time": 60,
      "end_time": 0,
      "aggregator": 60,
      "fields": ["temp"],
      "sensors": ["Pinheiros",]
    }, {
      headers: {
        "Access-Control-Allow-Origin": "*",
        "Access-Control-Allow-Headers": "*"
      }
    })
    .then((response) => {
      (sensors.value.data = response.data[0].data)
      // return { sensors }
      // console.log(sensors.value.data[3].time)
      heatmapData.forEach(item => {
        console.log(item.weight)
        console.log(sensors.value.data[1])
        // item.weight = sensors.value.data[1].value
        // console.log(item.weight)
      })
    })
    onMounted(fetchData)

    // heatmapData.forEach(item => {
    //   console.log(sensors.value.data[0])
    //   // console.log(sensors.value.data)
    //   // item.weight = sensors.value.data[item].value
    //   console.log(item.weight)
    // })


</script>

<template >
  <!-- <div v-for="(pokemon, i) in pokemons" :key="i">
    <p>{{ pokemon.name }}</p>
  </div> -->
  <GoogleMap
    :api-key="googleApi"
    :libraries="['visualization']"
    style="width: 100%; height: 600px; width: 100%"
    :center="inteli"
    :zoom="16"
    
    :disableDefaultUi="true"
  >
    <HeatmapLayer :options="{ 
      data: heatmapData,
      radius: 220,
      dissipating: true, 
      opacity: 0.4
       }" />
  </GoogleMap>
</template>

<style scoped>
.custom-btn {
  box-sizing: border-box;
  background: rgba(123, 123, 123, 0.8);
  height: 125px;
  width: 200px;
  margin: 0px 0px 1.5rem 0px;
  /* padding: 0px; */
  font-size: 1.25rem;
  appearance: none;

  user-select: none;
  overflow: hidden;

  font-family: Roboto, sans-serif;
  font-weight: 700;
  font-size: 14px;
  color: #fff;
  background-color: #79797987;
  padding: 15px 30px;
  border: none;
  box-shadow: rgb(0, 0, 0) 0px 5px 25px -5px;
  border-radius: 8px 8px 0px 0px;
  /* transition : 738ms; */
  /* transform: translateY(0); */
  display: flex;
  flex-direction: row;
  align-items: center;
  cursor: pointer;
  text-transform: lowercase;
}
</style>
