<template>
  <b-container>
    <transition name="post">
      <b-row align-h="between">
        <b-col cols="12" md="9" class="post">
          <div v-if="fetchState.error">{{fetchState.error.message}}</div>
          <Post v-else-if="!fetchState.pending"
                :title="state.title"
                :body="state.body"
                :thumbnail="state.thumbnail"
                :tags="state.tags"
                :linkList="state.tagLinks"
                :posted-at="state.postedAt"
          ></Post>
        </b-col>
        <b-col md="2" class="d-none d-md-block ">
          <div class="toc-container" v-if="!fetchState.error && !fetchState.pending && state.toc.length > 0">
            <div v-if="fetchState.error">{{fetchState.error.message}}</div>
            <div v-else-if="!fetchState.pending" class="pt-2">
              <a class="toc ml-3" v-for="content in state.toc" :href="`${content.link}`">{{content.name}}</a>
            </div>
          </div>
        </b-col>
      </b-row>
    </transition>
  </b-container>
</template>

<script lang="ts">
import { PostsApi } from "~/client";
import { buildConfiguration } from "~/client/configurationFactory"
import { defineComponent, reactive, computed, useFetch, useContext, useMeta } from "@nuxtjs/composition-api"
import Post from "~/components/molecules/Post.vue"
import Enumerable from "linq"
import moment from "moment";
import {markdown} from "~/application/posts/markdown";

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
      postedAt: "",
      tagLinks: computed(() => {
        const links = Enumerable.from(state.tags as string[]).select(function(t){return {name: t, link: `/?tag=${encodeURIComponent(t)}`}}).toArray()
        return {links: links}
      }),
      toc: computed(() => {
        return markdown(state.body)[1]
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
      state.postedAt = moment(post.createdAt).format("YYYY-MM-DD");

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
@import "~assets/colors";

.post-enter-active, .post-leave-active { transition: opacity .5s; }
.post-enter, .post-leave-active { opacity: 0; }

.post {
  padding: 0 0;
}

.toc-container{
  background: $background-white;
  border-radius: 2px;
  filter: drop-shadow(2px 2px 3px $shadow-color);
  min-height: 300px;
  .toc {
    display: block;
  }

  position: sticky;
  top: 10px;
}

</style>
