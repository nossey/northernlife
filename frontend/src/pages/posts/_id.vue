<template>
  <div></div>
</template>

<script lang="ts">
import {Context} from "@nuxt/types";
import { PostsApi } from "~/client";
import { buildConfiguration } from "~/client/configurationFactory"

export default {
  async asyncData(ctx: Context){
    const api = new PostsApi(buildConfiguration());
    const id = ctx.route.params['id'];
    await api.postsIdGet(id)
      .then(response => {
        console.log(response.data);
      }).catch(error => {
        ctx.error({statusCode: error.response.status});
      });
  }
}
</script>

<style scoped>

</style>
