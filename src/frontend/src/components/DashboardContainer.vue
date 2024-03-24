<script setup>

    import SideBar from './SideBar.vue' 
    import MainChart from './MainChart.vue';
    import GrapchA from './GrapchA.vue';
    import GrapchB from './GrapchB.vue';
    import GrapchC from './GrapchC.vue';
    
    import axios from 'axios';
    import { ref, onMounted } from 'vue';
    import draggable from 'vuedraggable';

    const elements = ref(null);

    const get_elements = async () =>{
        try {
            const url = '/api/platform/getComponents';
            
            const response = await axios.post( url,  {
            id: window.user,
            token: "token",
            });

            console.log(response.data.elements)
            if  (response.data != null){
                return response.data.elements}

        } catch (error) {
            console.error('Ocorreu um erro:', error);
            
        }
        return ["Tetse"]
    } 


   async function format_data() {
       var  raw_elements = await get_elements()
       elements.value = [
            raw_elements[0],
            raw_elements[1],
            raw_elements[2],
       ]
    }

    onMounted(() => {
    format_data();
  });


</script>

<template>
    <div class="grid grid-cols-6  w-screen h-auto min-h-96 ">
        <div class=" bg-slate-100 sideMenu"></div>
        <SideBar />

        <div class="bg-white xl:col-start-2 xl:col-end-9 flex-col min-h-screen col-start-1 h-fit  col-end-9 w-full">
            
            <draggable v-model="elements" tag="div" class="flex flex-wrap gap-10 h-full max-w-full m-10 items-center align-items: flex-start" :animation="300">
                <template #item="{ element: element_ }">
                    <div :class="`${element_.Value} flex justify-center  items-center min-h-52 h-auto bg-slate-700 text-center`">
                        <div class="flex flex-col w-auto">
                            <h1 class="m-3">{{element_.Index}}</h1>
                            <!-- <MainChart v-if="element_.Name == 'MainChart'" /> -->
                            <GrapchA v-if="element_.Name == 'GrapchA'" />
                            <GrapchB v-if="element_.Name == 'GrapchB'" />
                            <GrapchC v-if="element_.Name == 'GrapchC'" />
                        </div>
                    </div>
                </template>
            </draggable> 
        </div>
    </div>
</template>
  
