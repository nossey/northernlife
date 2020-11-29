<template>
  <div>
    Title
    <input type="text" v-model="state.title">
    Body
    <textarea v-model="state.body"></textarea>
    PlainBody
    <Button @click.native="postman">POST</Button>

    <Post v-bind:body="state.body"></Post>
  </div>
</template>

<script lang="ts">

import {defineComponent, reactive, computed} from "@nuxtjs/composition-api";
import {PostsApi} from "~/client";
import {buildConfiguration} from "~/client/configurationFactory";
import Button from "~/components/atoms/Button.vue"
import Post from "~/components/molecules/Post.vue"
import { createMarkdown } from "safe-marked";
const markdown = createMarkdown();
const { htmlToText } = require('html-to-text');

type Props = {
  isPosting: boolean
}

export default defineComponent({
  components: {
    Button,
    Post
  },
  name: "post",
  middleware: 'auth',
  setup(props:Props, context) {
    props.isPosting = false
    const state = reactive({
      title: "Hello world",
      body: "# Hello World ## This is power",
      renderedBody: computed(() => markdown(state.body)),
      plainBody: computed(() => htmlToText(state.renderedBody))
    });

    const postman = async () => {
      if (props.isPosting)
        return;
      props.isPosting = true;
      const api = new PostsApi(buildConfiguration());
      await api.postsPost({title: state.title, body: state.body, plain_body: state.plainBody, tags: []}).then(res => {
        // TODO:トーストとか色々出してあげる
        context.root.$router.push(`/posts/${res.data.post_id}`)
      }).catch(err => {
        // TODO:トーストとか色々出してあげる
        console.log(err)
        props.isPosting = false;
      });
    }

    return {
      props,
      state,
      postman
    }
  }
})
</script>

<style scoped>

</style>
