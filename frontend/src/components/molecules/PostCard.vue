<template>
  <b-container class="card">
    <b-row>
      <b-col><b-img class="img-fluid" :src="thumbnail" /></b-col>
    </b-row>
    <b-row>
      <b-col class="tag-area pt-2 pb-2"><Tag v-for="tagLink in tagLinkList.links" :key="tagLink.name" :to="tagLink.link" class="tag">{{tagLink.name}}</Tag></b-col>
    </b-row>
    <b-row class="mt-1">
      <b-container>
        <b-row><b-col><h2>{{title}}</h2></b-col></b-row>
        <b-row><b-col><div class="plain-body-area">{{plainBody}}</div></b-col></b-row>
      </b-container>
    </b-row>
    <b-row align-h="end">
      <b-col class="foot">
        <Button :to="`/posts/${id}`">READ</Button>
      </b-col>
    </b-row>
  </b-container>
</template>

<script lang="ts">

import {defineComponent, PropType} from "@nuxtjs/composition-api"
import Button from '~/components/atoms/Button.vue'
import Tag from "~/components/atoms/Tag.vue"

export interface ITagLink {
  name: string
  link: string
}

export interface ITagLinkList {
  links: Array<ITagLink>
}

export default defineComponent( {
  props: {
    id: {
      type: String,
      required: true
    },
    plainBody: {
      type: String,
      required: true
    },
    thumbnail: {
      type: String,
      required:true
    },
    title: {
      type: String,
      required: true
    },
    tagLinkList: {
      type: Object as PropType<ITagLinkList>,
      required: true
    }
  },
  components: {
    Button,
    Tag
  }
})
</script>

<style lang="scss" scoped>
@import "assets/colors";
.card {
  min-width: 400px;
  transition: all .1s;
  filter: drop-shadow(2px 2px 3px $shadow-color);
  &:hover{
    filter: drop-shadow(5px 5px 5px $shadow-color);
    transition: all .1s;
  }
  max-width: 600px;
  padding: 10px 15px;

  .plain-body-area {
    text-overflow: ellipsis;
    overflow: hidden;
    width: 100%;
    height: 50px;
  }

  .tag-area{
    .tag {
      margin-right: 5px;
    }
    height: 30px;
  }

  .foot {
    margin-top: 10px;
    text-align: right;
  }
}
</style>
