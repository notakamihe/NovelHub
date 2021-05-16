<template>
    <div style="width: 100%; height: 64vh; flex-direction: column;" class="d-flex align-items-center justify-content-center">
        <p>Presenting...</p>
        <h1 style="font-size: 5rem; font-weight: black;">NovelHub</h1>
        <h5 style="font-family: Rubik; font-weight: 100;" class="mt-3 text-muted">
            Your second-favorite hub ;)
        </h5>
        <router-link 
            class="btn2 mt-5" 
            style="font-weight: 900; font-weight: Rubik;"
            to="/books"
        >
            GO
        </router-link>
    </div>
</template>

<script>
export default {
    name: "Home"
}
</script>