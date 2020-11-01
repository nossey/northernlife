<template>
  <b-container>
    <b-row>
      <b-col>
        <ul v-if="result.posts.length > 0">
          <li v-for="post in result.posts"><nuxt-link :to="`/posts/${post.id}`">{{post.id}}</nuxt-link></li>
        </ul>
      </b-col>
    </b-row>
  </b-container>
</template>

<script lang="ts">
import { Context } from "@nuxt/types";
import { PostsApi } from "~/client";
import { buildConfiguration } from "~/client/configurationFactory"
import { defineComponent } from "@nuxtjs/composition-api"

export default defineComponent({
  async asyncData(ctx: Context): Promise<Object> {
    const post = new PostsApi(buildConfiguration());
    const response = await post.postsGet(1);
    return {result: response.data}
  },
});
</script>

<style>
</style>
