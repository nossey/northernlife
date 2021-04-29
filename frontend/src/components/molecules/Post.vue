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

const MarkdownIt = require('markdown-it');
const sanitizer = require('markdown-it-sanitizer');
const emoji = require('markdown-it-emoji');
const imsize = require('markdown-it-imsize');
import markdownItAnchor from 'markdown-it-anchor'
const markdownItTableOfContents = require('markdown-it-table-of-contents')

import hljs from 'highlight.js'

let tocs: Array<ITocItem> = new Array();
let anchorId = 1;
let tocId = 1;
const md = new MarkdownIt({
  html:true,
  breaks:true,
  langPrefix: 'hljs ',
  highlight: function(code, lang){
    return hljs.highlightAuto(code, [lang]).value
  }
})
.use(sanitizer)
.use(emoji)
.use(imsize)
.use(markdownItAnchor, {
  slugify: function(s){
    tocs.push({name: s, link: `#${anchorId.toString()}`})
    return (anchorId++);
  }
})
.use(markdownItTableOfContents, {
  includeLevel: [1, 2, 3],
  slugify: function(s){
    return (tocId++)
  }
})

import Tag from "~/components/atoms/Tag.vue"

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

export interface ITocItem {
  name: string
  link: string
}

export interface ITocLinkList {
  tableOfContents: Array<ITocItem>
}

import {markdown as renderMarkdown} from "~/application/posts/markdown"

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
        return renderMarkdown(props.body)[0];
      }),
      tags: computed(() => props.tags),
      linkList: computed(() => props.linkList),
      postedAt: computed(() => props.postedAt)
    });

    return {
      state
    }
  }
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
