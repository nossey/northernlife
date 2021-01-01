<template>
  <div>
    <b-container fluid>
      <b-row>
        <b-col class="col edit-area">
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
                <button @click="checker">Check</button>
                <b-row>
                  <b-col v-for="tag in state.tagSelections" :key="tag.name">
                    <input type="checkbox" v-bind:id="tag.name" v-bind:value="tag.name" v-model="state.selectedTags">
                    <label v-bind:for="tag.name">{{tag.name}}</label>
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
import { createMarkdown } from "safe-marked";
import Enumerable from "linq"
const markdown = createMarkdown({
  marked:{
    breaks: true
  }});
const { htmlToText } = require('html-to-text');

type Props = {
  isPosting: boolean
}

class TagSelection {
  constructor(name: string, selected: boolean) {
    this.name = name
    this.selected = selected
  }
 name: string;
 selected: boolean;
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
    let tagSelections: Array<TagSelection> = new Array<TagSelection>();
    const state = reactive({
      title: "Hello world",
      body: "# Hello World",
      renderedBody: computed(() => markdown(state.body)),
      plainBody: computed(() => htmlToText(state.renderedBody, {
        ignoreImage: true
      })),
      thumbnail: "https://northernlife-content.net/lunch.jpg",
      tagSelections: tagSelections,
      selectedTags: new Array<string>()
    });

    state.tagSelections = useAsync(async() => {
      const tagApi = new TagsApi(buildConfiguration());
      const result = (await tagApi.tagsGet()).data;
      return Enumerable.from(result).select(t => new TagSelection(t, false)).toArray();
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
      await api.postsPost({title: state.title, body: state.body, plainBody: state.plainBody, tags: [], thumbnail: state.thumbnail}).then(res => {
        // TODO:トーストとか色々出してあげる
        context.root.$router.push(`/posts/${res.data.postID}`)
      }).catch(err => {
        // TODO:トーストとか色々出してあげる
        console.log(err)
        console.log(err.response.data)
        props.isPosting = false;
      });
    }

    const checker = () => {
      console.log(state.tagSelections);
      console.log(state.selectedTags);
    }

    return {
      props,
      state,
      postman,
      checker,

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
