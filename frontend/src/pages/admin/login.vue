<template>
  <div>
    <input type="text" v-model="state.userId">
    <input type="password" v-model="state.password">
    <Button @click.native="login">Login</Button></div>
</template>

<script lang="ts">

import {defineComponent, reactive} from "@nuxtjs/composition-api"
import Button from "~/components/atoms/Button.vue"
import {buildConfiguration} from "~/client/configurationFactory";
import {AuthApi} from "~/client";

export default defineComponent( {
  name: "login",
  setup(props, context){
    const state = reactive<{userId: string, password: string}>({
      userId: "",
      password: ""
    });

    const login = async () => {
      const api = new AuthApi(buildConfiguration());
      const res = await api.authLoginPost({userID:state.userId, password: state.password}).catch(err => err.response);
      if (res.status === 200)
        context.root.$router.push("/")
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
