<template>
  <button type="button"><slot></slot></button>
</template>

<script lang="ts">
import { defineComponent } from '@nuxtjs/composition-api';

export default defineComponent({
  name: "Button",
});

</script>

<style lang="scss" scoped>
@import "assets/colors";

button {
  background-color: transparent;
  cursor: pointer;
  outline: none;
  padding: 5px 10px;
  appearance: none;
  border-radius: 5%;
  color: $main-theme;
  border: solid $main-theme 1px;
  transition: background-color .4s;
  font-weight: 480;

  &:hover {
    transition: background-color .4s;
    background-color: $main-theme;
    color: $font-color-white;
  }
}

</style>
