<template>
  <b-container class="mt-4">
    <b-row>
      <b-col cols="12" md="9">
        <transition name="post">
          <div v-if="fetchState.error">{{fetchState.error.message}}</div>
          <Post v-else-if="!fetchState.pending"
                :title="state.title"
                :body="state.body"
                :thumbnail="state.thumbnail"
                :tags="state.tags"
                :linkList="state.tagLinks"
          ></Post>
        </transition>
      </b-col>
      <b-col md="3" class="d-none d-md-block">toc</b-col>
    </b-row>
  </b-container>
</template>

<script lang="ts">
import { PostsApi } from "~/client";
import { buildConfiguration } from "~/client/configurationFactory"
import { defineComponent, reactive, computed, useFetch, useContext, useMeta } from "@nuxtjs/composition-api"
import Post from "~/components/molecules/Post.vue"
import Enumerable from "linq"
const toc = Array<{level: Number, slug: string, title: string}>();

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
        const links = Enumerable.from(state.tags as string[]).select(function(t){return {name: t, link: `/?tag=${encodeURIComponent(t)}`}}).toArray()
        return {links: links}
      })
    });

    const { title, meta } = useMeta();
    const {fetch, fetchState} = useFetch(async() => {
      const api = new PostsApi(buildConfiguration());
      const id = useContext().params.value["id"];
      const post = (await api.postsIdGet(id)).data;
      state.thumbnail = post.thumbnail;
      state.title = post.title;
      state.body = post.body;
      state.plainBody = post.plainBody;
      state.tags = (post.tags) ? post.tags : [];

      title.value = state.title;
      meta.value = [
        {
          hid: 'og:description',
          property: 'og:description',
          content: state.plainBody,
        },
        {
          hid: 'og:image',
          property: 'og:image',
          content: state.thumbnail,
        },
      ]
    });

    return {
      state,
      fetchState
    }
  }
})
</script>

<style lang="scss" scoped>
.post-enter-active, .post-leave-active { transition: opacity .5s; }
.post-enter, .post-leave-active { opacity: 0; }
</style>
