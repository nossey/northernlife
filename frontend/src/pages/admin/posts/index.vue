<template>
  <b-container>
    <nuxt-link to="/admin">Admin Top</nuxt-link>
    <b-row>
      <b-col>
        <b-tabs>
          <b-tab title="All">
            <b-container>
              <b-row v-for="post in allPosts.posts" :key="post.id">
                <nuxt-link :to="`/admin/posts/${post.id}`">{{post.title}}</nuxt-link>
              </b-row>
            </b-container>
          </b-tab>
          <b-tab title="Published">
            <b-container>
              <b-row v-for="post in publishedPosts.posts" :key="post.id">
                <nuxt-link :to="`/admin/posts/${post.id}`">{{post.title}}</nuxt-link>
              </b-row>
            </b-container>
          </b-tab>
          <b-tab title="Draft">
            <b-container v-if="draftPosts.posts.length > 0">
              <b-row v-for="post in draftPosts.posts" :key="post.id">
                <nuxt-link :to="`/admin/posts/${post.id}`">{{post.title}}</nuxt-link>
              </b-row>
            </b-container>
            <b-container v-else>No Draft Posts</b-container>
          </b-tab>
        </b-tabs>
      </b-col>
    </b-row>
  </b-container>
</template>

<script lang="ts">

import { buildConfiguration } from "~/client/configurationFactory"
import {AdminPostsApi} from "~/client";

export default {
  async asyncData(){
    const api = new AdminPostsApi(buildConfiguration());
    const allPosts = await api.adminPostsGet(1, "all");
    const draftPosts = await api.adminPostsGet(1, "draft");
    const published = await api.adminPostsGet(1, "published");
    return {
      allPosts: allPosts.data,
      draftPosts: draftPosts.data,
      publishedPosts: published.data
    }
  },
}

</script>

<style scoped>

</style>
