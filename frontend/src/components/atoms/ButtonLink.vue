<template>
  <nuxt-link class="button" :to="to"><slot></slot></nuxt-link>
</template>

<script lang="ts">
import {defineComponent} from "@nuxtjs/composition-api";

export default defineComponent({
  name: "ButtonLink",
  props: {
    to: {
      type: String,
      required: true
    }
  },
})

</script>

<style scoped lang="scss">
@import "~assets/colors";
a {
  text-decoration: none;
  cursor: pointer;
  outline: none;
  padding: 5px 10px;
  appearance: none;
  border-radius: 5%;
  color: $font-color-white;
  border: solid $main-theme 1px;
  background: $main-theme;
  transition: all .4s;
  font-weight: 480;

  &:hover {
    transition: all .4s;
    background-color: ToDarkenColor($main-theme);
    border-color: ToDarkenColor($main-theme);
  }

  &:active{
    transition: all .2s;
    border-color: $main-theme;
    background: $main-theme;
    box-shadow: 0 1px 1px 1px ToDarkenColor($main-theme);
  }
}

</style>
