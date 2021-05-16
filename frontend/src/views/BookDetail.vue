<template>
    <div class="m-0 p-0">
        <div class="col-11 mx-auto d-flex my-5" v-if="found">
            <div style="height: 80vh; flex-shrink: 0; max-width: 40rem; overflow-x: auto;">
                <img 
                    :src="book?.coverUrl ? `http://localhost:9000/uploads/${book?.coverUrl}` : '/assets/images/book-placeholder.jpg'" 
                    alt=""
                    style="height: 100%;"
                >
            </div>
            <div class="mx-5 d-flex" style="flex: 1; flex-direction: column;">
                <small style="align-self: flex-end;" :title="book.CreatedAt">
                    {{timeAgo(book?.CreatedAt)}}
                </small>
                <p style="font-family: Rubik; color: #000; font-size: 14px;" class="m-0">
                    ISBN: <b>{{formatISBN(book?.isbn)}}</b>
                </p>
                <h2 
                    style="font-family: Rubik; font-weight: bold; color: #000;"
                    class="m-0"
                >
                    {{book?.title}}
                </h2>
                <div class="d-flex align-items-center">
                    <Avatar :size="2" style="margin-right: 1rem;" :src="getPfp()" />
                    <div class="my-auto" style="font-family: Rubik;">
                        <router-link 
                            :to="`/author/${book.author?.username}`"
                            style="font-family: Rubik; font-weight: 900; color: #44f;"
                        >
                            {{book.author?.firstName}} {{book.author?.lastName}}
                        </router-link> 
                        <span v-if="book.publisher"> via {{book.publisher}}</span>
                    </div>
                </div>
                <p 
                    style="font-family: Rubik; color: #44f; text-transform: uppercase; font-weight: bold;"
                    class="mt-3"
                >
                    {{book.maturity}} <span class="text-dark">&bull;</span> {{book.genre}}
                </p>
                <ul style="list-style: none;">
                    <li class="text-uppercase" style="font-family: Rubik; font-size: 14px; font-weight: 900;">
                        Publish date: 
                        <span style="color: #000;">
                            {{moment(book.publishDate).format("MMMM D, YYYY")}}
                        </span>
                    </li>
                    <li 
                        class="text-uppercase" 
                        style="font-family: Rubik; font-size: 14px; font-weight: 900;"
                        v-if="book.editionNumber"
                    >
                        Edition: <span style="color: #000;">{{book.editionNumber}}</span>
                    </li>
                    <li class="text-uppercase" style="font-family: Rubik; font-size: 14px; font-weight: 900;">
                        Number of pages: <span style="color: #000;">{{book.numPages}}</span>
                    </li>
                </ul>
                <div class="my-3 d-flex align-items-baseline">
                    <p 
                        style="font-family: Rubik; font-size: 14px; font-weight: 900; margin-right: 1rem;" 
                        class="text-uppercase"
                    >
                        {{book.reads}} Reads
                    </p>
                    <div class="d-flex">
                        <button 
                            class="btn2 mx-2" 
                            style="width: 2.5rem; height: 2.5rem"
                            title="Bookmark"
                            @click="toggleBookmark"
                            :style="{backgroundColor: isLiked ? '#4f4' : '#44f'}"
                            v-if="$store.state.user.status.loggedIn"
                        >
                            <i class="fas fa-bookmark"></i>
                        </button>
                        <button 
                            class="btn2 mx-2" 
                            style="width: 2.5rem; height: 2.5rem"
                            title="Read"
                            @click="openPdf()"
                        >
                        
                            <i class="fas fa-book-open"></i>
                        </button>
                        <a 
                            class="btn2 mx-2" 
                            style="width: 2.5rem; height: 2.5rem"
                            title="Download"
                            href="http://localhost:9000/uploads/pdfs/20210515154933drylab.pdf"
                            @click.prevent="download"
                        >
                            <i class="fas fa-download text-light"></i>
                        </a>
                        <router-link class="no-decor" :to="book.ID + '/edit'">
                            <button
                                class="btn2 mx-2" 
                                style="width: 2.5rem; height: 2.5rem;"
                                title="Edit"
                                v-if="$store.state.user?.user?.ID == book.authorID"
                            >
                                <i class="fas fa-edit text-light"></i>
                            </button>
                        </router-link>
                        <button 
                            class="btn2 mx-2 bg-danger" 
                            style="width: 2.5rem; height: 2.5rem"
                            title="Delete"
                            @click="handleDeleteBook"
                            v-if="$store.state.user?.user?.ID == book.authorID"
                        >
                            <i class="fas fa-trash"></i>
                        </button>
                    </div>
                </div>
                <p style="font-size: smaller; font-weight: bold;">
                    {{book.synopsis}}
                </p>
            </div>
        </div>
        <NotFound v-if="found === false" text="Book Not Found." url="/books" />
    </div>
</template>

<script>
import Avatar from "../components/Avatar"
import axios from "axios"
import {formatISBN, getFileName, normalizeUrl, timeAgo} from "..//utils/utils"
import moment from 'moment'
import NotFound from "../components/NotFound"

export default {
    name: "BookDetail",
    components: {
        Avatar,
        NotFound
    },
    created () {
        this.getData()
    },
    data () {
        return {
            book: {},
            isLiked: false,
            found: null,
            formatISBN: formatISBN,
            timeAgo: timeAgo,
            moment: moment
        }
    },
    methods: {
        download () {
            if (this.book?.pdfUrl)
                axios.get(`http://localhost:9000/uploads/${this.book.pdfUrl}`, { responseType: 'blob' })
                    .then(response => {
                        const blob = new Blob([response.data], { type: 'application/pdf' })
                        const link = document.createElement('a')
                        link.href = URL.createObjectURL(blob)
                        link.download = getFileName(this.book.pdfUrl)
                        link.click()
                        URL.revokeObjectURL(link.href)
                    }).catch(console.error)
        },
        async getData () {
            axios.get(`http://localhost:9000/api/books/${this.$route.params.id}`)
                .then(res => {
                    this.found = true
                    this.book = res.data
                    this.isLiked = this.book.likes.map(l => l.authorID).includes(this.$store.state.user?.user?.ID)
                })
                .catch(err => {
                    console.log(err)
                    this.found = false
                })
        },
        getPfp: function () {
            return this.book?.author?.pfpUrl 
                ? normalizeUrl(`http://localhost:9000/uploads/${this.book?.author?.pfpUrl}`)
                : ""
        },
        handleDeleteBook: function () {
            if (window.confirm("Confirm the deletion of this book? It cannot be recovered.")) {
                axios.delete(`http://localhost:9000/api/books/${this.book.ID}`)
                    .then(() => {
                        this.$router.push(`/author/${this.$store.state.user.user.ID}`)
                    })
            }
        },
        toggleBookmark: function () {
            axios.put(`http://localhost:9000/api/books/${this.book.ID}/toggle-like/${this.$store.state.user.user.ID}`, {})
                .then(res => {
                    this.isLiked = res.data.likes.map(l => l.authorID).includes(this.$store.state.user.user.ID)
                })
                .catch(err => console.log(err))
        },
        openPdf () {
            if (this.book?.pdfUrl) {
                window.open(`http://localhost:9000/uploads/${this.book?.pdfUrl}`, '_blank');

                if (this.book.authorID != this.$store.state.user?.user?.ID) {
                    axios.put(`http://localhost:9000/api/books/${this.$route.params.id}/add-read`, {})
                        .then(res => {
                            this.book.plays = res.data.plays
                        })
                }
            }
        }
    }
}
</script>