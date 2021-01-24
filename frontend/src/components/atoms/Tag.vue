<template>
  <div class="tag" @click="clickedEvent"><slot></slot></div>
</template>

<script lang="ts">
import {defineComponent} from "@nuxtjs/composition-api";

export default defineComponent({
  props: {
    to: {
      type: String,
      required: false
    }
  },
  setup(props, context){
    const clickedEvent = () => {
      if (props.to){
        context.root.$router.push(props.to)
      }
      else{
        context.emit('click')
      }
    }
    return {
      clickedEvent
    }
  }
})

</script>

<style lang="scss" scoped>
@import "assets/colors";
.tag {
  background: $background-white;
  display: inline;
  color: $main-theme;
  border: $main-theme solid 1px;
  border-radius: 2px;
  padding: 3px 5px 3px 5px;
}

</style>
