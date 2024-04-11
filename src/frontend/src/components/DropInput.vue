<template>
    <CButton color="primary" @click="() => { visibleLiveDemo = true }">Launch demo modal</CButton>
    <CModal 
      :visible="visibleLiveDemo"
      @close="() => { visibleLiveDemo = false }"
      aria-labelledby="LiveDemoExampleLabel"
    >
      <CModalHeader>
        <CModalTitle id="LiveDemoExampleLabel">Modal title</CModalTitle>
      </CModalHeader>
      <CModalBody>Woohoo, you're reading this text in a modal!</CModalBody>
      <CModalFooter>
        <CButton color="secondary" @click="() => { visibleLiveDemo = false }">
          Close
        </CButton>
        <CButton color="primary">Save changes</CButton>
      </CModalFooter>
    </CModal>
</template>
  <script>


  import { CButton,  CModal, CModalHeader, CModalTitle, CModalFooter, CModalBody} from '@coreui/vue';
    export default {
      data() {
        return { 
          visibleLiveDemo: false,
        }
      }
    }
  </script>