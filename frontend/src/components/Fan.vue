<template>
    <router-link 
        class="d-flex user my-5" 
        style="height: 10rem;"
        :to="`/author/${fan.username}`"
    >
        <Avatar :size="10" :rounded="false" style="flex-shrink: 0;" class="user-img" />
        <div style="text-align: left; height: 100%" class="mx-3">
            <div>
                <div class="my-0 d-flex">
                    <h4 style="font-family: Rubik; font-weight: bold; color: #44f;">
                        {{fan.firstName}} {{fan.lastName}}
                    </h4>
                </div>
                <h5 style="font-family: Rubik; font-weight: bold;" class="text-muted my-0">
                    {{fan.title}}
                </h5>
                <p>@{{fan.username}}</p>
            </div>
            
            <p class="p1 user-about" style="line-height: 1.15rem; max-height: 4.5rem; overflow-y: hidden;">
                {{fan.about}}
            </p>
        </div>
    </router-link>
</template>

<script>
import Avatar from "./Avatar"

export default {
    name: "Fan",
    components: {
        Avatar
    },
    props: {
        fan: Object
    }
}
</script>

<style lang="scss" scoped>
.user:hover {
    cursor: pointer;
    box-shadow: 0.15rem 0.15rem 1rem 0.1rem #000;
    padding: 1rem;
    transition: all 0.25s;
    height: auto!important;
}

.user:hover .user-img {
    box-shadow: 0.15rem 0.15rem 1rem 0.1rem #000;
    margin-right: 1rem;
}

.user:hover .user-about {
    max-height: initial!important;
    transition: all 0.5s;
}

.user:not(:hover) {
    transition: all 0.5s;
}
</style>