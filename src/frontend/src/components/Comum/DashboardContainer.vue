<script setup>

    import SideBar from '../SideBar.vue' 
    import MainChart from '../MainChart.vue';
    import GrapchA from '../GrapchA.vue';
    import GrapchB from '../GrapchB.vue';
    import HeatMap from '../HeatMap.vue';
    import GraphRain from '../GraphRain.vue'

</script>

<template>
    <div class="grid grid-cols-6  w-screen h-auto min-h-96 ">
        <div class="bg-secudary  sideMenu"></div>
        <SideBar />

        <div class="xl:col-start-2 xl:col-end-9 flex-col min-h-screen col-start-1 h-fit  col-end-9 w-full">
            
            
            <div class="flex flex-wrap gap-10 h-full max-w-full m-10 items-center align-items: flex-start" >
                <div :class="`w-full flex justify-center  items-center min-h-52 h-auto  bg-white soft-shadow text-center`">
                   <HeatMap />
                </div>
                <div :class="`w-96 flex justify-center  items-center min-h-52 h-auto bg-white soft-shadow text-center`">
                    <div class="flex flex-col w-auto">
                        <h1 class="m-3">MainChart</h1>
                        <MainChart />
                    </div>
                </div>
                <div :class="`w-3/12 flex justify-center  items-center min-h-52 h-auto  bg-white soft-shadow text-center`">
                    <div class="flex flex-col w-auto">
                        <h1 class="m-3">GrapchA</h1>        
                        <GrapchA />
                    </div>
                </div>
                <div :class="`w-96 flex justify-center  items-center min-h-52 h-auto bg-white soft-shadow text-center`">
                    <div class="flex flex-col w-auto">
                        <h1 class="m-3">GraphRain</h1>
                        <GraphRain  />
                    </div>
                </div>

            </div> 
        </div>
    </div>
</template>
  
