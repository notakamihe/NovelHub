<template>
    <div style="padding: 3rem;" class="mx-auto">
        <p style="text-align: center; font-size: 2rem;" class="mb-5">Create a book</p>
        <form class="col-11 mx-auto" @submit.prevent="handleCreateBook">
            <div class="col-12">
                <div class="col-12 my-2 d-flex">
                    <input 
                        type="text" 
                        name="title" 
                        id="title" 
                        placeholder="Title"
                        class="input1 col-12 rounded"
                        v-model="title"
                    >
                </div>
                <div class="col-12 my-2 d-flex">
                    <div style="flex: 5;" class="mx-2">
                        <input 
                            type="text" 
                            name="isbn" 
                            id="isbn" 
                            placeholder="ISBN-13" 
                            class="input1 col-12 rounded"
                            maxlength="13"
                            v-model="isbn"
                        >
                    </div>
                    <div style="flex: 4;" class="mx-2">
                            <select name="genre" id="genre" class="mx-1 input1 col-12" v-model="genre">
                                <option value="" selected disabled hidden>Genre</option>
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
                    <div style="flex: 3;" class="mx-2">
                        <input 
                            type="text" 
                            name="publishDate" 
                            id="publishDate" 
                            placeholder="Publish Date" 
                            class="input1 col-12 rounded"
                            onfocus="this.type = 'date'"
                            v-model="publishDate"
                        >
                    </div>
                </div>
                <div class="col-12 my-2 d-flex">
                    <input 
                        type="text" 
                        name="publisher" 
                        id="publisher" 
                        placeholder="Publisher"
                        class="input1 col-12 rounded"
                        v-model="publisher"
                    >
                </div>
            </div>
            <div class="mt-4 d-flex">
                <div class="mx-2 darken-on-hover" style="flex-shrink: 0; width: 15rem; height: 27rem; position: relative;">
                    <div class="rounded" style="background-color: #000; width: 100%; height: 95%;">
                        <div 
                            class="fit-bg rounded darken-on-hover-target" 
                            style="width: 100%; height: 100%; background-image: url('/assets/images/book-placeholder.jpg');"
                            :style="{backgroundImage: 'url(' + coverUrl + ')'}"
                        >

                        </div>
                    </div>
                    <p class="info1 pt-2">Portrait aspect ratios recommended</p>
                    <input 
                        type="file" 
                        class="d-none" 
                        name="cover" 
                        id="cover" 
                        accept="image/*"
                        @change="handleChangeCover"
                    >
                    <label for="cover" class="btn2 position-absolute d-none" style="top: 45%; left: 45%;">
                        <i class="fas fa-upload"></i>
                    </label>
                </div>
                <div 
                    class="mx-2 rounded darken-on-hover" 
                    style="flex-shrink: 0; width: 15rem; height: 26.666rem; background-color: #000; position: relative;"
                >
                    <div 
                        class="rounded"
                        style="background-color: #000; width: 100%; height: 100%; "
                    >
                        <iframe     
                            v-if="pdf"
                            class="fit-bg rounded darken-on-hover-target" 
                            style="width: 100%; height: 100%; border: solid #000 2px;"
                            :src="pdfUrl"
                        ></iframe>
                    </div>
                    <p 
                        class="position-absolute text-center darken-on-hover-target-remove" 
                        style="top: 45%; left: 0; right: 0; color: #fff8; text-transform: uppercase; font-family: Rubik;"
                        v-if="!pdf"
                    >
                        Choose pdf
                    </p>
                    <input 
                        type="file" 
                        class="d-none" 
                        name="pdf" 
                        id="pdf" 
                        accept="application/pdf"
                        @change="handleChangePfp"
                    >
                    <label for="pdf" class="btn2 position-absolute d-none" style="top: 45%; left: 43%;">
                        <i class="fas fa-upload"></i>
                    </label>
                </div>
                <div class="d-flex mx-2" style="flex-grow: 1">
                    <div class="col-12">
                        <textarea 
                            name="synopsis" 
                            id="synopsis" 
                            class="col-12 input1" 
                            rows="15"
                            placeholder="Synopsis (min. 15 words)"
                            v-model="synopsis"
                        ></textarea>
                        <div class="d-flex my-2">
                            <div style="flex: 1" class="mx-1">
                                <input 
                                    type="number" 
                                    name="editionNumber" 
                                    id="editionNumber"
                                    placeholder="Edition No. (opt.)"
                                    class="col-12 input1"
                                    v-model="editionNumber"
                                >
                            </div>
                            <div style="flex: 1" class="mx-1">
                                <input 
                                    type="number" 
                                    name="numPages" 
                                    id="numPages"
                                    placeholder="No. Pages"
                                    class="col-12 input1"
                                    v-model="numPages"
                                    min="1"
                                >
                            </div>
                            <select name="maturity" id="maturity" style="flex: 2" class="mx-1 input1" v-model="maturity">
                                <option value="" selected disabled hidden>Maturity</option>
                                <option value="childrens">Childrens</option>
                                <option value="teen">Teen</option>
                                <option value="young adult">Young Adult</option>
                                <option value="adult">Adult</option>
                                <option value="any">Any</option>
                            </select>
                        </div>
                    </div>
                </div>
            </div>
            <div class="d-flex mt-4" style="align-items: center;">
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

export default {
    name: "CreateBook",
    components: {
        Error
    },
    created () {
        if (!this.$store.state.user?.status.loggedIn) {
            this.$router.push("/login")
        }
    },
    data () {
        return {
            title: "",
            isbn: "",
            genre: "",
            publishDate: "",
            publisher: "",
            synopsis: "",
            editionNumber: null,
            numPages: 1,
            maturity: "",
            cover: null,
            pdf: null,

            coverUrl: "/assets/images/book-placeholder.jpg",
            pdfUrl: "",

            error: ""
        }
    },
    methods: {
        handleCreateBook: function () {
            this.error = ""

            if (!this.pdf) {
                this.error = "Must provide a PDF File."
                window.scrollTo(0, document.body.scrollHeight)
                return
            }

            axios.post("http://localhost:9000/api/books", {
                isbn: this.isbn,
                title: this.title,
                synopsis: this.synopsis,
                publishDate: this.publishDate,
                genre: this.genre,
                editionNumber: parseInt(this.editionNumber),
                numPages: parseInt(this.numPages),
                publisher: this.publisher,
                maturity: this.maturity,
                authorID: this.$store.state.user.user.ID
            })
            .then(async res => {
                let formData = new FormData()

                formData.append("cover", this.cover)
                formData.append("pdf", this.pdf)

                if (this.cover)
                    await fetch(`http://localhost:9000/api/books/${res.data.ID}/cover`, {body: formData, method: "put"})
        
                await fetch(`http://localhost:9000/api/books/${res.data.ID}/pdf`, {body: formData, method: "put"})

                this.$router.push(`/books/${res.data.ID}`)
            })
            .catch(err => {
                if (err?.response?.status == 400) {
                    this.error = err.response.data
                    window.scrollTo(0, document.body.scrollHeight)
                    return
                }
            })
        },
        handleChangeCover: function (e) {
            this.cover = e.target.files[0]
            this.coverUrl = this.cover ? window.URL.createObjectURL(this.cover) : "/assets/images/book-placeholder.jpg"
        },
        handleChangePfp: function (e) {
            this.pdf = e.target.files[0]
            this.pdfUrl = this.pdf ? window.URL.createObjectURL(this.pdf) : ''
        }
    }
}
</script>

<style lang="scss">
.darken-on-hover:hover .darken-on-hover-target {
    opacity: 0.3;
    transition: 0.25s all;
}

.darken-on-hover:hover .darken-on-hover-target-remove {
    display: none;
}

.darken-on-hover:hover > label {
    display: flex!important;
}
</style>