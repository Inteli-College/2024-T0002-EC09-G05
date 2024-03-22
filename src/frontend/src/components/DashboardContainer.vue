
<script setup>

    import SideBar from './SideBar.vue' 
    import MainChart from './MainChart.vue';
    
    import axios from 'axios';
    import { ref, onMounted } from 'vue';


    const elements = ref(null);

    async function get_layout(){
        try {
            const url = '/api/platform/getComponents';
            
            const response = await axios.post( url,  {
            id: window.user,
            token: "token",
            });

            console.log(response.data.elements)
            if  (response.data != null){
                elements.value = response.data.elements}

        } catch (error) {
            console.error('Ocorreu um erro:', error);
            
        }
        return ["Tetse"]
    } 

    onMounted(() => {
    get_layout();
  });


</script>

<template>
    <div class="grid grid-cols-6  w-screen h-auto min-h-96 ">
        <div class=" bg-slate-100 sideMenu"></div>
        <SideBar />

        <div class="bg-white xl:col-start-2 xl:col-end-9 flex-col min-h-screen col-start-1 h-fit  col-end-9 w-full">
            
            <div class="flex flex-col gap-10 h-full max-w-full m-10 items-center align-items: flex-start">
                <!-- Os gráficos são renderizados aqui -->
                <div style="width: 100%; max-height: 300px; height: 300px;">
                    <MainChart />
                </div>
                <div class="w-full h-96 bg-slate-700 text-center" v-for="element in elements" >{{element.Name}}</div>

            </div>
        </div>
    </div>
</template>
  
