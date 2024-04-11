<script setup>
import { CCard, CCardText, CCardTitle, CCol, CRow } from '@coreui/vue'
</script>

<template>
<CCard class="mb-1">
    <CRow class="g-0 flex"> 
        <CCol class="flex items-center"> 
            <img src="../assets/icon3.svg" class="custom-image">
        </CCol>
        <CCol class="ml-4">
            <CCardTitle >Butantã teve sua maior temperatura dos últimos anos</CCardTitle>    
            <CCardText><small class="text-body-secondary" style="color: #718EBF;">5 de Abril de 2024</small></CCardText>
        </CCol>
    </CRow>
</CCard>
<CCard class="mb-1">
    <CRow class="g-0 flex"> 
        <CCol class="flex"> 
            <img src="../assets/icon2.svg" class="custom-image">
        </CCol>
        <CCol class="ml-4">
            <CCardTitle >Centro de São Paulo teve sua menor temperatura dos últimos anos</CCardTitle>    
            <CCardText><small class="text-body-secondary" style="color: #718EBF;">5 de Abril de 2024</small></CCardText>
        </CCol>
    </CRow>
</CCard>
<CCard class="mb-1">
    <CRow class="g-0 flex"> 
        <CCol class="flex"> 
            <img src="../assets/icon1.svg" class="custom-image">
        </CCol>
        <CCol class="ml-4">
            <CCardTitle >Aumento de gastos com controle atmosférico</CCardTitle>    
            <CCardText><small class="text-body-secondary" style="color: #718EBF;">5 de Abril de 2024</small></CCardText>
        </CCol>
    </CRow>
</CCard>
</template>

<style scoped>
.custom-image {
    width: 100px; 
    height: 100px; 
}
</style>
