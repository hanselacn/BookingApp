import { createRouter, createWebHistory } from "vue-router";

const routes =[
    {
        path: '/bookings',
        name: 'bookings.index',
        component: ()=> import("../views/bookings/index.vue")

    },
    {
        path: '/login',
        name: 'login.login',
        component: ()=> import("../views/login/login.vue")

    },
    {
        path: '/logout',
        name: 'login.logout',
        component: ()=> import("../views/login/logout.vue")

    },
    {
        path: '/register',
        name: 'homepage.register',
        component: ()=> import("../views/homepage/register.vue")

    },
    {
        path: '/bookings/roomday',
        name: 'bookings.bookbyroomday',
        component: ()=> import("../views/bookings/bookbyroomday.vue")

    },
    {
        path: '/mybooking',
        name: 'bookings.mybooking',
        component: ()=> import("../views/bookings/mybooking.vue")

    },
    {
        path: '/edit',
        name: 'users.edit',
        component: ()=> import("../views/users/edit.vue")

    },
    {
        path: '/room',
        name: 'room.room',
        component: ()=> import("../views/room/room.vue")

    }
]

const router = createRouter({
    history : createWebHistory(),
    routes
})

export default router