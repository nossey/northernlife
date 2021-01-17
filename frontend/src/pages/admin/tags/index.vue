<template>
  <div>
    <b-container>
      <nuxt-link to="/admin">Admin Top</nuxt-link>
      <b-row>
        <b-col>
          <input type="text" v-model="state.newTag">
          <Button @click.native="tryCreateTag">Create</Button>
        </b-col>
      </b-row>
      <b-row>
        <b-col>
          <div v-if="fetchState.pending">Loading...</div>
          <div v-else-if="fetchState.error">Error Occured</div>
          <div v-else>{{state.tags}}</div>
        </b-col>
      </b-row>
    </b-container>
  </div>
</template>

<script lang="ts">

import {defineComponent, useFetch, reactive} from "@nuxtjs/composition-api"
import {AdminTagsApi} from "~/client"
import {buildConfiguration} from "~/client/configurationFactory"
import Button from "~/components/atoms/Button.vue"

export default defineComponent({
  middleware: ['auth'],
  components: {Button},
  setup() {
    const state = reactive({
      tags: new Array<string>(),
      newTag: ""
    });

    const api = new AdminTagsApi(buildConfiguration());
    const{fetch, fetchState} = useFetch(async() => {
      const result = (await api.adminTagsGet()).data;
      state.tags = (result.tags) ? result.tags : new Array<string>();
    })
    const tryCreateTag = async() => {
      await api.adminTagsPost({tagName: state.newTag})
        .then(res => {
          state.newTag = ""
          fetch();
        })
        .catch(err => {
        // TODO:エラートーストとか出してあげる
        console.log(err);
      });
    }

    return {
      state,
      tryCreateTag,
      fetch,
      fetchState
    }
  }
})

</script>

<style scoped>

</style>
