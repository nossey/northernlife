<template>
  <button type="button" @click.once="tryJumpToLink"><slot></slot></button>
</template>

<script lang="ts">
import { defineComponent } from '@nuxtjs/composition-api';

type Props = {
  to: string
};

export default defineComponent({
  name: "Button",
  props: {
    to: {
      type: String,
      required: false
    }
  },
  setup(props: Props, context){
    const tryJumpToLink = () => {
      if (props.to)
        context.root.$router.push(props.to)
    }
    return {
      tryJumpToLink
    }
  }
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

  &:active{
    transition: all .2s;
    border-color: $main-theme;
    background: $main-theme;
    box-shadow: 0 1px 1px 1px ToDarkenColor($main-theme);
  }
}

</style>
