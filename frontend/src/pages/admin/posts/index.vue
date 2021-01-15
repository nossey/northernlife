<template>
  <b-container>
    <nuxt-link to="/admin">Admin Top</nuxt-link>
    <b-row>
      <b-col>
        <Button to="/admin/post">Create New Post</Button>
        <b-tabs>
          <b-tab title="All">
            <b-container>
              <b-row v-if="fetchState.pending">Loading...</b-row>
              <b-row v-else v-for="post in allPosts.posts" :key="post.Id">
                <nuxt-link :to="`/admin/posts/${post.Id}`">{{post.Title}}</nuxt-link>
                <Button @click.native="deletePost(post.Id)">Delete</Button>
              </b-row>
            </b-container>
          </b-tab>
          <b-tab title="Published">
            <b-container>
              <b-row v-for="post in publishedPosts.posts" :key="post.Id">
                <nuxt-link :to="`/admin/posts/${post.Id}`">{{post.Title}}</nuxt-link>
                <Button @click.native="deletePost(post.Id)">Delete</Button>
              </b-row>
            </b-container>
          </b-tab>
          <b-tab title="Draft">
            <b-container v-if="draftPosts.posts.length > 0">
              <b-row v-for="post in draftPosts.posts" :key="post.Id">
                <nuxt-link :to="`/admin/posts/${post.Id}`">{{post.Title}}</nuxt-link>
                <Button @click.native="deletePost(post.Id)">Delete</Button>
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
import Button from "~/components/atoms/Button.vue"
import {defineComponent, reactive, useFetch} from "@nuxtjs/composition-api";

interface PostListItem {
  Id: string;
  Title: string;
}

export default defineComponent({
  middleware: ['auth'],
  components: {Button},
  setup() {
    const allPosts = reactive<{posts: PostListItem[]}>({posts: new Array<PostListItem>()}) ;
    const draftPosts = reactive<{posts: Array<PostListItem>}>({posts: new Array<PostListItem>()}) ;
    const publishedPosts = reactive<{posts: Array<PostListItem>}>({posts: new Array<PostListItem>()}) ;

    const {fetch, fetchState} = useFetch(async() => {
      const api = new AdminPostsApi(buildConfiguration());
      await api.adminPostsGet(1, "all")
        .then((res) => {
          allPosts.posts = new Array<PostListItem>();
          const result = res.data;
          if (result.posts)
          {
            result.posts.forEach(p => {
                if (p && p.id && p.title){
                  allPosts.posts.push({Id: p.id, Title: p.title})
                }
              }
            )
          }
        }).catch((err) => {
          console.log(err);
        });

      const draftPostsResult = await api.adminPostsGet(1, "draft");
      const publishedResult = await api.adminPostsGet(1, "published");

      draftPosts.posts = new Array<PostListItem>();
      publishedPosts.posts = new Array<PostListItem>();

      if (draftPostsResult.data.posts)
      {
        draftPostsResult.data.posts.forEach(p => {
            if (p && p.id && p.title){
              draftPosts.posts.push({Id: p.id, Title: p.title})
            }
          }
        )
      }
      if (publishedResult.data.posts)
      {
        publishedResult.data.posts.forEach(p => {
            if (p && p.id && p.title){
              publishedPosts.posts.push({Id: p.id, Title: p.title})
            }
          }
        )
      }
    });
    return {
      fetchState,
      fetch,
      allPosts,
      publishedPosts,
      draftPosts
    }
  },
  methods: {
    async deletePost(id) {
      const api = new AdminPostsApi(buildConfiguration());
      await api.adminPostsIdDelete(id).then((result) => {
        this.fetch()
      }).catch((err) => {
        // TODO:
      })
    }
  }
})

</script>

<style scoped>

</style>
