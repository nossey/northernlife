<template>
  <div>
    <b-container class="pl-0" v-if="(attachedTags.tags) && attachedTags.tags.length > 0">
      <Tag class="mr-2" :to="`/?tag=${encodeURIComponent(tag.name)}`" v-for="tag in attachedTags.tags" :key="tag.name">{{tag.name}}:{{tag.attachedCount}}</Tag>
    </b-container>
    <b-container class="mt-1 mb-1" v-if="searchWord">
      検索ワード:{{searchWord}}の検索結果
    </b-container>
    <b-container class="mt-1" v-if="tagName">
      タグ:{{tagName}}の検索結果
    </b-container>
    <b-container class="mt-1">
      全{{result.totalCount}}件
    </b-container>
    <b-container class="pl-0 pr-0">
      <b-row>
        <b-container>
          <b-row>
            <b-col cols="12" md="9">
              <b-container>
                <b-row>
                  <b-col v-for="post in result.posts" :key="post.id" cols="12" lg="6">
                    <PostCard
                      class="mb-2"
                      :id="post.id"
                      :title="post.title"
                      :plainBody="post.plainBody"
                      :thumbnail="post.thumbnail"
                      :tag-link-list="toTagLinks(post.tags)"
                      :posted-at="post.createdAt"
                    ></PostCard>
                  </b-col>
                </b-row>
              </b-container>
            </b-col>
            <b-col md="3" class="d-none d-md-block">
              <Profile></Profile>
            </b-col>
          </b-row>
        </b-container>
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
import { PostsApi, TagsApi } from "~/client";
import { buildConfiguration } from "~/client/configurationFactory"
import Enumerable from "linq";
import { defineComponent } from "@nuxtjs/composition-api"
import Button from "~/components/atoms/Button.vue"
import Tag from "~/components/atoms/Tag.vue"
import PostCard from "~/components/molecules/PostCard.vue"
import Profile from "~/components/atoms/Profile.vue";

export default defineComponent({
  async asyncData(ctx: Context): Promise<Object> {
    const post = new PostsApi(buildConfiguration());
    const tag = new TagsApi(buildConfiguration());
    const pageQuery = ctx.query["page"];
    let search = "";
    if (ctx.query["search"])
      search = ctx.query["search"] as string;
    let tags: Array<string> | undefined = undefined;
    const tagName = ctx.query["tag"] as string;
    if (ctx.query["tag"])
    {
      tags = Array<string>();
      tags.push(tagName);
    }

    let page = 1;
    if (pageQuery)
      page = Number(pageQuery)
    const response = await post.postsGet(page, tags, search);
    const tagResponse = await tag.tagsGet();

    return {
      result: response.data,
      attachedTags: tagResponse.data,
      page: page,
      searchWord: search,
      tagName: tagName
    }
  },
  methods: {
    toTagLinks(tags: string[]){
      const links =  Enumerable.from(tags).select(function (t) {return {name: t, link:`/?tag=${encodeURIComponent(t)}`}}).toArray()
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
      const tag = this.$nuxt.context.query['tag'];
      if (tag)
      {
        params['tag'] = tag as string;
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
  watchQuery: ['page', 'search', 'tag'],
  components: {
    Button,
    Tag,
    PostCard,
    Profile
  }
});
</script>

<style lang="scss" scoped>

.home-enter-active, .home-leave-active { transition: opacity .5s; }
.home-enter, .home-leave-active { opacity: 0; }

</style>
