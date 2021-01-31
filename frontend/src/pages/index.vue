<template>
  <div>
    <b-container>
      <b-row>
       <b-col v-for="post in result.posts" :key="post.id" cols="12" lg="6">
        <PostCard
          class="margin"
          :id="post.id"
          :title="post.title"
          :plainBody="post.plainBody"
          :thumbnail="post.thumbnail"
          :tag-link-list="toTagLinks(post.tags)"
        ></PostCard>
        </b-col>
      </b-row>
    </b-container>
  </div>
</template>

<script lang="ts">
import { Context } from "@nuxt/types";
import { PostsApi } from "~/client";
import { buildConfiguration } from "~/client/configurationFactory"
import Enumerable from "linq";
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
  setup(){
    const toTagLinks = (tags: string[]) => {
      const links =  Enumerable.from(tags).select(function (t) {return {name: t, link:`/tags/${t}`}}).toArray()
      return {links: links}
    }

    return {
      toTagLinks
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
