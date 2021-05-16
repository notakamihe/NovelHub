<template>
    <div>
        <div
            class="d-flex book rounded my-4 no-decor" 
            style="min-height: 10.666rem; border-bottom: 2px solid #44f;"
            v-if="variant == 1"
        >
            <router-link 
                class="fit-bg"
                style="width: 6rem; height: 10.666rem; border-bottom-left-radius: 0.5rem; border-top-left-radius: 0.5rem; flex-shrink: 0;"
                :style="{backgroundImage: `url('${getCoverUrl()}')`}"
                :to="`/books/${book.ID}`"
            >
            </router-link>
            <div style="flex: 1; text-align: left; flex-direction: column; min-height: 10.666rem;" class="px-3 d-flex">
                <small class="m-0" v-if="showAuthorInfo">
                    <router-link :to="`author/${book.author?.username}`">
                        <span style="color: #44f;">{{book.author?.firstName}} {{book.author?.lastName}}</span> 
                    </router-link> 
                    <span v-if="book.publisher">
                        &bull; 
                        <span style="color: #44f;">{{book.publisher}}</span>
                    </span>
                </small>
                <router-link :to="`/books/${book.ID}`" style="flex-grow: 1;">
                    <h5 style="font-weight: bold; font-family: Rubik; color: #000">{{book.title}}</h5>
                    <p style="font-size: small; font-family: Rubik; font-weight: 100;" class="text-capitalize">
                        {{book.isbn}} 
                        &bull;
                        {{new Date(book.publishDate).getUTCFullYear()}} 
                        &bull;
                        {{book.genre}}
                        <span v-if="book.editionNumber">
                            &bull;
                            Edition {{book.editionNumber}} 
                        </span>
                        &bull;
                        {{book.numPages}} Pages
                    </p>
                </router-link>
                <div class="margin-top: auto">
                    <button class="btn3 my-2" @click="showSynopsis = !showSynopsis">
                        {{showSynopsis ? "Close" : "Read"}} synopsis
                    </button>
                </div>
                <p
                    v-if="showSynopsis" 
                    class="synopsis"
                    style="font-size: smaller; color: #000; font-style: italic; line-height: normal;"
                >
                    {{book.synopsis}}
                </p>
            </div>
            <p class="p1">
                {{book.reads}} 
                <i class="fas fa-book-reader"></i>
            </p>
        </div>
        <router-link 
            class="rounded my-4 d-inline-block book2 no-decor" 
            style="width: 25rem;"
            v-if="variant == 2"
            :to="`/books/${book.ID}`"
        >
            <div class="d-flex">
                <div 
                    class="fit-bg book2-cover"
                    style="width: 6rem; height: 10.666rem; flex-shrink: 0;"
                    :style="{backgroundImage: `url('${getCoverUrl()}')`}"
                >
                </div>
                <div 
                    style="flex: 1; text-align: left; flex-direction: column; min-height: 10.666rem; border-top-right-radius: 0.75rem;" class="p-3 d-flex book-bg">
                    <small class="m-0" style="font-size: small; font-weight: bold;" v-if="showAuthorInfo">
                        <router-link :to="`/author/${book.author?.username}`">
                            <span style="color: #44f;">{{book.author.firstName}} {{book.author.lastName}}</span> 
                        </router-link> 
                        <span style="color: #44f;" v-if="book.publisher">
                            &bull; {{book.publisher}}
                        </span>
                    </small>
                    <p style="font-size: 12px; font-family: Rubik; font-weight: 100;" class="m-0">
                        {{book.isbn}} 
                    </p>
                    <h6 style="font-weight: bold; font-family: Rubik; color: #000">{{book.title}}</h6>
                    <div style="margin-left: auto; text-align: right;" class="mt-auto p-0 d-flex col-12">
                        <small style="font-family: Rubik; margin-right: auto;" class="p-0">
                            {{new Date(book.publishDate).getUTCFullYear()}}
                        </small>
                        <small style="font-family: Rubik;" class="p-0">
                            {{book.reads}} <i class="fas fa-book-reader m-0"></i>
                        </small>
                    </div>
                </div>
            </div>
            <div 
                class="mx-auto px-2 pt-2 pb-1" 
                style="background-color: #44f; border-radius: 0 0 0.75rem 0.75rem;">
                <p style="font-size: small; font-family: Rubik; font-weight: 100; text-transform: capitalize; color: #fff;">
                    {{book.maturity}}
                    &bull;
                    {{book.genre}}
                    <span v-if="book.editionNumber" class="text-light">
                        &bull; Edition
                        {{book.editionNumber || null}}
                    </span>
                    &bull;
                    {{book.numPages}}p
                </p>
                <p class="p1 scrollbar1" style="line-height: 1.15rem; max-height: 9.2rem; overflow-y: scroll; color: #fff;; font-weight: bold;">
                    {{book.synopsis}}
                </p>
            </div>
        </router-link>
    </div>
</template>

<script>
import { normalizeUrl } from '../utils/utils'
export default {
    name: "Book", 
    props: {
        variant: {
            type: Number,
            default: 1
        },
        book: Object,
        showAuthorInfo: {
            type: Boolean,
            default: true
        }
    },
    data () {
        return {
            showSynopsis: false
        }
    },
    methods: {
        getCoverUrl: function () {
            return this.book?.coverUrl 
                ? normalizeUrl(`http://localhost:9000/uploads/${this.book?.coverUrl}` )
                : '/assets/images/book-placeholder.jpg'
        }
    }
}
</script>

<style lang="scss" scoped>
.book:hover {
    background-color: #0001;
    cursor: pointer;
    box-shadow: 0.15rem 0.15rem 1rem 0.1rem #000;
    padding: 1rem;
    transition: all 0.25s;
}

.book:not(:hover) {
    transition: all 0.5s;
}

.book-bg {
    border-width: 2px 2px 0 0;
    border-color: #44f;
    border-style: solid solid none none;
}

.book2:hover {
    box-shadow: 0 0 2rem 0.1rem #000;
    background-color: #0000;
    cursor: pointer;
    transition: all 1.5s;
    transform: translateY(-1rem);
}

.book2:hover .book2-cover {
    box-shadow: 0.25rem 0 0.5rem 0.05rem #000a;
    transition: 0.25s all;
}
</style>