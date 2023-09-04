// src/router/index.js
import {createRouter, createWebHistory} from 'vue-router';
import ApplicationHome from '../components/ApplicationHome.vue';
import UserHome from '../components/UserHome.vue';

const routes = [
    {path: '/', component: ApplicationHome},
    {path: '/home', component: UserHome}
];

const router = createRouter({
    history: createWebHistory(),
    routes
});

export default router;
