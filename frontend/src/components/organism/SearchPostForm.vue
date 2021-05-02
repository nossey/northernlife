<template>
  <SearchForm :on-submit="submit"></SearchForm>
</template>

<script lang="ts">

import {defineComponent,useContext} from "@nuxtjs/composition-api";
import SearchForm from "~/components/atoms/SearchForm.vue";

export default defineComponent({
  components: {SearchForm},
  setup(){
    const context = useContext();
    const submit = (text: string) => {
      context.redirect(`/?search=${text}`);
    };

    return {
      submit
    }
  }
})
</script>

<style scoped>

</style>
