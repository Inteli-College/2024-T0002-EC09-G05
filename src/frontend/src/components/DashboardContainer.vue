
<template>
    <div class="grid grid-cols-6  w-screen h-auto min-h-96 ">
        <div class="sideMenu bg-slate-100 w-full min-h-full  xl:block hidden">
            <div class="flex flex-col gap-10 items-center m-4 bg-slate-700 h-auto"> 
                <div class="h-10 w-full text-center bg-amber-200">componete da infos</div>
                <div class="h-10 w-full text-center bg-amber-200">componete da infos</div>
                <div class="h-10 w-full text-center bg-amber-200">componete da infos</div>
            </div>
        </div>
        <div class="bg-stone-400 xl:col-start-2 xl:col-end-9 flex-col min-h-screen col-start-1  col-end-9 w-full h-7">
            Grapcs Area...
            <div class="flex flax-col max-w-full m-10 items-center justify-center">
                <div class="w-full h-96 bg-slate-500 rounded-md">
                    Grapcs
                </div>
            </div>
        </div>
    </div>
</template>
  
  <script setup>
  import { Disclosure, Menu, MenuButton, MenuItem, MenuItems } from '@headlessui/vue'
  
  </script>
