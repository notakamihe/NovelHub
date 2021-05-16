<template>
    <div class="d-flex py-3 px-5 bg-light fixed-top">
        <h1 style="flex-grow: 1;">
            <router-link to="/" class="no-decor" style="color: #44f">
                Novel <span class="hub">Hub</span>
            </router-link>
        </h1>
        <div class="d-flex align-items-center">
            <div v-if="!isLoggedIn" class="d-flex align-items-center">
                <router-link to="/register" class="mx-2 link1">Register</router-link>
                <router-link to="/login" class="mx-2 link1">Login</router-link>
            </div>
            <div v-if="isLoggedIn" class="d-flex align-items-center">
                <router-link to="/books" class="mx-2 link1 ">Explore</router-link>
                <router-link to="/books/create" class="mx-2 link1">Create</router-link>
                <router-link :to="`/author/${$store.state.user?.user?.username}`" class="mx-2">
                    <Avatar :size="3" :src="getUrl()" />
                </router-link>
            </div>
        </div>
    </div>
</template>

<script>
import { normalizeUrl } from '../utils/utils'
import Avatar from "./Avatar"

export default {
    name: "Header",
    components: {
        Avatar
    },
    data () {
        return {
            isLoggedIn: this.$store.state.user.status.loggedIn
        }
    },
    methods: {
        getUrl: function () {
            const url = this.$store.state.user?.user?.pfpUrl
            return url ? normalizeUrl(`http://localhost:9000/uploads/${url}`) : ""
        }
    },
    watch:{
        $route (){
            this.isLoggedIn = this.$store.state.user.status.loggedIn
        }
    } 
}
</script>

<style lang="scss" scoped>
.hub {
    color: white;
    background-color: #44f;
    font-weight: 900;
    padding: 0.5rem;
    border-radius: 0.5rem;
}
</style>