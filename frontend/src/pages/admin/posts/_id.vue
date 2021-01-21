<template>
  <b-container>
    <b-row>
      <b-col>
        <div v-if="fetchState.pending">
          Loading
        </div>
        <div v-else-if="fetchState.error">{{fetchState.error}}</div>
        <div v-else>
          <PostEditor
            :title="state.title"
            :body="state.body"
            :tags="state.tags"
            :thumbnail="state.thumbnail"
            @updated="updated($event)"
          ></PostEditor>
        </div>
        <Button @click.native="putPost(true)">Save</Button>
        <Button @click.native="putPost(false)">Save as draft</Button>
        <Button @click.native="deletePost">Delete</Button>
      </b-col>
      <b-col>
        <div v-if="fetchState.pending">
          Loading
        </div>
        <div v-else-if="fetchState.error">{{fetchState.error}}</div>
        <div v-else>
          <Post
            v-bind:title="state.title"
            v-bind:body="state.body"
            v-bind:tags="state.selectedTags"
            v-bind:thumbnail="state.thumbnail"
          ></Post>
        </div>
      </b-col>
    </b-row>
  </b-container>
</template>

<script lang="ts">
import {defineComponent, useFetch, reactive, useContext, computed} from "@nuxtjs/composition-api";
import {AdminPostsApi, AdminTagsApi} from "~/client";
import {buildConfiguration} from "~/client/configurationFactory";
import { createMarkdown } from "safe-marked";
const markdown = createMarkdown({
  marked:{
    breaks: true
  }});
const { htmlToText } = require('html-to-text');

import PostEditor from "~/components/molecules/PostEditor.vue"
import Post from "~/components/molecules/Post.vue"
import Button from "~/components/atoms/Button.vue";

export default defineComponent({
  middleware: ['auth'],
  components: {Button, PostEditor, Post},
  setup(props, context){
    const state = reactive({
      title: "",
      body: "",
      selectedTags: new Array<string>(),
      tags: new Array<string>(),
      thumbnail: "",
      renderedBody: computed(() => markdown(state.body)),
      plainBody: computed(() => htmlToText(state.renderedBody, {
        ignoreImage: true
      })),
    });
    const ctx = useContext();
    const id = ctx.params.value["id"];
    const {fetchState} = useFetch(async() => {
      const api = new AdminPostsApi(buildConfiguration());
      const tagApi = new AdminTagsApi(buildConfiguration());
      const postResult = await api.adminPostsIdGet(id);
      const tags = await tagApi.adminTagsGet()

      state.title = (postResult.data.title) ? postResult.data.title : "";
      state.tags = tags.data.tags;
      state.body =  (postResult.data.body) ?  postResult.data.body : "";
      state.thumbnail = (postResult.data.thumbnail) ? postResult.data.thumbnail : "";
      state.selectedTags = (postResult.data.tags) ? postResult.data.tags : new Array<string>();
    })

    const updated = (event) => {
      state.title = event.title
      state.body = event.body
      state.thumbnail = event.thumbnail
      state.selectedTags = event.selectedTags
    };

    const putPost = async (publish: boolean) => {
      const api = new AdminPostsApi(buildConfiguration());
      await api.adminPostsIdPut(id, {
        title: state.title,
        body: state.body,
        plainBody: state.plainBody,
        tags: state.selectedTags,
        thumbnail: state.thumbnail,
        published: publish})
      .then(res => {
        // TODO:トーストとか色々出してあげる
        context.root.$router.push(`/posts/${res.data.postID}`)
      }).catch(err => {
        // TODO:トーストとか色々出してあげる
        console.log(err)
        console.log(err.response.data)
      });
    }

    const deletePost = async () => {
      const api = new AdminPostsApi(buildConfiguration());
      await api.adminPostsIdDelete(id)
        .then(res => {
          // TODO:トーストとか色々出してあげる
          context.root.$router.push(`/admin/posts/`)
        }).catch(err => {
          // TODO:トーストとか色々出してあげる
          console.log(err)
          console.log(err.response.data)
        });
    }

    return {
      state,
      fetchState,
      updated,
      putPost,
      deletePost
    }
  }
})
</script>

<style scoped>

</style>
