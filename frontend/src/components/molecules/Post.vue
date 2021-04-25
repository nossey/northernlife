<template>
  <b-container class="post-container pt-2">
    <b-row>
      <b-col><h1 class="mb-0">{{state.title}}</h1></b-col>
    </b-row>
    <b-row>
      <b-col v-if="state.postedAt" class="posted-time-area">
        {{state.postedAt}}
      </b-col>
    </b-row>
    <b-row v-if="state.tags.length > 0">
      <b-col class="pt-2 pb-2">
        <Tag v-for="link in state.linkList.links" :key="link.name" :to="link.link" class="tag">{{link.name}}</Tag>
      </b-col>
    </b-row>
    <b-row>
      <b-col>
        <b-img class="img-fluid" :src="state.thumbnail" />
      </b-col>
    </b-row>
    <b-row>
      <b-col class="rendered-area">
        <div v-html="state.renderedBody"></div>
      </b-col>
    </b-row>
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/highlight.js/9.15.10/styles/agate.min.css">
  </b-container>
</template>

<script lang="ts">
import {computed, defineComponent, reactive, PropType} from "@nuxtjs/composition-api"

let MarkdownIt = require('markdown-it');
const md = new MarkdownIt();

import { createMarkdown } from "safe-marked";
import hljs from 'highlight.js'
const renderer = new marked.Renderer();
const toc = Array<{level: Number, slug: string, title: string}>();
renderer.heading = (text, level) => {
  const slug = encodeURI(text.toLowerCase());
  toc.push({
    level: level as Number,
    slug: slug,
    title: text as string
  })
  return '<h'
    + level
    + ' id="'
    + slug
    + '">'
    + text
    + '</h'
    + level
    + '>\n'
}
const markdown = createMarkdown({
  marked:{
    breaks:true,
    renderer: renderer,
    langPrefix: 'hljs ',
    highlight(code: string, lang: string, callback?: (error: any, code?: string) => void): string | void {
      return hljs.highlightAuto(code, [lang]).value
    }
  }
});

import Tag from "~/components/atoms/Tag.vue"
import marked from "marked";

type Props = {
  body: string,
  title: string,
  thumbnail: string,
  tags: string[],
  linkList: ITagLinkList,
  postedAt: string
}

export interface ITagLink {
  name: string
  link: string
}

export interface ITagLinkList {
 links: Array<ITagLink>
}

export default defineComponent({
  props: {
    title: {
      type: String,
      required: true
    },
    body: {
      type: String,
      required: true
    },
    postedAt: {
     type: String,
     required: false
    },
    thumbnail: {
      type: String,
      required: true
    },
    tags: {
      type: Array,
      required: true
    },
    linkList: {
      type: Object as PropType<ITagLinkList>,
      required: true
    }
  },
  components: {Tag},
  setup(props: Props) {
    const state = reactive({
      thumbnail: computed(() => props.thumbnail),
      title: computed(() => props.title),
      renderedBody: computed(() =>{
        return md.render(props.body)
      }),
      tags: computed(() => props.tags),
      linkList: computed(() => props.linkList),
      toc: computed(() => {
        markdown(props.body);
        return toc;
      }),
      postedAt: computed(() => props.postedAt)
    });

    return {
      state
    }
  },
  name: "Post"
})
</script>

<style scoped lang="scss">
@import "~assets/colors.scss";
.post-container {
  background: $background-white;
  filter: drop-shadow(2px 2px 3px $shadow-color);

  .posted-time-area {
    font-size: 12px;
  }

  .rendered-area {
    ::v-deep img {
      max-width: 100%;
      height: auto;
    }
  }

  .tag {
    margin-right: 5px;
  }
}

</style>
