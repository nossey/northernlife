<template>
  <div>
    <input type="text" v-model="state.userId">
    <input type="password" v-model="state.password">
    <Button @click="login">Login</Button></div>
</template>

<script lang="ts">

import {defineComponent, reactive} from "@nuxtjs/composition-api"
import Button from "~/components/atoms/Button.vue"

export default defineComponent( {
  name: "login",
  setup(props, context){
    const state = reactive<{userId: string, password: string}>({
      userId: "",
      password: ""
    });

    const login = async () => {
      // SPEC:型推論が利かないためWorkAround
      await (context.root as any).$auth.loginWith('local', {
       data: {
         userID: state.userId,
         password: state.password,
       }
      })
    }
    return  {
      login,
      state
    }
  },
  components: {
    Button
  }
})
</script>

<style scoped>

</style>
