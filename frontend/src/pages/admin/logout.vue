<template>
  <div><Button @click.native="logout">Logout</Button></div>
</template>

<script lang="ts">
import Button from "~/components/atoms/Button.vue"
import { defineComponent } from "@nuxtjs/composition-api"
export default defineComponent( {
  name: "logout",
  middleware: 'auth',
  components: {Button},
  setup(props, context){
    const logout = () => {
      // SPEC:型推論が利かないためWorkAround
      (context.root as any).$auth.logout();
    }

    return {
      logout
    }
  }
})
</script>

<style scoped>

</style>
