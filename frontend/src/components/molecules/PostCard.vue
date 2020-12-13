<template>
  <b-container class="card">
    <b-row>
      <b-col><b-img class="img-fluid" :src="thumbnail" /></b-col>
    </b-row>
    <b-row class="mt-1">
      <b-container>
        <b-row><b-col><h2>{{title}}</h2></b-col></b-row>
        <b-row><b-col>{{plainBody}}</b-col></b-row>
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

import {defineComponent} from "@nuxtjs/composition-api"
import Button from '~/components/atoms/Button.vue'
export default defineComponent( {
  name: "PostCard",
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
    }
  },
  components: {
    Button
  }
})
</script>

<style lang="scss" scoped>
@import "assets/colors";
.card {
  max-width: 400px;
  transition: all .1s;
  filter: drop-shadow(2px 2px 3px $shadow-color);
  &:hover{
    filter: drop-shadow(5px 5px 5px $shadow-color);
    transition: all .1s;
  }
  max-width: 600px;
  padding: 10px 15px;

  .foot {
    margin-top: 10px;
    text-align: right;
  }
}
</style>
