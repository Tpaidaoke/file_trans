// router/index.js
import {
	createRouter,
	createWebHistory
} from 'vue-router'
import SendPage from '@/views/SendPage.vue'
import ReceivePage from '@/views/ReceivePage.vue'

const routes = [
  {
    path: '/',
    redirect: '/sendPackage'
  },
  {
    path: '/sendPackage',
    name: 'Send',
    component: SendPage
  },
  {
    path: '/receivePackage', 
    name: 'Receive',
    component: ReceivePage
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router