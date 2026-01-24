import { createWebHashHistory, createRouter } from 'vue-router';

import Login from '../views/Login.vue';
import About from '../views/About.vue';
import Home from '../views/Home.vue';

const routes = [
    { path: '/', component: Login },
    { path: '/home', component: Home },
    { path: '/about', component: About },
];

const router = createRouter({
    history: createWebHashHistory(),
    routes,
});

router.beforeEach((to, from, next) => {
    const isAuthenticated = localStorage.getItem('authToken');

    if (to.meta.requiresAuth && !isAuthenticated) {
        next("/");
    } else {
        next();
    }
});

export default router;
