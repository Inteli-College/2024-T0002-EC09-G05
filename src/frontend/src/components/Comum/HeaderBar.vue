<template>
  <Disclosure as="nav" class="bg-slate-100 border-botton">
    <div class="mx-auto max-w-7xl px-2 sm:px-6 lg:px-8">
      <div class="relative flex h-16 items-center justify-between">
        <a href="/" class="flex flex-shrink-0 items-center">
          <img
            class="h-8 w-auto"
            src="../../assets/logo.png"
            alt="Your Company"
          />
        </a>
        <div class="flex items-center w-full sm:hidden">
          <!-- Mobile menu button-->
          <SearchBar class="w-50 ml-11 transition-all" />
        </div>

        <div
          class="absolute inset-y-0 right-0 flex items-center pr-2 sm:static sm:inset-auto sm:ml-6 sm:pr-0"
        >
          <SearchBar class="hidden sm:ml-6 sm:block mr-3" />
          <button
            type="button"
            class="relative rounded-full bg-slate-100 p-1 text-gray-400 hover:text-white focus:outline-none focus:ring-2 focus:ring-gray-700 hover:bg-gray-700 focus:ring-offset-2 focus:ring-offset-bg-slate-100"
          >
            <span class="absolute -inset-1.5" />
            <span class="sr-only">View notifications</span>
            <BellIcon class="h-6 w-6" aria-hidden="true" />
          </button>

          <a href="/auth"  class="mx-4 hover:bg-slate-600 rounded-full bg-slate-700 px-10 text-gray-50 py-2">Sigin</a href="/auth">
        </div>
      </div>
    </div>
  </Disclosure>
</template>

<script setup>
import { Disclosure, Menu, MenuButton, MenuItem, MenuItems } from '@headlessui/vue'
import { BellIcon } from '@heroicons/vue/24/outline'
import SearchBar from '../SearchBar.vue'
</script>
