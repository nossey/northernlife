<template>
  <div>
    <PostCard
      class="margin"
      v-for="post in result.posts"
      :id="post.id"
      :title="post.title"
      :body="post.body"
      :key="post.id"></PostCard>
  </div>
</template>

<script lang="ts">
import { Context } from "@nuxt/types";
import { PostsApi, TagsApi } from "~/client";
import { buildConfiguration } from "~/client/configurationFactory"
import { defineComponent } from "@nuxtjs/composition-api"
import Button from "~/components/atoms/Button.vue"
import PostCard from "~/components/molecules/PostCard.vue"

export default defineComponent({
  async asyncData(ctx: Context): Promise<Object> {
    const post = new PostsApi(buildConfiguration());
    const tagApi = new TagsApi(buildConfiguration());
    const response = await post.postsGet(1);
    const tagsResponse = await tagApi.tagsGet();
    console.log(tagsResponse.data)
    return {
      result: response.data,
      tags: tagsResponse.data
    }
  },
  components: {
    Button,
    PostCard
  }
});
</script>

<style lang="scss" scoped>

.margin {
  margin: 10px 5px;
}

</style>
