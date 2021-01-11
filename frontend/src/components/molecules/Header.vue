<template>
  <b-container fluid class="header">
    <b-row class="text-center h-100">
      <b-col v-if="loggedIn"></b-col>
      <b-col class="center" cols="8">
        <nuxt-link to="/">NorthernLife</nuxt-link>
      </b-col>
      <b-col v-if="loggedIn">
        <nuxt-link to="/admin">Admin</nuxt-link>
      </b-col>
    </b-row>
  </b-container>
</template>

<script lang="ts">
import {defineComponent, ref} from "@nuxtjs/composition-api";

export default defineComponent({
  middleware: 'auth',
  setup(_props, context) {
    const root = context.root as any;
    const loggedIn = ref(root.$auth.loggedIn);

    return {
      loggedIn
    }
  }
})
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
