<template>
  <b-container fluid="">
   <b-row>
     <b-col><nuxt-link to="/admin/posts">Posts</nuxt-link></b-col>
     <b-col><nuxt-link to="/admin/tags">Tags</nuxt-link></b-col>
   </b-row>
  </b-container>
</template>

<script lang="ts">

import {defineComponent} from "@nuxtjs/composition-api";

export default defineComponent({
  middleware: ['auth'],
})

</script>

<style scoped>

</style>
