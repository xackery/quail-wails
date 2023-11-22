import { createRouter, createWebHashHistory } from 'vue-router';
import DashboardView from './views/DashboardView.vue'

const routes = [
    {path:"/", name:"Home" , component: DashboardView },
]

export const router = createRouter({
    history: createWebHashHistory(),
    routes,
})

