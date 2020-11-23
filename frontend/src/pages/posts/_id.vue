<template>
  <div>
    <Post :body="this.post.body"></Post>
  </div>
</template>

<script lang="ts">
import { Context } from "@nuxt/types";
import { PostsApi } from "~/client";
import { buildConfiguration } from "~/client/configurationFactory"
import Post from "~/components/molecules/Post.vue"

export default {
  components: {
   Post
  },
  async asyncData(ctx: Context){
    const api = new PostsApi(buildConfiguration());
    const id = ctx.route.params['id'];
    try
    {
      const post = await api.postsIdGet(id)
      return {
        post: post.data
      }
    }
    catch (error)
    {
      ctx.error({statusCode: error.response.status});
    }
  }
}
</script>

<style scoped>

</style>
