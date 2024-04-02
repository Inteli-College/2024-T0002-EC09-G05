<script setup>
const props = defineProps({
    title: String
})
</script>

<template>
    <h1 class="text-stone-600">{{ title }}</h1>
</template>