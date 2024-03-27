<script setup>
import SideBar from './SideBar.vue'
import MainChart from './MainChart.vue';
import HeatMap from './HeatMap.vue'

</script>

<template>
    <div class="grid grid-cols-6  w-screen h-auto min-h-96 ">
        <div class=" bg-slate-100 sideMenu"></div>
        <SideBar />

        <div class="bg-white xl:col-start-2 xl:col-end-9 flex-col min-h-screen col-start-1 h-fit  col-end-9 w-full">
            
            <div class="flex flex-col gap-10 h-full max-w-full m-10 items-center justify-center align-items: flex-start">
                <div style="width: 100%; ">
                    <HeatMap />
                </div>
                <div style="width: 100%; max-height: 300px; height: 300px;">
                    <MainChart />
                </div>
            </div>
        </div>
    </div>
</template>
  
