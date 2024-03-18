<script setup>
import InsightCard from './InsightCard.vue'
import SideBarTitleSection from './SideBarTitleSection.vue'
import SideBarStats from './SideBarStats.vue'
</script>
<template>
        <div class="SideBar  bg-slate-100  min-h-full  xl:block max-w-sm fixed">      
            <div class="flex flex-col gap-10 items-center m-6 h-auto"> 
                <SideBarTitleSection title="Insights Recentes"/>
                <InsightCard/>
                <SideBarTitleSection title="Estatisticas Rápidas"/>
                <SideBarStats/>
            </div>
        </div>
</template>
