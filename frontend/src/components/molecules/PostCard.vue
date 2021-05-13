<template>
  <b-container class="card">
    <div class="img-area">
      <b-img :src="thumbnail" />
    </div>
    <b-container>
      <b-row class="mt-1">
        <b-container>
          <b-row class="title-area"><b-col>{{title}}</b-col></b-row>
          <b-row><b-col class="posted-time-area">{{convertTime(postedAt)}}</b-col></b-row>
          <b-row>
            <b-col class="tag-area" v-if="tagLinkList.links.length > 0">
              <Tag v-for="tagLink in tagLinkList.links" :key="tagLink.name" :to="tagLink.link" class="tag mt-1">{{tagLink.name}}</Tag>
            </b-col>
          </b-row>
          <b-row><b-col><div class="plain-body-area">{{plainBody}}</div></b-col></b-row>
        </b-container>
      </b-row>
      <b-row align-h="end">
        <b-col class="foot">
          <ButtonLink :to="`/posts/${id}`">READ</ButtonLink>
        </b-col>
      </b-row>
    </b-container>
  </b-container>
</template>

<script lang="ts">

import {defineComponent, PropType} from "@nuxtjs/composition-api"
import Button from '~/components/atoms/Button.vue'
import ButtonLink from '~/components/atoms/ButtonLink.vue'
import Tag from "~/components/atoms/Tag.vue"
import moment from 'moment';

const convert = (time: string) => {
  moment("time", "YYYY-MM-DD")
}

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
    postedAt: {
      type: String,
      require: false
    },
    tagLinkList: {
      type: Object as PropType<ITagLinkList>,
      required: true
    }
  },
  setup(){
   const convertTime = (time: string) => {
     return moment(time).format("YYYY/MM/DD")
   }

   return {
     convertTime
   }
  },
  components: {
    Button,
    ButtonLink,
    Tag
  }
})
</script>

<style lang="scss" scoped>
@import "~assets/colors";
.card {
  transition: all .1s;
  filter: drop-shadow(2px 2px 3px $shadow-color);
  &:hover{
    filter: drop-shadow(5px 5px 5px $shadow-color);
    transition: all .1s;
  }
  max-width: 600px;
  padding: 0px;

  .img-area {
    width: 100%;
    padding-top: 56.25%;
    position: relative;

    img {
      position: absolute;
      top: 0;
      left: 0;
      bottom: 0;
      right: 0;
      width: 100%;
      height: 100%;
    }
  }

  .title-area {
    margin-bottom: 0;
    font-weight: 500;
    font-size: 1.5rem;
  }

  .posted-time-area {
    font-size: 12px;
  }

  .plain-body-area {
    text-overflow: ellipsis;
    overflow: hidden;
    width: 100%;
    line-height: 1rem;
    height: 2rem;
    font-size: .875rem;
    font-weight: 400;
  }

  .tag-area{
    .tag {
      margin-right: 5px;
    }
    min-height: 30px;
    height: auto;
  }

  .foot {
    margin-top: 10px;
    padding-bottom: 15px;
    text-align: right;
  }
}
</style>
