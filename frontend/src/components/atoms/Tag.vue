<template>
  <button class="tag" @click="clickedEvent"><slot></slot></button>
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
      else {
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
@import "~assets/colors";
.tag {
  font-size: 12px;
  background: $background-white;
  display: inline;
  color: $main-theme;
  border: $main-theme solid 1px;
  border-radius: 4px;
  padding: 3px 6px 3px 6px;
  transition: .5s;
  outline: none;

  &:hover {
    cursor: pointer;
    background-color: $main-theme;
    color: $font-color-white;
    transition: .5s;
  }
  &:active{
    background-color: ToDarkenColor($main-theme);
  }
}

</style>
