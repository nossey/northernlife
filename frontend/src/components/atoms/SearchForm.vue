<template>
  <form @submit.prevent="submit">
    <input type="search" placeholder="検索">
  </form>
</template>

<script lang="ts">

import {defineComponent} from "@nuxtjs/composition-api";

export default defineComponent({
 setup(){
   const submit = () => {
     console.log("Hello")
   }

   return {
     submit
   }
 }

})

</script>

<style lang="scss" scoped>

form {
  input {
    border: none;
    padding: 3px;
  }
}

</style>
