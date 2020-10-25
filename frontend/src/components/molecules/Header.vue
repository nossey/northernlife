<template>
  <div class="header"><nuxt-link to="/">Northernlife</nuxt-link></div>
</template>

<script>
export default {
name: "Header"
}
</script>

<style lang="scss" scoped>
@import "assets/colors.scss";

.header {
  background-color: $main-theme;
  height: 40px;

  a {
    color: $font-color-white;
    text-decoration: none;
  }
}

</style>
