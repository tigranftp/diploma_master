import {createApp,} from 'vue'
import {createRouter, createWebHistory} from "vue-router";
import App from './App.vue'

const router = createRouter({
    history: createWebHistory(),
    routes: [{
        name: "Main",
        path: '/',
        component: () => import("./components/main.vue"),
    },
        {
            name: "Login",
            path: '/login',
            component: () => import("./components/login.vue"),
        },
        {
            name: "SignUp",
            path: '/sign_up',
            component: () => import("./components/sign_up.vue"),
        },
        {
            name: "CreateModel",
            path: '/create_model',
            component: () => import("./components/create_model.vue"),
        },
    ]
})

createApp(App).use(router).mount('#app')
