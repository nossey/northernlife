<template>
  <div>
    <b-container fluid>
      <b-row>
        <b-col>
          <b-container class="p-3">
            <b-row>
              <b-container>
               <b-row>Title</b-row>
                <b-row>
                  <input type="text" v-model="state.title">
                </b-row>
              </b-container>
            </b-row>
            <b-row class="mt-2">
              <b-container>
                <b-row>Body</b-row>
                <b-row class="body">
                  <textarea v-model="state.body" @dragenter="dragEnter" @dragleave="dragLeave" @drop.prevent="dropFile"></textarea>
                </b-row>
              </b-container>
            </b-row>
            <b-row class="mt-2">
              Thumbnail
              :: TODO
            </b-row>
            <b-row class="mt-2">
              <Button @click.native="postman">POST</Button>
            </b-row>
          </b-container>
        </b-col>
        <b-col>
          <Post v-bind:body="state.body"></Post>
        </b-col>
      </b-row>
    </b-container>
  </div>
</template>

<script lang="ts">

import {defineComponent, reactive, computed} from "@nuxtjs/composition-api";
import {ContentsApi, PostsApi} from "~/client";
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
      body: "# Hello World",
      renderedBody: computed(() => markdown(state.body)),
      plainBody: computed(() => htmlToText(state.renderedBody, {
        ignoreImage: true
      }))
    });

    const dragEnter = () => {
      console.log("dragenter")
    }
    const dragLeave = () => {
      console.log("dragleave")
    }
    const dropFile = async (event) => {
     const file = event.dataTransfer.files[0];
     const reader = new FileReader();
     reader.onload = async () => {
       const uploader = new ContentsApi(buildConfiguration());
       const result = await uploader.contentsPost({image: reader.result as string});
       console.log(result.data.url);
     }
     reader.readAsDataURL(file);
    }

    const postman = async () => {
      if (props.isPosting)
        return;
      props.isPosting = true;
      const api = new PostsApi(buildConfiguration());
      await api.postsPost({title: state.title, body: state.body, plainBody: state.plainBody, tags: [], thumbnail: "https://northernlife-content.net/lunch.jpg"}).then(res => {
        // TODO:トーストとか色々出してあげる
        context.root.$router.push(`/posts/${res.data.postID}`)
      }).catch(err => {
        // TODO:トーストとか色々出してあげる
        console.log(err)
        console.log(err.response.data)
        props.isPosting = false;
      });
    }

    return {
      props,
      state,
      postman,

      dragEnter,
      dragLeave,
      dropFile
    }
  }
})
</script>

<style scoped lang="scss">

textarea {
  width: 100%;
  height: 300px;
  box-sizing: border-box;
}

</style>
