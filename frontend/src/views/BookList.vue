<template>
    <div>
        <form class="col-10 mx-auto" @submit.prevent="filterBooks" id="form">
            <div class="col-12 d-flex">
                <input 
                    type="search" 
                    name="search" 
                    id="search" 
                    style="flex: 1;" 
                    class="input1 rounded"
                    placeholder="Search by title, ISBN, publisher or author"
                >
                <button type="submit" class="btn2" style="margin-left: 1rem;">
                    <i class="fas fa-search"></i>
                </button>
            </div>
            <div>
                <div class="d-flex justify-content-center my-3" style="align-items: baseline;">
                    <p style="font-weight: bold;">Filter</p> 
                    <button class="mx-3 btn2">
                        <i class="fas fa-filter"></i>
                    </button>
                </div>
                <div class="d-flex align-items-center justify-content-center">
                    <div style="flex: 1;" class="mx-2">
                        <select name="genre" id="genre" class="input1 col-12">
                            <option value="" selected disabled hidden>Genre</option>
                            <option value="">None</option>
                            <option value="fiction">Fiction</option>
                            <option value="narrative">Narrative</option>
                            <option value="science fiction">Science Fiction</option>
                            <option value="non-fiction">Non-fiction</option>
                            <option value="mystery">Mystery</option>
                            <option value="historcial fiction">Historical fiction</option>
                            <option value="biography">Biography</option>
                            <option value="horror">Horror</option>
                            <option value="poem">Poem</option>
                            <option value="humor">Humor</option>
                            <option value="short">Short</option>
                            <option value="autobiography">Autobiography</option>
                            <option value="romance">Romance</option>
                            <option value="crime">Crime</option>
                            <option value="fairy tale/folklore">Fairy tale/Folklore</option>
                            <option value="moral">Moral</option>
                            <option value="fantasy">Fantasy</option>
                            <option value="comic">Comic</option>
                            <option value="essay">Essay</option>
                            <option value="dystopian">Dystopian</option>
                            <option value="drama">Drama</option>
                            <option value="satire">Satire</option>
                            <option value="memoir">Memoir</option>
                            <option value="suspense">Suspense</option>
                            <option value="other">Other</option>
                        </select>
                    </div>
                    <div style="flex: 1;" class="mx-2 d-flex">
                        <div style="flex: 1" class="mx-1">
                            <input 
                                type="text" 
                                name="fromDate" 
                                id="fromDate" 
                                placeholder="Start Date"
                                class="input1 rounded"
                                onfocus="this.type = 'date'"
                            >
                        </div>
                        <div style="flex: 1" class="mx-1">
                            <input 
                                type="text" 
                                name="toDate" 
                                id="toDate" 
                                placeholder="End Date"
                                class="input1 rounded"
                                onfocus="this.type = 'date'"
                            >
                        </div>
                    </div>
                    <div style="flex: 1;" class="mx-2 d-flex">
                        <select 
                            name="comparison" 
                            id="comparison" 
                            style="flex: 2;"
                            class="col-12 input1"
                        >
                            <option value="equals">&equals;</option>
                            <option value="lt">&lt;</option>
                            <option value="lte">&le;</option>
                            <option value="gt">&gt;</option>
                            <option value="gte">&ge;</option>
                        </select>
                        <input 
                            type="number" 
                            name="numPages" 
                            id="numPages" 
                            placeholder="No. Pages"
                            class="col-12 input1" 
                            style="flex: 6;"
                        >
                    </div>
                    <div style="flex: 1;" class="mx-2">
                        <select name="maturity" id="maturity" class="input1 col-12">
                            <option value="" selected disabled hidden>Maturity</option>
                            <option value="">None</option>
                            <option value="childrens">Childrens</option>
                            <option value="teen">Teen</option>
                            <option value="young adult">Young Adult</option>
                            <option value="adult">Adult</option>
                            <option value="any">Any</option>
                        </select>
                    </div>
                </div>
            </div>
        </form>
        <div class="mt-5 col-12 text-center">
            <div class="d-inline-block mx-2" :key="book" v-for="book in books">
                <Book :variant="2" :book="book" />
            </div>
        </div>
    </div>
</template>

<script>
import Book from "../components/Book"
import axios from "axios"

export default {
    name: "BookList",
    components: {
        Book
    },
    created () {
        axios.get("http://localhost:9000/api/books")
            .then(res => {
                this.books = res.data
            })
            .catch(err => console.log(err))
    },
    data () {
        return {
            books: []
        }
    },
    methods: {
        filterBooks: function () {
            const form = document.getElementById("form")
            const params = new URLSearchParams(new FormData(form)).toString()

            axios.get("http://localhost:9000/api/books?" + params)
                .then(res => {
                    this.books = res.data
                })
        }
    }
}
</script>