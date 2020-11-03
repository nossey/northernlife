<template>
  <b-container class="card">
    <b-row>
     <b-col>
       <h2>{{title}}</h2>
     </b-col>
    </b-row>
    <b-row>
      <b-col>{{body}}</b-col>
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
    body: {
      type: String,
      required: true
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
  border: $sub-theme 2px solid;
  max-width: 600px;
  padding: 10px 15px;

  .foot {
    margin-top: 10px;
    text-align: right;
  }
}
</style>
