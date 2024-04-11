<script setup>
import { ref } from "vue";
import Dropdown from 'primevue/dropdown';

defineProps({
  userId: Number,
  roles: String,
  directorates: Number,
  elements: Number,
  props: Number,

  textLabel: String,
})

const selectedValue = ref();
const values = ref([
    { name: 'User1', code: 'u1' },
    { name: 'User2', code: 'u2' },
    { name: 'User3', code: 'u3' },
    { name: 'User4', code: 'u4' },

]);
</script>

<template>
    <div class="card flex justify-content-center">
        <Dropdown v-model="selectedValue" :options="values" optionLabel="name" placeholder={{ textLabel }} class="w-full md:w-14rem" />
    </div>
</template>