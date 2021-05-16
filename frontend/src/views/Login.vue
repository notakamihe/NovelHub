<template>
    <div class="p-5">
        <p style="text-align: center; font-size: 2rem;" class="mb-5">Login</p>
        <form class="col-10 mx-auto" @submit.prevent="handleLogin">
            <div>
                <div class="d-flex col-12 my-4">
                    <div style="flex: 4;" class="mx-2">
                        <input 
                            type="text" 
                            name="username" 
                            id="username" 
                            class="col-12 input1 rounded"
                            placeholder="Username"
                            v-model="username"
                        >
                    </div>
                </div>
                <div class="d-flex col-12 my-4">
                    <div style="flex: 1;" class="mx-2">
                        <input 
                            type="password" 
                            name="password" 
                            id="password" 
                            class="col-12 input1 rounded"
                            placeholder="Password"
                            v-model="password"
                        >
                    </div>
                </div>
            </div>
            <div class="d-flex" style="align-items: center;">
                <div style="flex: 1">
                    <input 
                        type="submit" 
                        class="btn1 p-2 rounded mx-2" style="font-size: larger; font-weight: bold;"
                    >
                </div>
                <Error :msg="error" />
            </div>
        </form>
    </div>
</template>

<script>
import Error from "../components/Error";
import { mapActions } from "vuex";

export default {
    name: "Login", 
    components: {
        Error
    },
    created () {
        if (this.$store.state.user?.status.loggedIn) {
            this.$router.push(`/author/${this.$store.state.user?.user.username}`)
        }
    },
    data () {
        return {
            username: "",
            password: "",
            
            error: ""
        }
    },
    methods: {
        ...mapActions(['login']),
        handleLogin: function () {
            this.login({username: this.username, password: this.password})

            if (this.$store.state.user.error) {
                this.error = this.$store.state.user.error.response.data
            }
        }
    }
}
</script>