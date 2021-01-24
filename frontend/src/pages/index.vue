<template>
  <div>
    <Tag to="/admin/login">Hello</Tag>
    <PostCard
      class="margin"
      v-for="post in result.posts"
      :id="post.id"
      :title="post.title"
      :plainBody="post.plainBody"
      :thumbnail="post.thumbnail"
      :key="post.id"></PostCard>
  </div>
</template>

<script lang="ts">
import { Context } from "@nuxt/types";
import { PostsApi, TagsApi } from "~/client";
import { buildConfiguration } from "~/client/configurationFactory"
import { defineComponent } from "@nuxtjs/composition-api"
import Button from "~/components/atoms/Button.vue"
import Tag from "~/components/atoms/Tag.vue"
import PostCard from "~/components/molecules/PostCard.vue"

export default defineComponent({
  async asyncData(ctx: Context): Promise<Object> {
    const post = new PostsApi(buildConfiguration());
    const response = await post.postsGet(1);
    return {
      result: response.data,
    }
  },
  components: {
    Button,
    Tag,
    PostCard
  }
});
</script>

<style lang="scss" scoped>

.margin {
  margin: 10px 5px;
}

</style>
