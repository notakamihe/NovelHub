<template>
    <div class="p-5">
        <p style="text-align: center; font-size: 2rem;" class="mb-5">Register</p>
        <form class="col-10 mx-auto" @submit.prevent="handleRegister">
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
                    <div style="flex: 3;" class="mx-2">
                        <input 
                            type="text" 
                            name="firstName" 
                            id="firstName" 
                            class="col-12 input1 rounded"
                            placeholder="First Name"
                            v-model="firstName"
                        >
                    </div>
                    <div style="flex: 3;" class="mx-2">
                        <input 
                            type="text" 
                            name="lastName" 
                            id="lastName" 
                            class="col-12 input1 rounded"
                            placeholder="Last Name"
                            v-model="lastName"
                        >
                    </div>
                </div>
                <div class="d-flex col-12 my-4">
                    <div style="flex: 2;" class="mx-2">
                        <input 
                            type="email" 
                            name="email" 
                            id="email" 
                            class="col-12 input1 rounded"
                            placeholder="Email"
                            v-model="email"
                        >
                    </div>
                    <div style="flex: 1;" class="mx-2">
                        <input 
                            type="text" 
                            name="dob" 
                            id="dob" 
                            class="col-12 input1 rounded"
                            placeholder="Date of birth"
                            onfocus="this.type = 'date'"
                            v-model="dob"
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
                    <div style="flex: 1;" class="mx-2">
                        <input 
                            type="password" 
                            name="confirmPassword" 
                            id="confirmPassword" 
                            class="col-12 input1 rounded"
                            placeholder="Confirm Password"
                            v-model="confirmPassword"
                        >
                    </div>
                </div>
                <div class="col-12 d-flex px-2">
                    <input 
                        type="text" 
                        name="title" 
                        id="title" 
                        class="input1 rounded col-12"
                        placeholder="Title"
                        v-model="title"
                    >
                </div>
                <div class="d-flex col-12 my-4">
                    <div style="flex: 1;" class="mx-2">
                        <textarea 
                            name="about" 
                            id="about" 
                            class="col-12 input1" 
                            rows="5"
                            placeholder="About"
                            v-model="about"
                        ></textarea>
                    </div>
                    <div class="d-flex align-items-center justify-content-center mx-2">
                        <div>
                            <input 
                                type="file" 
                                name="pfp" 
                                id="pfp" 
                                class="d-none" 
                                @change="handleChangeFile"
                                accept="image/*"
                            >
                            <div 
                                class="rounded position-relative fit-bg"
                                style="width: 9rem; height: 9rem; background-color: #000"
                                :style="{backgroundImage: 'url(' + pfpUrl() + ')'}"
                            >
                                <label 
                                    for="pfp" 
                                    class="position-absolute text-center"
                                    style="top: 30%; left: 0rem; right: 0rem; color: #fffa; height: 100%"
                                >
                                    Choose profile image
                                </label>
                            </div>
                        </div>
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
import { mapActions } from 'vuex';
import Error from "../components/Error";

export default {
    name: "Register",
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
            firstName: "",
            lastName: "",
            dob: "",
            email: "",
            password: "",
            confirmPassword: "",
            title: "",
            about: "",
            pfp: null,
            
            error: ""
        }
    },
    methods: {
        ...mapActions(["register"]),
        handleRegister: function () {
            this.error = ""

            if (this.password != this.confirmPassword) {
                window.scrollTo(0, document.body.scrollHeight)
                this.error = "Passwords must match."
                return
            }

            this.register(
            {
                user: {
                    username: this.username,
                    firstName: this.firstName,
                    lastName: this.lastName,
                    dob: this.dob,
                    email: this.email,
                    password: this.password,
                    title: this.title,
                    about: this.about
                },
                pfp: this.pfp
            })

            if (this.$store.state.user.error) {
                this.error = this.$store.state.user.error.response.data
            }
        },
        handleChangeFile: function (e) {
            this.pfp = e.target.files[0]
        },
        pfpUrl: function () {
            return this.pfp ? window.URL.createObjectURL(this.pfp) : ""
        }
    }
}
</script>