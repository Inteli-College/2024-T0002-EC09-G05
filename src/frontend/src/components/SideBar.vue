<script setup>
import InsightCard from './InsightCard.vue'
import SideBarTitleSection from './SideBarTitleSection.vue'
import SideBarStats from './SideBarStats.vue'
import Btn from './Btn.vue'
</script>
<template>
        <div class="SideBar min-h-full bg-secudary xl:block max-w-sm fixed">      
            <div class="flex flex-col gap-10 items-center m-6 h-auto"> 
                <div class="flex w-full items-center justify-between">
                    <SideBarTitleSection title="Insights Recentes" class="font-bold text-xl"/>
                    <Btn />
                </div>
                <InsightCard/>
                <div>
                    <SideBarTitleSection title="Estatisticas Rápidas" class="font-bold text-xl"/>

                </div>
                <SideBarStats/>
            </div>
        </div>
</template>
