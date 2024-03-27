<template>
    <div class="container">
      <div class="block" v-for="(block, index) in blocks" :key="index">
        <h2>{{ block.title }}</h2>
        <p>{{ block.content }}</p>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        blocks: [
          { title: 'Data', content: 'Conteúdo do bloco 1' },
          { title: 'Hora', content: 'Conteúdo do bloco 2' },
          { title: 'Max Temperatura', content: 'Conteúdo do bloco 3' },
          { title: 'Min Temperatura', content: 'Conteúdo do bloco 4' }
        ]
      };
    },
    mounted() {
    this.fetchData();
  },
  methods: {
    async fetchData() {
      try {
        const response = await fetch('http://localhost:8080/data/getDataByRelativeTime', {
          method: 'POST', // ou 'GET', 'PUT', 'DELETE', etc., dependendo do tipo de requisição
          mode: 'no-cors',
          headers: {
            // 'Content-Type': 'application/json',
            // 'Accept': '*/*',  
            // 'Authorization': 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MTI0MDQzNTMsInJvbGUiOjEsInVzZXJfaWQiOjF9.KZmnOXt9FSL_iOG4r5eDRicwagHs-uLYAUOqiVfQxDY',
            // Outros cabeçalhos conforme necessário
          },
          body: JSON.stringify({
            // Se você tiver dados para enviar no corpo da requisição
            start_time: '60',
            end_time: '0',
            aggregator: '300',
            fields: '[temp]',
            sensors: '[Paulista]',
            // Adicione outras chaves e valores conforme necessário
          })
        });
        
        if (!response.ok) {
          throw new Error('Erro ao carregar dados da API');
        }
        
        const data = await response.json();
        this.responseData = data;
      } catch (error) {
        console.error('Erro:', error);
      }
    }
  }
  };
  </script>
  
  <style scoped>
  .container {
    width: 85%;
    display: flex;
    justify-content: center;

    box-sizing: border-box;
    background: rgba(123, 123, 123, 0.8);
    height: 125px;
    /* width: 200px; */
    margin: 0px 0px 1.5rem 0px;
    /* padding: 0px; */
    font-size: 1.25rem;
    appearance: none;

    user-select: none;
    overflow: hidden;

    font-family: Roboto, sans-serif;
    font-weight: 700;
    color: #fff;
    background-color: #79797987;

    border: none;
    box-shadow: rgb(0, 0, 0) 0px 5px 25px -5px;
    border-radius: 8px 8px 0px 0px;
    /* transition : 738ms; */
    /* transform: translateY(0); */
    /* display: flex; */
    flex-direction: row;
    align-items: center;
    /* cursor: pointer; */
    /* text-transform: lowercase;   */
  }
  
  .block {
    width: 75%; /* Ajuste conforme necessário */
    height: 100%;
    background-color: #79797987;
    padding: 15px 30px;
  }
  
  h2 {
    font-size: 14px;
  }
  
  p {
    padding-top: 1.25rem;
    font-size: 24px;
  }
  </style>
  