<template>
  <b-container fluid class="post-container pt-2">
    <b-row>
      <b-col><h1>{{state.title}}</h1></b-col>
    </b-row>
    <b-row v-if="state.tags.length > 0">
      <b-col v-for="tag in state.tags" :key="tag">#{{tag}}</b-col>
    </b-row>
    <b-row>
      <b-col>
        <b-img class="img-fluid" :src="state.thumbnail" />
      </b-col>
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
import Tag from "~/components/atoms/Tag.vue"

type Props = {
  body: string,
  title: string
  thumbnail: string,
  tags: string[]
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
    },
    tags: {
      type: Array,
      required: true
    }
  },
  components: {Tag},
  setup(props: Props) {
    const state = reactive({
      thumbnail: computed(() => props.thumbnail),
      title: computed(() => props.title),
      renderedBody: computed(() => markdown(props.body)),
      tags: computed(() => props.tags)
    });

    return {
      state
    }
  },
  name: "Post"
})
</script>

<style scoped lang="scss">
@import "assets/colors.scss";
.post-container {
  background: $background-white;
}

</style>
