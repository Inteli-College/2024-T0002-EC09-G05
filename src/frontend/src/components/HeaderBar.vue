<template>
  <Disclosure as="nav" class="bg-slate-100   border-botton">
    <div class="mx-auto max-w-7xl px-2 sm:px-6 lg:px-8">
      <div class="relative flex h-16 items-center justify-between">
        <a href="/" class="flex flex-shrink-0 items-center">
          <img
            class="h-8 w-auto"
            src="../assets/logo.png"
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

          <!-- Profile dropdown -->
          <Menu
            as="div"
            class="hover:outline hover:outline-offset-2 hover:outline-4 rounded-full relative ml-3"
          >
            <div>
              <MenuButton
                class="relative flex rounded-full bg-slate-100 text-sm focus:outline-none focus:ring-2 focus:ring-gray-700 focus:ring-offset-2 focus:ring-offset-bg-slate-100"
              >
                <span class="absolute -inset-1.5" />
                <span class="sr-only">Open user menu</span>
                <img
                  class="h-8 w-8 rounded-full"
                  src="https://images.unsplash.com/photo-1472099645785-5658abf4ff4e?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80"
                  alt=""
                />
              </MenuButton>
            </div>
            <transition
              enter-active-class="transition ease-out duration-100"
              enter-from-class="transform opacity-0 scale-95"
              enter-to-class="transform opacity-100 scale-100"
              leave-active-class="transition ease-in duration-75"
              leave-from-class="transform opacity-100 scale-100"
              leave-to-class="transform opacity-0 scale-95"
            >
              <MenuItems
                class="absolute right-0 z-10 mt-2 w-48 origin-top-right rounded-md bg-white py-1 shadow-lg ring-1 ring-gray-700 ring-opacity-5 focus:outline-none"
              >
                <MenuItem v-if="role==1" v-slot="{ active }">
                  <a 
                    href="/settings"
                    :class="[active ? 'bg-gray-100' : '', 'block px-4 py-2 text-sm text-gray-700']"
                    >Settings</a
                  >
                  
                </MenuItem>
                <MenuItem v-slot="{ active }">
                  <a
                  href="/auth"
                  @click="singout_handler"
                    :class="[active ? 'bg-gray-100' : '', 'block px-4 py-2 text-sm text-gray-700']"
                    >Sign out</a
                  >
                </MenuItem>
              </MenuItems>
            </transition>
          </Menu>
        </div>
      </div>
    </div>
  </Disclosure>
</template>

<script setup>
import { Disclosure, Menu, MenuButton, MenuItem, MenuItems } from '@headlessui/vue'
import { BellIcon } from '@heroicons/vue/24/outline'
import SearchBar from '../components/SearchBar.vue'
import { useCookies } from 'vue3-cookies'

var role = useCookies().cookies.get("role")
console.log(role)
function singout_handler(event) {
  useCookies().cookies.remove("authToken")
  
}
</script>
