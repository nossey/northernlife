<template>
  <div><div @click="postman">ClickME!</div></div>
</template>

<script lang="ts">

import {defineComponent} from "@nuxtjs/composition-api";
import {AuthApi, Configuration} from "~/client";
import {buildConfiguration} from "~/client/configurationFactory";

export default defineComponent({
  name: "post",
  middleware: 'auth',
  setup(props, context) {
    const postman = async () => {
      const config = buildConfiguration();
      const token = (context.root.$store as any).$auth.getToken('local');
      config.apiKey = token;
      const api = new AuthApi(config);
      const user = await api.authUserGet();
      console.log(user.data.userID);
    }

    return {
      postman
    }
  }
})
</script>

<style scoped>

</style>
