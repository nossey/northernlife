<template>
  <b-container fluid class="header">
    <b-row class="text-center h-100">
      <b-col class="center">
        <nuxt-link to="/">NorthernLife</nuxt-link>
      </b-col>
    </b-row>
  </b-container>
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
  height: 60px;
  .center {
    margin: auto;
  }
  a {
    color: $font-color-white;
    text-decoration: none;
    font-size: 24px;
  }
}

</style>
