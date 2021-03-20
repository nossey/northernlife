<template>
  <div>
    <b-container class="pt-4 pb-2">
      <b-row>
        <b-col>
          <SearchForm :on-submit="search"></SearchForm>
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
    let search = "";
    if (ctx.query["search"])
      search = ctx.query["search"] as string;

    let page = 1;
    if (pageQuery)
      page = Number(pageQuery)
    const response = await post.postsGet(page, undefined, search);

    return {
      result: response.data,
      page: page,
    }
  },
  methods: {
    toTagLinks(tags: string[]){
      const links =  Enumerable.from(tags).select(function (t) {return {name: t, link:`/tags/${t}/posts`}}).toArray()
      return {links: links}
    },
    search(text: string){
      this.$nuxt.context.redirect(`/?search=${text}`);
    },
    pageClicked(page :string){
      const pageNumber = Number(page);
      const params: {[key: string]: string} = {};
      if (pageNumber != 1)
      {
        params['page'] = pageNumber.toString();
      }
      const searchWord = this.$nuxt.context.query["search"];
      if (searchWord)
      {
        params['search'] = searchWord as string;
      }

      const p = this.createQueryParams(params);
      this.$nuxt.context.redirect(`/${p}`);
    },
    createQueryParams(params: {[key: string]: string}){
      if(Object.keys(params).length === 0)
        return "";

      let result = "?";
      let added = false;
      for (const key in params)
      {
        if (added)
        {
          result += `&${key}=${params[key]}`
        }
        else
        {
          result += `${key}=${params[key]}`
        }
        added = true
      }
      return result;
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
  watchQuery: ['page', 'search'],
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
