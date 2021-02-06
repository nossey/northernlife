<template>
  <form @submit.prevent="submit">
    <fa :icon="faSearch"></fa>
    <input type="search" placeholder="検索">
  </form>
</template>

<script lang="ts">

import {defineComponent, computed} from "@nuxtjs/composition-api";

import {faSearch} from "@fortawesome/free-solid-svg-icons";

export default defineComponent({
 setup(){
   const submit = () => {
     console.log("Hello")
   }

   return {
     submit,
     faSearch
   }
 }

})

</script>

<style lang="scss" scoped>
@import "assets/colors";

form {
  display: inline-block;
  border-radius: 4px;
  background: $background-white;
  svg {
    color: $shadow-color;
    margin-left: 10px;
  }
  input {
    border: none;
    padding: 3px;
    outline: none;
  }
}

</style>
