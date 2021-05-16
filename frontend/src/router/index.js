import {createRouter, createWebHistory} from "vue-router"
import {AuthorDetail, BookDetail, BookList, CreateBook, EditAuthor, EditBook, Home, Login, Register} from "../views"
import NotFound from "../components/NotFound"

const routes = [
    {
        path: "/",
        name: "Home",
        component: Home
    },
    {
        path: "/register",
        name: "Register",
        component: Register
    },
    {
        path: "/login",
        name: "Login",
        component: Login
    },
    {
        path: "/books",
        name: "BookList",
        component: BookList
    },
    {
        path: "/books/create",
        name: "CreateBook",
        component: CreateBook
    },
    {
        path: "/books/:id",
        name: "BookDetail",
        component: BookDetail
    },
    {
        path: "/books/:id/edit",
        name: "EditBook",
        component: EditBook
    },
    {
        path: "/author/:username",
        name: "AuthorDetail",
        component: AuthorDetail
    },
    {
        path: "/author/:username/edit",
        name: "EditAuthor",
        component: EditAuthor
    },
    {
        path: "/:pathMatch(.*)*",
        component: NotFound
    }
]

const router = createRouter({
    history: createWebHistory(process.env.BASE_URL),
    routes
})

export default router