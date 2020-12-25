<template>
  <b-container fluid>
    <b-row>
      <b-col>
        <b-img class="img-fluid" :src="state.thumbnail" />
      </b-col>
    </b-row>
    <b-row>
      <b-col>{{state.title}}</b-col>
    </b-row>
    <b-row>
      <b-col>
        <div v-html="state.renderedBody"></div>
      </b-col>
    </b-row>
  </b-container>
</template>

<script lang="ts">
import {computed, defineComponent, reactive} from "@nuxtjs/composition-api"
import { createMarkdown } from "safe-marked";
const markdown = createMarkdown({
  marked:{
    breaks:true
  }
});

type Props = {
  body: string,
  title: string
  thumbnail: string
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
    thumbnail: {
      type: String,
      required: true
    }
  },
  setup(props: Props) {
    const state = reactive({
      thumbnail: computed(() => props.thumbnail),
      title: computed(() => props.title),
      renderedBody: computed(() => markdown(props.body)),
    });
    return {
      state
    }
  },
  name: "Post"
})
</script>

<style scoped>

</style>
