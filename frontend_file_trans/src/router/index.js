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
    redirect: '/send'
  },
  {
    path: '/send',
    name: 'Send',
    component: SendPage
  },
  {
    path: '/receive', 
    name: 'Receive',
    component: ReceivePage
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router