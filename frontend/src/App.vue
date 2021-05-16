<template>
    <div class="m-0 p-0">
        <div style="margin-bottom: 7rem;">
            <Header />
        </div>
        <router-view :key="$route.fullPath"></router-view>
        <div style="margin-top: 7rem;">
            <Footer />
        </div>
    </div>
</template>

<script>
import Header from "./components/Header"
import Footer from "./components/Footer"

export default {
  name: 'App',
  components: {
      Header,
      Footer
  }
}
</script>

<style lang="scss">
* {
    font-family: Exo;
}

.no-decor {
    text-decoration: none;
    color: initial;
}

.no-decor * {
    text-decoration: none;
    color: initial;
}

.no-outline {
    outline: none;
}

.no-btn {
    border: none;
    background-color: #0000;
}

h1 {
    font-family: Exo;
    font-weight: 900;
}

.link1 {
    @extend .no-decor;
    text-transform: uppercase;
    color: #0009;
    font-family: Rubik;
}

.input1 {
    @extend .no-outline;
    background-color: #0002;
    border: none;
    padding: 0.5rem;
    font-family: Rubik;
}

.btn1 {
    background-color: #44f;
    color: white;
    border: none;
}

.fit-bg {
    background-size: cover;
    background-position: center;
    background-repeat: no-repeat;
}

.info1 {
    font-family: Rubik;
    font-size: 12px;
    text-transform: uppercase;
    line-height: 1;
    color: #000c;
}

.btn2 {
    background-color: #44f;
    border-radius: 50%;
    color: #fff;
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 0.75rem;
    border: none;
    cursor: pointer;
    transition: 0.2s background-color;
}

.avatar {
    @extend .fit-bg;
    border-radius: 50%;
}

.p1 {
    font-family: Rubik;
    line-height: normal;
    font-size: 13px;
    color: #000;
    letter-spacing: 1px;
    font-weight: 100;
}

.btn3 {
    @extend .no-btn;
    color: #0009;
    font-weight: bolder;
    font-size: 13px;
}

.scrollbar1::-webkit-scrollbar {
    height: 0.2rem;
    width: 0.2rem;
}

.scrollbar1::-webkit-scrollbar-thumb {
    background-color: #fff;
    border-radius: 2rem
}

.btn4 {
    @extend .no-decor;
    @extend .no-btn;
    font-family: Rubik;
    padding: 0.35rem;
    font-size: 14px;
    text-transform: uppercase;
    transition: all 0.25s;
}

.btn4:hover {
    color: #fff;
    background-color: #44f;
    transition: all 0.25s;
    border-radius: 0.5rem;
    font-weight: 900;
}

</style>
