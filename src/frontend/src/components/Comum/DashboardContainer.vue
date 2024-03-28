<script setup>

    import SideBar from '../SideBar.vue' 
    import MainChart from '../MainChart.vue';
    import GrapchA from '../GrapchA.vue';
    import GrapchB from '../GrapchB.vue';
    import GrapchC from '../GrapchC.vue';

</script>

<template>
    <div class="grid grid-cols-6  w-screen h-auto min-h-96 ">
        <div class=" bg-slate-100  sideMenu"></div>
        <SideBar />

        <div class="xl:col-start-2 xl:col-end-9 flex-col min-h-screen col-start-1 h-fit  col-end-9 w-full">
            
            <div class="flex flex-wrap gap-10 h-full max-w-full m-10 items-center align-items: flex-start" >
                <div :class="`w-96 flex justify-center  items-center min-h-52 h-auto bg-white soft-shadow text-center`">
                    <div class="flex flex-col w-auto">
                        <h1 class="m-3">MainChart</h1>
                        <MainChart />
                    </div>
                </div>
                <div :class="`w-96 flex justify-center  items-center min-h-52 h-auto  bg-white soft-shadow text-center`">
                    <div class="flex flex-col w-auto">
                        <h1 class="m-3">GrapchA</h1>        
                        <GrapchA />
                    </div>
                </div>
                <div :class="`w-96 flex justify-center  items-center min-h-52 h-auto  bg-white soft-shadow text-center`">
                    <div class="flex flex-col w-auto">
                        <h1 class="m-3">GrapchB</h1>
                        <GrapchB  />
                    </div>
                </div>
                <div :class="`w-96 flex justify-center  items-center min-h-52 h-auto  bg-white soft-shadow text-center`">
                    <div class="flex flex-col w-auto">
                        <h1 class="m-3">GrapchC</h1>
                        <GrapchC />
                    </div>
                </div>
            </div> 
        </div>
    </div>
</template>
  
