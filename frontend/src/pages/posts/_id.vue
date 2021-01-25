<template>
  <b-container class="mt-4">
    <b-row>
      <b-col cols="12" md="9">
        <Post
          :title="title"
          :body="body"
          :thumbnail="thumbnail"
          :tags="tags"
          :linkList="tagLinks"
        ></Post>
      </b-col>
      <b-col md="3" class="d-none d-md-block">Table of contents</b-col>
    </b-row>
  </b-container>
</template>

<script>
import { PostsApi } from "~/client";
import { buildConfiguration } from "~/client/configurationFactory"
import Post  from "~/components/molecules/Post.vue"
import Enumerable from "linq"

export default {
  components: {
   Post
  },
  head(){
    return {
      title: this.title,
      meta: [
        // OGP
        { hid: 'og:title', property: 'og:title', content: this.title },
        { hid: 'og:description', property: 'og:description', content: this.plainBody },
        { hid: 'og:image', property: 'og:image', content: this.thumbnail },
      ]
    }
  },
  async asyncData(ctx){
    const api = new PostsApi(buildConfiguration());
    const id = ctx.route.params['id'];
    try
    {
      const post = (await api.postsIdGet(id)).data
      const links = Enumerable.from(post.tags).select(function(t) { return {name: t, link: `/tags/${t}`}}).toArray()
      return {
        thumbnail: post.thumbnail,
        title: post.title,
        body: post.body,
        tags: post.tags,
        tagLinks: {links: links}
      }
    }
    catch (error)
    {
      ctx.error({statusCode: error.response.status});
      return {
        thumbnail: "",
        title: "",
        body: "",
        tags: []
      }
    }
  },
}
</script>

<style scoped>

</style>
