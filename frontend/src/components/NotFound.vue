<template>
    <div style="margin-top: 15rem">
        <h1 class="text-center" style="color: #44f; font-weight: bold;">{{text}}</h1>
        <div class="text-center mt-3">
            <router-link class="btn4" :to="url">Back to safety</router-link>
        </div>
    </div>
</template>

<script>
export default {
    name: "NotFound",
    props: {
        text: {
            type: String,
            default: "Page not Found."
        },
        url: {
            type: String,
            default: "/"
        }
    }
}
</script>

