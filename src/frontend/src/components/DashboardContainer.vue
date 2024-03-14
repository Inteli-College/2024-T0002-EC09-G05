<script setup>
import InsightCard from './InsightCard.vue'
</script>

<template>
    <div class="grid grid-cols-6  w-screen h-auto min-h-96 ">
        <div class="sideMenu bg-slate-100 w-full min-h-full  xl:block hidden">
            <div class="flex flex-col gap-10 items-center m-4 bg-slate-700 h-auto"> 
                <!-- Componentes de info aqui -->
                <InsightCard/>
            </div>
        </div>
        <div class="bg-stone-400 xl:col-start-2 xl:col-end-9 flex-col min-h-screen col-start-1  col-end-9 w-full h-7">
            Grapcs Area...
            <div class="flex flax-col max-w-full m-10 items-center justify-center">
                <div class="w-full h-96 bg-slate-500 rounded-md">
                    <!-- Componentes de gráficos aqui -->
                    Grapcs
                </div>
                
            </div>
        </div>
    </div>
</template>
  
