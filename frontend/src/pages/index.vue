<template>
  <b-container>
    <b-row>
      <b-col>
        <ul v-if="result.posts.length > 0">
          <li v-for="post in result.posts"><nuxt-link :to="`/posts/${post.id}`">{{post.id}}</nuxt-link> <Button>READ</Button></li>
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
import Button from "~/components/atoms/Button.vue"

export default defineComponent({
  async asyncData(ctx: Context): Promise<Object> {
    const post = new PostsApi(buildConfiguration());
    const response = await post.postsGet(1);
    return {result: response.data}
  },
  components: {
    Button
  }
});
</script>

<style lang="scss" scoped>

li {
  margin-top: 10px;
}

</style>
