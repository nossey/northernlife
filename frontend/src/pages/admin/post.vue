<template>
  <div>
    Title
    <input type="text" v-model="state.title">
    Body
    <!--<input type="textarea" v-model="props.body">-->
    <textarea v-model="state.body"></textarea>
    PlainBody
    <input type="text" v-model="state.plain_body">
    <Button @click.native="postman">POST</Button>

    <div v-html="state.renderedBody"></div>
  </div>
</template>

<script lang="ts">

import {defineComponent, reactive, computed} from "@nuxtjs/composition-api";
import {PostsApi, ModelPostCreateBody} from "~/client";
import {buildConfiguration} from "~/client/configurationFactory";
import Button from "~/components/atoms/Button.vue"
import marked_1 from "marked/lib/marked.esm";

type Props = {
  createPost: ModelPostCreateBody,
  isPosting: boolean
}

export default defineComponent({
  components: {
    Button
  },
  name: "post",
  middleware: 'auth',
  setup(props:Props, context) {
    props.isPosting = false
    const state = reactive({
      title: "Hello world",
      body: "# Hello World ## This is power",
      plain_body: "# Hello World ## This is power",
      renderedBody: computed(() => marked_1(state.body)),
    });

    const postman = async () => {
      if (props.isPosting)
        return;
      props.isPosting = true;
      const api = new PostsApi(buildConfiguration());
      await api.postsPost({title: state.title, body: state.body, plain_body: state.plain_body}).then(res => {
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
