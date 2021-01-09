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
            :title="post.title"
            :body="post.body"
            :tags="post.tags"
            :thumbnail="post.thumbnail"
            @updated="update($event)"
          ></PostEditor>
        </div>
      </b-col>
      <b-col>
        <div v-if="fetchState.pending">
          Loading
        </div>
        <div v-else-if="fetchState.error">{{fetchState.error}}</div>
        <div v-else>
          <Post
            :title="post.title"
            :body="post.body"
            :tags="post.selectedTags"
            :thumbnail="post.thumbnail"
          ></Post>
        </div>
      </b-col>
    </b-row>
  </b-container>
</template>

<script lang="ts">
import {defineComponent, useFetch, reactive, ref, toRefs, useContext, SetupContext} from "@nuxtjs/composition-api";
import {AdminPostsApi, TagsApi} from "~/client";
import {buildConfiguration} from "~/client/configurationFactory";

import PostEditor from "~/components/molecules/PostEditor.vue"
import Post from "~/components/molecules/Post.vue"
import Button from "~/components/atoms/Button.vue";

export default defineComponent({
  components: {Button, PostEditor, Post},
  setup(){
    let post = reactive<{
      title: String,
      body: String,
      plainBody: String,
      selectedTags: Array<string>,
      tags: string[],
      thumbnail: String
    }>({
      title: "",
      body: "",
      plainBody: "",
      selectedTags: [],
      tags: [],
      thumbnail: ""
    });
    const {fetchState} = useFetch(async() => {
      const context = useContext();
      const id = context.params.value["id"];
      const api = new AdminPostsApi(buildConfiguration());
      const tagApi = new TagsApi(buildConfiguration());
      const postResult = await api.adminPostsIdGet(id);
      const tags = await tagApi.tagsGet();

      post.title = (postResult.data.title) ? postResult.data.title : "";
      post.tags = tags.data;
      post.body =  (postResult.data.body) ?  postResult.data.body : "";
      post.thumbnail = (postResult.data.thumbnail) ? postResult.data.thumbnail : "";
      post.selectedTags = (postResult.data.tags) ? postResult.data.tags : [];
    })

    const update = (event) => {
      console.log(event.title)
      post.title = event.title
      post.body = event.body
      post.thumbnail = event.thumbnail
      post.selectedTags = event.selectedTags
    };

    return {
      post,
      fetchState,
      update
    }
  }
})
</script>

<style scoped>

</style>
