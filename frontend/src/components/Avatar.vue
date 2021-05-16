<template>
    <div 
        class="avatar" 
        :style="{
            width: size + 'rem', 
            height: size + 'rem',
            backgroundImage: `url(${src || '/assets/images/avatar.png'})`,
            backgroundColor: '#fff',
            borderRadius: rounded ? '50%' : '0.5rem'
        }"
    >
    </div>
</template>

<script>
export default {
    name: "Avatar",
    props: {
        size: Number,
        src: String,
        rounded: {
            type: Boolean,
            default: true
        }
    }
}
</script>