<template>
  <b-container>
    <b-row>
      <b-col>
        <button @click="getPost">GetPost</button>
      </b-col>
    </b-row>
  </b-container>
</template>

<script lang="ts">
import Vue from "vue";
import { Context } from "@nuxt/types";
import { PostsApi } from "~/client";
import { buildConfiguration } from "~/client/configurationFactory"

export default Vue.extend({
  methods: {
    async getPost(ctx: Context){
      const post = new PostsApi(buildConfiguration());
      const result = await post.postsGet(1)
      console.log(result.data);
   }},
  async asyncData(ctx: Context): Promise<void> {
    const post = new PostsApi(buildConfiguration());
    const result = await post.postsGet(1);
    console.log(result.data);
  },
});
</script>

<style>
</style>
