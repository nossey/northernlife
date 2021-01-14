<template>
  <b-container fluid="">
   <b-row>
     <b-col><nuxt-link to="/admin/posts">Posts</nuxt-link></b-col>
   </b-row>
  </b-container>
</template>

<script lang="ts">

import {defineComponent} from "@nuxtjs/composition-api";
import Button from "~/components/atoms/Button.vue"

export default defineComponent({
  middleware: ['auth'],
  components: {Button},
})

</script>

<style scoped>

</style>
