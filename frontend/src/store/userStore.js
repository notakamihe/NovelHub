import axios from "axios"
import router from "../router/index"

const currentUser = JSON.parse(localStorage.getItem("user"))

const state = currentUser ? {status: {loggedIn: true}, user: currentUser} : {status: {}, user: null}

const actions = {
    deleteUser ({commit}, id) {
        axios.delete(`http://localhost:9000/api/authors/${id}`)
            .then(() => {
                commit("logout")
                router.push("/")
            })
            .catch(err => console.log(err))
    },
    login ({commit}, {username, password}) {
        axios.post("http://localhost:9000/api/login", {username, password})
            .then(res => {
                axios.get("http://localhost:9000/api/author", {headers: {"x-access-token": res.data}})
                    .then(r => {
                        commit("loginSuccess", r.data)
                        router.push(`/author/${r.data.username}`)
                    })
                    .catch(e => {
                        commit("loginFailure", e)
                    })
            })
            .catch(err => {
                commit("loginFailure", err)
            })
    },
    
    logout ({commit}) {
        commit("logout")
        router.push("/login")
    },

    register ({commit}, {user, pfp}) {
        axios.post("http://localhost:9000/api/register", user)
            .then(res => {
                axios.get("http://localhost:9000/api/author", {headers: {"x-access-token": res.data}})
                    .then(async r => {
                        let formData = new FormData()

                        formData.append("pfp", pfp)

                        fetch(`http://localhost:9000/api/authors/${r.data.ID}/pfp`, {
                            body: formData, 
                            method: "put"
                        })
                        .then(response => response.json())
                        .then(data => {
                            commit("registerSuccess", data)
                            router.push(`/author/${data.username}`)
                        })
                    })
                    .catch(e => {
                        commit("registerFailure", e)
                    })
            })
            .catch(err => {
                commit("registerFailure", err)
            })
    },

    updateUser ({commit}, {user, data, changePfp, pfp}) {
        axios.put(`http://localhost:9000/api/authors/${user.ID}`, {
                username: data.username,
                email: data.email,
                firstName: data.firstName,
                lastName: data.lastName,
                dob: data.dob,
                title: data.title,
                about: data.about
            })
            .then(async res => {
                let formData = new FormData()

                formData.append("pfp", pfp)

                if (changePfp) {
                    const data = await (await fetch(`http://localhost:9000/api/authors/${user.ID}/pfp`, {body: formData, method: "put"})).json()

                    res.data = data
                }

                commit("updateUserSuccess", res.data)

                router.push(`/author/${res.data.username}`)
            })
            .catch(err => {
                commit("updateUserFailure", err)
            })
    }
}

const mutations = {
    loginFailure (state, err) {
        state.status = {};
        state.user = null;
        state.error = err
    },
    loginSuccess (state, user) {
        localStorage.setItem('user', JSON.stringify(user));
        state.status = { loggedIn: true };
        state.user = user;
    },
    logout () {
        localStorage.removeItem("user")
        state.status = {};
        state.user = null;
    },
    registerFailure (state, err) {
        state.status = {};
        state.user = null;
        state.error = err
    },
    registerSuccess (state, user) {
        localStorage.setItem('user', JSON.stringify(user));
        state.status = { loggedIn: true };
        state.user = user;
    },
    updateUserSuccess (state, user) {
        localStorage.setItem("user", JSON.stringify(user))
        state.status = {loggedIn: true}
        state.user = user
    },
    updateUserFailure (state, err) {
        state.error = err
    }
}

const getters = {
    getUser (state) {
        return state.user
    }
}

export const user = {
    state,
    actions,
    mutations,
    getters
}