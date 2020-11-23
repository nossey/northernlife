<template>
  <div v-html="state.renderedBody"></div>
</template>

<script lang="ts">
import {computed, defineComponent, reactive} from "@nuxtjs/composition-api"
import { createMarkdown } from "safe-marked";
const markdown = createMarkdown();

type Props = {
  body: string
}

export default defineComponent({
  props: {
    body: {
      type: String,
      required: true
    }
  },
  setup(props: Props) {
    const state = reactive({
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
