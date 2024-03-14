<script setup>
import InsightCard from './InsightCard.vue'
</script>

<template>
    <div class="grid grid-cols-6  w-screen h-auto min-h-96 ">
        <div class="sideMenu bg-slate-100 w-full min-h-full  xl:block hidden">
            <div class="flex flex-col gap-10 items-center m-6 h-auto"> 
                <InsightCard/>
                <InsightCard/>
                <InsightCard/>
            </div>
        </div>
        <div class="bg-stone-400 xl:col-start-2 xl:col-end-9 flex-col min-h-screen col-start-1 h-fit  col-end-9 w-full">
            
            <div class="flex flex-col gap-10 h-full max-w-full m-10 items-center">
                <!-- Componentes dos gráficos aqui -->
            </div>
            </div>
        </div>
    </div>
</template>
  
