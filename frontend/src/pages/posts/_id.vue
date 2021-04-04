<template>
  <b-container class="mt-4">
    <b-row>
      <b-col cols="12" md="9">
        <div v-if="fetchState.pending">Loading...</div>
        <div v-else-if="fetchState.error">{{fetchState.error.message}}</div>
        <Post v-else
          :title="state.title"
          :body="state.body"
          :thumbnail="state.thumbnail"
          :tags="state.tags"
          :linkList="state.tagLinks"
        ></Post>
      </b-col>
      <b-col md="3" class="d-none d-md-block">Table of contents</b-col>
    </b-row>
  </b-container>
</template>

<script>
import { PostsApi } from "~/client";
import { buildConfiguration } from "~/client/configurationFactory"
import { defineComponent, reactive, computed, useFetch, useContext, useMeta } from "@nuxtjs/composition-api"
import Post from "~/components/molecules/Post.vue"
import Enumerable from "linq"

export default defineComponent({
  components: {
   Post
  },
  head: {},
  setup(){
    const state = reactive({
      thumbnail: "",
      title: "",
      body: "",
      plainBody: "",
      tags: [],
      tagLinks: computed(()=>{
        const links = Enumerable.from(state.tags).select(function(t){return {name: t, link: `/?tag=${encodeURIComponent(t)}`}}).toArray()
        return {links: links}
      })
    });

    const { title } = useMeta();
    const {_, fetchState} = useFetch(async() => {
      const api = new PostsApi(buildConfiguration());
      const id = useContext().params.value["id"];
      const post = (await api.postsIdGet(id)).data;
      state.thumbnail = post.thumbnail;
      state.title = post.title;
      state.body = post.body;
      state.plainBody = post.plainBody;
      state.tags = (post.tags) ? post.tags : [];

      title.value = state.title;
    });

    return {
      state,
      fetchState
    }
  }
})
</script>

<style scoped>

</style>
