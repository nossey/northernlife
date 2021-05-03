<template>
  <b-container class="post-container pt-2">
    <b-row>
      <b-col class="title-area"><h1 class="mb-0">{{state.title}}</h1></b-col>
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
      <b-col class="thumbnail-area">
        <img :src="state.thumbnail">
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

  .title-area{
    h1 {
      background: $background-color;
      border-left: solid 8px $main-theme;
      border-bottom: solid 3px $shadow-color;
      padding: 0.2em 0 0.2em 0.3em;
    }
  }

  .posted-time-area {
    font-size: 12px;
  }

  .thumbnail-area{
    img {
      width: 100%;
      height: auto;
    }
  }

  .rendered-area {
    ::v-deep {
      letter-spacing: .05rem;
      font-size: 1.1rem;
    }
    ::v-deep img {
      max-width: 100%;
      height: auto;
    }
    ::v-deep h1, ::v-deep h2, ::v-deep h3, ::v-deep h4 {
      color: $font-color-white;
      background-color: $main-theme;
      margin-top: 0.25em;
      padding: 0.2em 0 0.2em 0.4em;
    }
  }

  .tag {
    margin-right: 5px;
  }
}

</style>
