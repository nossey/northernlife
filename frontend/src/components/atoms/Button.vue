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
}

</style>
