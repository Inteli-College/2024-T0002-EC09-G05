<script setup>
import { CCard, CCardImage, CCardBody, CCardText, CCardTitle, CCol } from '@coreui/vue'
</script>

<template>
<CCard class="mb-3">
    <CRow class="g-0">
        <CCol :md="4" class="max-w-12">
            <!-- <CCardImage class="rounded-0" src="../assets/logo.png" /> -->
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="w-6 h-6">
                <path strokeLinecap="round" strokeLinejoin="round" d="M16.5 12a4.5 4.5 0 1 1-9 0 4.5 4.5 0 0 1 9 0Zm0 0c0 1.657 1.007 3 2.25 3S21 13.657 21 12a9 9 0 1 0-2.636 6.364M16.5 12V8.25" />
            </svg>

        </CCol>
        <CCol :md="8">
        <CCardBody>
            <CCardTitle>Butantã teve sua maior temperatura dos últimos anos</CCardTitle>
            <!-- <CCardText>.</CCardText> -->
            <CCardText><small class="text-body-secondary" style="color: #718EBF;">25 de Março de 2024</small></CCardText>
        </CCardBody>
        </CCol>
    </CRow>
</CCard>
</template>
