<script setup>

import { ref } from 'vue';

const isModalOpen = ref(false);

const closeModal = () => {
  isModalOpen.value = false;
};


</script>

<template>
    <div class="flex flex-col h-screen items-center justify-center w-screen">
        <h3 class="text-7xl">Erro 404</h3>
        <div class="flex flex-wrap h-screen w-auto">
            <h1>Verifique a rota</h1>
        </div>

    </div>
</template>

<style scoped>
/* Modal Styles */
.modal {
  display: flex;
  position: fixed;
  z-index: 999;
  left: 0;
  top: 0;
  width: 100%;
  height: 100%;
  align-items: center;
  justify-content: center;
  background-color: rgba(0, 0, 0, 0.4);
}

/* Modal Content */
.modal-content {
  background-color: #fefefe;
  padding: 20px;
  border-radius: 5px;
}

/* Close Button */
.close {
  color: #aaa;
  float: right;
  font-size: 28px;
  font-weight: bold;
  cursor: pointer;
}
</style>