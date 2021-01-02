<template>
  <b-container>
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
  </b-container>
</template>

<script lang="ts">

import {defineComponent, reactive, watch} from "@nuxtjs/composition-api";
import {ContentsApi} from "~/client";
import {buildConfiguration} from "~/client/configurationFactory";

export default defineComponent( {
  props: {
    title: {
      type: String,
      required: true
    },
    body: {
      type: String,
      required: true
    },
    tags: {
      type: Array,
      required: true
    },
    thumbnail: {
      type: String,
      required:true
    }
  },
  setup(props, context){

    const state = reactive({
      title: props.title,
      body: props.body,
      tags: props.tags,
      thumbnail: props.thumbnail,
      selectedTags: new Array<string>(),
    });

    watch(() => state, () => {context.emit("updated", state)}, {deep: true})

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
        state.thumbnail = (result.data.url) ? result.data.url : "";
      }
      reader.readAsDataURL(file);
    }

    return {
      state,
      dropFile,
      uploadThumbnail
    }
  }
})
</script>

<style lang="scss" scoped>

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
