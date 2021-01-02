<template>
  <div>
    <b-container fluid>
      <b-row>
        <b-col class="col edit-area">
          <PostEditor
            :title="state.title"
            :body="state.body"
            :tags="state.tags"
            :thumbnail="state.thumbnail"
            @updated="updated($event)"
          ></PostEditor>
          <b-container class="p-3">
            <b-row>
              <b-container>
               <b-row>Title</b-row>
                <b-row>
                  <input type="text" v-model="state.title">
                </b-row>
              </b-container>
            </b-row>
            <b-row>
              <b-container>
                <b-row>
                  <b-col v-for="(tag, index) in state.tags" :key="index">
                    <input type="checkbox" v-bind:id="tag" v-bind:value="tag" v-model="state.selectedTags">
                    <label v-bind:for="tag">{{tag}}</label>
                  </b-col>
                </b-row>
              </b-container>
            </b-row>
            <b-row class="mt-2">
              <b-container>
                <b-row>Body</b-row>
                <b-row class="body">
                  <textarea v-model="state.body" @drop.prevent="dropFile"></textarea>
                </b-row>
              </b-container>
            </b-row>
            <b-row class="mt-2">
              <b-container>
                <b-row>Thumbnail:</b-row>
                <b-row>{{state.thumbnail}}</b-row>
                <b-row>
                  <input @change="uploadThumbnail" type="file">
                </b-row>
              </b-container>
            </b-row>
            <b-row class="mt-2">
              <Button @click.native="postman">POST</Button>
            </b-row>
          </b-container>
        </b-col>
        <b-col class="col">
          <Post
            v-bind:thumbnail="state.thumbnail"
            v-bind:title="state.title"
            v-bind:body="state.body"
            v-bind:tags="state.selectedTags"
          ></Post>
        </b-col>
      </b-row>
    </b-container>
  </div>
</template>

<script lang="ts">

import {defineComponent, reactive, computed, useAsync, useFetch} from "@nuxtjs/composition-api";
import {ContentsApi, PostsApi, TagsApi} from "~/client";
import {buildConfiguration} from "~/client/configurationFactory";
import Button from "~/components/atoms/Button.vue"
import Post from "~/components/molecules/Post.vue"
import PostEditor from "~/components/molecules/PostEditor.vue"
import { createMarkdown } from "safe-marked";
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
  middleware: 'auth',
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
      selectedTags: new Array<string>()
    });

    state.tags = useAsync(async() => {
      const tagApi = new TagsApi(buildConfiguration());
      return (await tagApi.tagsGet()).data;
    });

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

    const uploadThumbnail = (event) => {
      const file = event.target.files[0];
      const reader = new FileReader();
      reader.onload = async () => {
        const uploader = new ContentsApi(buildConfiguration());
        const result = await uploader.contentsPost({image: reader.result as string});
        state.thumbnail = result.data.url;
      }
      reader.readAsDataURL(file);
    }

    const postman = async () => {
      if (props.isPosting)
        return;
      props.isPosting = true;
      const api = new PostsApi(buildConfiguration());
      await api.postsPost({title: state.title, body: state.body, plainBody: state.plainBody, tags: state.selectedTags, thumbnail: state.thumbnail}).then(res => {
        // TODO:トーストとか色々出してあげる
        context.root.$router.push(`/posts/${res.data.postID}`)
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

      dropFile,
      uploadThumbnail
    }
  }
})
</script>

<style scoped lang="scss">
@import "assets/colors.scss";

textarea {
  width: 100%;
  height: 300px;
  box-sizing: border-box;
}

.edit-area {
  border-right: 1px dotted $shadow-color;
}

</style>
