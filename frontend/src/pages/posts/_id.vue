<template>
  <b-container>
    <transition name="post">
      <b-row align-h="between">
        <b-col cols="12" md="9" class="post">
          <Post
            :title="postResult.title"
            :body="postResult.body"
            :postedAt="postResult.createdAt"
            :thumbnail="postResult.thumbnail"
            :tags="postResult.tags"
            :linkList="links">
          </Post>
          {{toc}}
        </b-col>
        <b-col md="3" class="d-none d-md-block">
          <!--<div class="side-area" v-if="state.toc.length > 0">
          <div class="toc">
            <div class="title">Table of contents</div>
            <a v-for="content in state.toc" :href="`${content.link}`" class="pl-3">{{content.name}}</a>
          </div>
            <Profile class="mt-2"></Profile>
          </div>-->
        </b-col>
      </b-row>
    </transition>
  </b-container>
</template>

<script>
import { PostsApi } from "~/client";
import { buildConfiguration } from "~/client/configurationFactory"
import { defineComponent, reactive, computed, useFetch, useAsync, useContext, useMeta } from "@nuxtjs/composition-api"
import Post from "~/components/molecules/Post.vue"
import Profile from "~/components/atoms/Profile.vue"
import Enumerable from "linq"
import moment from "moment";
import {markdown} from "~/application/posts/markdown";

export default ({
  components: {
   Post, Profile
  },
  head: {},
  async asyncData(ctx){
    const api = new PostsApi(buildConfiguration());
    const postResult = (await api.postsIdGet(ctx.params["id"])).data;
    return {
      postResult,
    }
  },
  computed: {
    links: function(){
      return {links: Enumerable.from(this.postResult.tags).select(function(t){return {name: t, link: `/?tag=${encodeURIComponent(t)}`}}).toArray()};
    },
    toc: function(){
      return markdown(this.postResult.body)[1]
    }
  }
  //setup(){
  //  const state = reactive({
  //    thumbnail: "",
  //    title: "Title",
  //    body: "",
  //    plainBody: "",
  //    tags: [],
  //    postedAt: "",
  //    tagLinks: computed(() => {
  //      const links = Enumerable.from(state.tags as string[]).select(function(t){return {name: t, link: `/?tag=${encodeURIComponent(t)}`}}).toArray()
  //      return {links: links}
  //    }),
  //    toc: computed(() => {
  //      return markdown(state.body)[1]
  //    })
  //  });

  //  const { title, meta } = useMeta();
  //  const postResult = useAsync(async() => {
  //    const api = new PostsApi(buildConfiguration());
  //    const id = useContext().params.value["id"];
  //    return (await api.postsIdGet(id)).data;
  //  });
  //  title.value = postResult.value?.title as string;
  //  meta.value = [
  //    {
  //      hid: 'og:description',
  //      property: 'og:description',
  //      content: postResult.value?.plainBody as string,
  //    },
  //    {
  //      hid: 'og:image',
  //      property: 'og:image',
  //      content: postResult.value?.thumbnail as string,
  //    },
  //  ]
  //  console.log(process.server)
  //  console.log(postResult.value)

  //  return {
  //    state,
  //    postResult: postResult.value
  //  }
  //}
})
</script>

<style lang="scss" scoped>
@import "~assets/colors";

.post-enter-active, .post-leave-active { transition: opacity .5s; }
.post-enter, .post-leave-active { opacity: 0; }

.post {
  padding: 0 0;
}

.side-area {
  position: sticky;
  top: 10px;
  .toc {
    .title {
      text-align: center;
      font-weight: bold;
    }
    border-radius: 2px;
    filter: drop-shadow(2px 2px 3px $shadow-color);
    background: $background-white;
    min-height: 300px;
    a {
      display: block;
      word-break: break-all;
    }
  }
}

</style>
