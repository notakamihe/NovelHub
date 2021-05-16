<template>
    <div>
        <div class="pb-5" v-if="found">
            <div class="d-flex align-items-center justify-content-center" style="margin-top: 10rem;">
                <div class="position-relative">
                    <Avatar :size="12" :src="pfpUrl" />
                    <button 
                        class="btn2 mx-auto position-absolute" 
                        style="width: 3rem; height: 3rem; bottom: 0; left: 0;"
                        :style="{backgroundColor: isAFan ? '#4f4' : '#44f'}"
                        @click="toggleFan"
                        v-if="$store.state.user.status.loggedIn"
                    >
                        <i class="fas fa-user-plus"></i>
                    </button>
                </div>
                <div class="mx-3">
                    <h2 style="font-family: Rubik; font-weight: bold; color: #44f;">{{user.firstName}} {{user.lastName}}</h2>
                    <h4 style="font-family: Rubik; font-weight: bold;" class="text-muted">@{{user.username}}</h4>
                    <h6 style="font-weight: 600;">{{user.email}}</h6>
                </div>
            </div>
            <div>
                <div class="mx-auto text-center mt-5 col-9">
                    <MDBTabs v-model="activeTabId1">
                        <MDBTabNav tabsClasses="mb-3" class="mx-auto text-center">
                            <div class="mx-auto d-flex">
                                <MDBTabItem tabId="about">About</MDBTabItem>
                                <MDBTabItem tabId="books">{{books?.length}} Books</MDBTabItem>
                                <MDBTabItem tabId="fans">{{user?.fans?.length}} Fans</MDBTabItem>
                                <MDBTabItem tabId="bookmarked">{{liked?.length}} Bookmarked</MDBTabItem>
                                <MDBTabItem tabId="account" v-if="$store.state.user?.user?.ID == user.ID">
                                    Account
                                </MDBTabItem>
                            </div>
                        </MDBTabNav>
                        <MDBTabContent>
                            <MDBTabPane tabId="about">
                                <p style="font-family: Rubik; font-size: 13px; font-weight: 100;">
                                    Born: 
                                    <span style="font-weight: black;">
                                        {{moment(user.dob).format("MMMM DD, YYYY")}} ({{calculateYears(user.dob)}})
                                    </span>
                                </p>
                                <h3 style="font-family: Rubik; font-weight: bold;">
                                    {{user.title}}
                                </h3>
                                <p class="p1 mt-3">
                                    {{user.about}}
                                </p>
                            </MDBTabPane>
                            <MDBTabPane tabId="books">
                                <div :key="book" v-for="book in books">
                                    <Book :showAuthorInfo="false" :book="book" />
                                </div>
                            </MDBTabPane>
                            <MDBTabPane tabId="fans">
                                <div :key="fan" v-for="fan in fans">
                                    <Fan :fan="fan" />
                                </div>
                            </MDBTabPane>
                            <MDBTabPane tabId="account">
                                <div class="mt-5">
                                    <div class="my-3">
                                        <button class="btn4" type="submit" @click="handleLogout">Log out</button>
                                    </div>
                                    <div class="my-3">
                                        <router-link class="btn4" :to="`${user?.username}/edit`">Edit</router-link>
                                    </div>
                                    <div class="my-3" @click="handleDeleteAuthor">
                                        <button class="btn4">Delete</button>
                                    </div>
                                </div>
                            </MDBTabPane>
                            <MDBTabPane tabId="bookmarked">
                                <div :key="like" v-for="like in liked">
                                    <Book :book="like.book" />
                                </div>
                            </MDBTabPane>
                        </MDBTabContent>
                    </MDBTabs>
                </div>
            </div>
        </div>
        <NotFound v-if="found === false" text="Author Not Found." url="/" />
    </div>
</template>

<script>
import { MDBTabs, MDBTabNav, MDBTabContent, MDBTabItem, MDBTabPane } from 'mdb-vue-ui-kit';
import { ref } from 'vue';

import Avatar from "../components/Avatar"
import Book from "../components/Book"
import Fan from "../components/Fan"

import moment from "moment"
import { mapActions } from 'vuex';
import axios from 'axios';
import { normalizeUrl } from '../utils/utils';
import NotFound from '../components/NotFound';

export default {
    name: "AuthorDetail",
    components: {
        Avatar,
        Book,
        Fan,
        MDBTabs,
        MDBTabNav,
        MDBTabContent,
        MDBTabItem,
        MDBTabPane,
        NotFound
    },
    created () {
        this.getData()
    },
    data () {
        return {
            books: [],
            fans: [],
            liked: [],
            user: {},
            found: null,
            isAFan: false,
            pfpUrl: "",
            moment: moment
        }
    },
    methods: {
        ...mapActions(['logout', "deleteUser"]),
        calculateYears: (date) => {
            var ageDifMs = Date.now() - new Date(date);
            var ageDate = new Date(ageDifMs);
            return Math.abs(ageDate.getUTCFullYear() - 1970);
        },
        getData: function () {
            axios.get(`http://localhost:9000/api/authors/username/${this.$route.params.username}`)
            .then(res => {
                this.found = true
                this.user = res.data

                axios.get(`http://localhost:9000/api/books/author/${this.user?.ID}`)
                    .then(res => {
                        this.books = res.data
                        this.books.reverse()
                    })
                    .catch(err => console.log(err))

                axios.get(`http://localhost:9000/api/authors/${this.user?.ID}/fans`)
                    .then(res => {
                        this.fans = res.data
                        this.isAFan = this.fans.map(f => f.ID).includes(this.$store.state.user?.user?.ID)
                    })
                    .catch(err => console.log(err))

                axios.get(`http://localhost:9000/api/books/liked/${this.user?.ID}`)
                    .then(res => {
                        this.liked = res.data
                        this.liked.reverse()
                    })
                    .catch(err => console.log(err))

                if (this.user?.pfpUrl)
                    this.pfpUrl = normalizeUrl(`http://localhost:9000/uploads/${this.user.pfpUrl}`)
            })
            .catch(err => {
                console.log(err);
                this.found = false
            })
        },
        handleDeleteAuthor: function () {
            if (window.confirm("Are you sure you want to delete your profile. All associated books and likes will be removed as well and the profile will not be recoverable.")) {
                if (window.confirm("Last chance. Confirm?")) {
                    this.deleteUser(this.user.ID)
                }
            }
        },
        handleLogout: function () {
            this.logout({})
        },
        toggleFan: function () {
            axios.put(`http://localhost:9000/api/authors/${this.user.ID}/toggle-fan/${this.$store.state.user?.user?.ID}`, {})
                .then(() => {
                    this.getData()
                })
                .catch(err => console.log(err))
        }
    },
    setup() {
        const activeTabId1 = ref('about');

        return { activeTabId1 };
    },
}
</script>