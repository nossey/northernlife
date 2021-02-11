<template>
  <div>
    <div v-if="fetchState.pending">Loading...</div>
    <div v-else-if="fetchState.error">Error happens!</div>
    <div v-else>
      Posts:TODO<br>
      TotalCount:{{state.totalCount}}<br>
      PerPageCount:{{state.perPageCount}}
    </div>
  </div>
</template>

<script lang="ts">

import {defineComponent, useFetch, reactive} from "@nuxtjs/composition-api";
import {PostsApi} from "~/client";
import {buildConfiguration} from "~/client/configurationFactory";

export default defineComponent({
 setup(_, context){
   const state = reactive({
    perPageCount: 0,
    totalCount:0
   });
   const {fetch, fetchState} = useFetch(async() => {
     const tag = context.root.$route.params["id"];
     const api = new PostsApi(buildConfiguration());
     const result = (await api.postsGet(1, new Array<string>(tag))).data;
     state.perPageCount = (result.perPageCount) ? result.perPageCount : 0;
     state.totalCount = (result.totalCount) ? result.totalCount : 0;
   })

   return {
     fetchState,
     state
   }
 }
})

</script>

<style scoped>

</style>
