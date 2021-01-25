<template>
  <div>
    <b-container fluid>
      <b-row v-if="fetchState.pending">Loading...</b-row>
      <b-row v-else-if="fetchState.error">Error!</b-row>
      <b-row v-else>
        <b-col class="col edit-area">
          <PostEditor
            :title="state.title"
            :body="state.body"
            :tags="state.tags"
            :thumbnail="state.thumbnail"
            :selectedTags="state.selectedTags"
            @updated="updated($event)"
          ></PostEditor>
          <b-container class="p-3">
            <b-row class="mt-2">
              <Button @click.native="postman(true)">Save</Button>
              <Button @click.native="postman(false)">Save as draft</Button>
            </b-row>
          </b-container>
        </b-col>
        <b-col class="col">
          <Post
            v-bind:thumbnail="state.thumbnail"
            v-bind:title="state.title"
            v-bind:body="state.body"
            v-bind:tags="state.selectedTags"
            v-bind:linkList="state.linkList"
          ></Post>
        </b-col>
      </b-row>
    </b-container>
  </div>
</template>

<script lang="ts">

import {defineComponent, reactive, computed, useAsync, useFetch} from "@nuxtjs/composition-api";
import {AdminPostsApi, AdminTagsApi} from "~/client";
import {buildConfiguration} from "~/client/configurationFactory";
import Button from "~/components/atoms/Button.vue"
import Post from "~/components/molecules/Post.vue"
import PostEditor from "~/components/molecules/PostEditor.vue"
import { createMarkdown } from "safe-marked";
import Enumerable from "linq";
const markdown = createMarkdown({
  marked:{
    breaks: true
  }});
const { htmlToText } = require('html-to-text');

type Props = {
  isPosting: boolean
}

export default defineComponent({
  components: {
    PostEditor,
    Button,
    Post
  },
  name: "post",
  middleware: ['auth'],
  setup(props:Props, context) {
    props.isPosting = false
    const state = reactive({
      title: "Hello world",
      body: "# Hello World",
      renderedBody: computed(() => markdown(state.body)),
      plainBody: computed(() => htmlToText(state.renderedBody, {
        ignoreImage: true
      })),
      thumbnail: "https://northernlife-content.net/lunch.jpg",
      tags: new Array<string>(),
      selectedTags: new Array<string>(),
      linkList: computed(() => {
        const links = Enumerable.from(state.selectedTags).select(function(t) {return {name: t, link: `/tags/${t}`}}).toArray();
        return {links: links}
      })
    });

    const {fetch: fetchTags, fetchState} = useFetch(async() => {
      const api = new AdminTagsApi(buildConfiguration());
      state.tags = (await api.adminTagsGet()).data.tags;
    });

    const postman = async (publish: boolean) => {
      if (props.isPosting)
        return;
      props.isPosting = true;
      const api = new AdminPostsApi(buildConfiguration());
      await api.adminPostsPost({title: state.title, body: state.body, plainBody: state.plainBody, tags: state.selectedTags, thumbnail: state.thumbnail, published: publish}).then(res => {
        // TODO:トーストとか色々出してあげる
        if (publish)
          context.root.$router.push(`/posts/${res.data.postID}`)
        else
          context.root.$router.push(`/admin/posts/${res.data.postID}`)
      }).catch(err => {
        // TODO:トーストとか色々出してあげる
        console.log(err)
        console.log(err.response.data)
        props.isPosting = false;
      });
    }

    const updated = (event) => {
      state.title = event.title
      state.body = event.body
      state.thumbnail = event.thumbnail
      state.selectedTags = event.selectedTags
    };

    return {
      props,
      state,
      postman,
      updated,
      fetchState
    }
  }
})
</script>

<style scoped lang="scss">
@import "assets/colors.scss";

.edit-area {
  border-right: 1px dotted $shadow-color;
}

</style>
