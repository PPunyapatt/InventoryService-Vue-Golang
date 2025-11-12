import { createMemoryHistory, createRouter } from 'vue-router'

import InventoryList from './pages/InventoryList.vue'


const routes = [
    {path: '/', component: InventoryList}
]

export const router = createRouter({
  history: createMemoryHistory(),
  routes,
})