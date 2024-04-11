<script setup>
    import InputText from 'primevue/inputtext';
    
    defineProps({
        label: String,

    })    
</script>

<template>
    <div class="flex flex-column gap-2">
        <label for="username">{{  label  }}</label>
        <InputText id="username" v-model="value" aria-describedby="username-help" />
        <!-- <small id="username-help"></small> -->
    </div>
    
</template>


