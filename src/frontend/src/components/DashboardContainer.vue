<script setup>

    import SideBar from './SideBar.vue' 
    import MainChart from './MainChart.vue';
    import GrapchA from './GrapchA.vue';
    import GrapchB from './GrapchB.vue';
    
    import HeatMap from './HeatMap.vue';
    
    import axios from 'axios';
    import { useCookies } from 'vue3-cookies'
    import { ref, onMounted } from 'vue';
    import draggable from 'vuedraggable';

    const elements = ref(null);

    const authTokenKey = 'authToken';
    

    const get_elements = async () =>{
        try {
            const url = '/api/platform/getComponents';
            
            const response = await axios.post( url,  {
            Id: Number(useCookies().cookies.get(authTokenKey)),
            });

            console.log(response.data.elements)
            if  (response.data.elements != null){
                return response.data.elements}

        } catch (error) {
            console.error('Ocorreu um erro:', error);
            
        }
        return [{"Name": "MainChart","Index": 1,"Value": "w-96"}]
    } 


   async function format_data() {
       var  raw_elements = await get_elements()
       elements.value = raw_elements
    }

    onMounted(() => {
    format_data();
  });


</script>

<template>
    <div class="grid grid-cols-6  w-screen h-auto min-h-96 ">
        <div class="bg-secudary  sideMenu"></div>
        <SideBar />

        <div class=" xl:col-start-2 xl:col-end-9 flex-col min-h-screen col-start-1 h-fit  col-end-9 w-full">
            
            <draggable v-model="elements" tag="div" class="flex flex-wrap gap-10 h-full max-w-full m-10 items-center align-items: flex-start" :animation="300">
                <template #item="{ element: element_ }">
                    <div :class="`${element_.Value} flex justify-center cursor-grab  items-center min-h-52 h-auto bg-slate-50 soft-shadow text-center`">

                            
                            <MainChart v-if="element_.Name == 'MainChart'" />
                            <GrapchA v-if="element_.Name == 'GrapchA'" />
                            <GrapchB v-if="element_.Name == 'GrapchB'" />
                            <HeatMap v-if="element_.Name == 'HeatMap'" />

                    </div>
                </template>
            </draggable> 
        </div>
    </div>
</template>
  
