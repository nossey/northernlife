<template>
  <div>
    <Post :body="this.post.body"></Post>
  </div>
</template>

<script>
import { PostsApi } from "~/client";
import { buildConfiguration } from "~/client/configurationFactory"
import Post from "~/components/molecules/Post.vue"

export default {
  components: {
   Post
  },
  head(){
    return {
      title: this.title,
      meta: [
        // OGP
        { hid: 'og:title', property: 'og:title', content: this.post.title },
        { hid: 'og:description', property: 'og:description', content: this.post.plainBody },
        { hid: 'og:image', property: 'og:image', content: this.post.thumbnail },
      ]
    }
  },
  async asyncData(ctx){
    const api = new PostsApi(buildConfiguration());
    const id = ctx.route.params['id'];
    try
    {
      const post = await api.postsIdGet(id)
      return {
        post: post.data,
        title: post.data.title
      }
    }
    catch (error)
    {
      ctx.error({statusCode: error.response.status});
      return {
        post: null
      }
    }
  },
}
</script>

<style scoped>

</style>
