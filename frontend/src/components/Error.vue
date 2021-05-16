<template>
    <div v-if="msg">
        <p>{{ msg }}</p>
    </div>
</template>

<script>
export default {
    name: "Error",
    props: {
        msg: String
    }
}
</script>

<style lang="scss" scoped>
p {
    background-color: #f003; 
    color: #f00; 
    font-weight: 500; 
    text-transform: uppercase; 
    font-family: Rubik;
    padding: 0.5rem 1rem;
    margin: 0.4rem;
    border-radius: 0.25rem;
}
</style>