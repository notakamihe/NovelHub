<template>
    <div class="p-5">
        <p style="text-align: center; font-size: 2rem;" class="mb-5">Update your profile</p>
        <form class="col-10 mx-auto" @submit.prevent="handleUpdateAuthor">
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
                            type="date" 
                            name="dob" 
                            id="dob" 
                            class="col-12 input1 rounded"
                            placeholder="Date of birth"
                            v-model="dob"
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
                                :style="{backgroundImage: 'url(' + url + ')'}"
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
import axios from 'axios';
import Error from "../components/Error";
import { normalizeUrl } from '../utils/utils';
import { mapActions } from 'vuex';

export default {
    name: "EditAuthor",
    components: {
        Error
    },
    created () {
        this.getData()
    },
    data () {
        return {
            author: {},

            username: "",
            firstName: "",
            lastName: "",
            dob: "",
            email: "",
            title: "",
            about: "",
            pfp: null,

            isPfpChanged: false,
            url: "",
            
            error: ""
        }
    },
    methods: {
        ...mapActions(['updateUser']),
        getData: function () {
            axios.get(`http://localhost:9000/api/authors/username/${this.$route.params.username}`)
                .then(res => {
                    this.author = res.data

                    if (this.$store.state.user?.user?.ID != this.author.ID) {
                        window.history.back()
                        return
                    }

                    this.username = this.author.username
                    this.email = this.author.email
                    this.firstName = this.author.firstName
                    this.lastName = this.author.lastName
                    this.dob = this.author.dob
                    this.title = this.author.title
                    this.about = this.author.about

                    if (this.author.pfpUrl)
                        this.url = normalizeUrl(`http://localhost:9000/uploads/${this.author.pfpUrl}`)
                })
                .catch(() => this.$router.push("/books"))
        },
        handleChangeFile: function (e) {
            this.isPfpChanged = true
            this.pfp = e.target.files[0]
            this.url = this.pfp ? window.URL.createObjectURL(this.pfp) : ""
        },
        handleUpdateAuthor: function () {
            this.error = ""

            this.updateUser(
                {
                    user: this.author, 
                    data: {
                        username: this.username,
                        email: this.email,
                        firstName: this.firstName,
                        lastName: this.lastName,
                        dob: this.dob,
                        title: this.title,
                        about: this.about
                    },
                    changePfp: this.isPfpChanged,
                    pfp: this.pfp
                }
            )

            if (this.$store.state.user.error) {
                this.error = this.$store.state.user.error?.response?.data
            }
        }
    }
}
</script>