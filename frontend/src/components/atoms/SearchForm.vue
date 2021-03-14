<template>
  <form @submit.prevent="submit">
    <fa :icon="faSearch"></fa>
    <input type="search" placeholder="検索" v-model="state.text">
  </form>
</template>

<script lang="ts">

import {defineComponent, reactive} from "@nuxtjs/composition-api";
import {faSearch} from "@fortawesome/free-solid-svg-icons";

type Props = {
  onSubmit: (text: string) => {},
}

export default defineComponent({
  props: {
    onSubmit: {
      type: Function,
      required: false
    }
  },
  setup(props: Props){
    const state = reactive({
      text: ""
    });
    const submit = () => {
      if (props.onSubmit)
      {
        props.onSubmit(state.text)
      }
    }

    return {
      submit,
      state,
      faSearch
    }
  }

})

</script>

<style lang="scss" scoped>
@import "assets/colors";

form {
  display: inline-block;
  border-radius: 4px;
  background: $background-white;
  svg {
    color: $shadow-color;
    margin-left: 10px;
  }
  input {
    border: none;
    padding: 3px;
    outline: none;
  }
}

</style>
