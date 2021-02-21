<template>
  <div>
    <b-container class="pt-4 pb-2">
      <b-row>
        <b-col>
          <SearchForm></SearchForm>
        </b-col>
      </b-row>
    </b-container>
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
      <b-pagination
         :total-rows="result.totalCount"
         :per-page="result.perPageCount"
         v-model="page"
         @input="pageClicked"
      ></b-pagination>
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
import SearchForm from "~/components/atoms/SearchForm.vue"

export default defineComponent({
  async asyncData(ctx: Context): Promise<Object> {
    const post = new PostsApi(buildConfiguration());
    const pageQuery = ctx.query["page"];
    let page = 1;
    if (pageQuery)
      page = Number(pageQuery)
    const response = await post.postsGet(page);
    const pageClicked = (page :string) => {
      const pageNumber = Number(page);
      if (pageNumber === 1)
      {
        ctx.redirect(`/`);
      }
      else
        ctx.redirect(`/?page=${pageNumber}`);
    }
    return {
      result: response.data,
      page: page,
      pageClicked: pageClicked
    }
  },
  async validate(ctx: Context){
    const page = ctx.query["page"];
    if (!page)
      return true
    if (isNaN(Number(page)))
      return false
    return true
  },
  watchQuery: ["page"],
  setup(){
    const toTagLinks = (tags: string[]) => {
      const links =  Enumerable.from(tags).select(function (t) {return {name: t, link:`/tags/${t}/posts`}}).toArray()
      return {links: links}
    }

    return {
      toTagLinks
    }
  },
  components: {
    Button,
    Tag,
    PostCard,
    SearchForm
  }
});
</script>

<style lang="scss" scoped>

.margin {
  margin: 10px 5px;
}

</style>
