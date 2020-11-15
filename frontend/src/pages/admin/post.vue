<template>
  <div>
    Title
    <input type="text" v-model="props.createPost.title">
    Body
    <input type="text" v-model="props.createPost.body">
    PlainBody
    <input type="text" v-model="props.createPost.plain_body">
    <Button @click.native="postman">POST</Button>
  </div>
</template>

<script lang="ts">

import {defineComponent} from "@nuxtjs/composition-api";
import {PostsApi, ModelPostCreateBody} from "~/client";
import {buildConfiguration} from "~/client/configurationFactory";
import Button from "~/components/atoms/Button.vue"

type Props = {
  createPost: ModelPostCreateBody,
  isPosting: boolean
}

export default defineComponent({
  components: {
    Button
  },
  props: {
    CreateBody:  {
      type: Object as () => ModelPostCreateBody,
    }
  },
  name: "post",
  middleware: 'auth',
  setup(props:Props, context) {
    props.createPost = {title: "Title", body: "Body", plain_body: "This will be the day"}
    props.isPosting = false
    const postman = async () => {
      //if (props.isPosting)
      //  return;
      //props.isPosting = true;
      const config = buildConfiguration();
      const token = (context.root.$store as any).$auth.getToken('local');
      config.apiKey = token;
      const api = new PostsApi(config);
      await api.postsPost(props.createPost).then(res => {
        context.root.$router.push(`/posts/${res.data.post_id}`)
      }).catch(err => {
        // TODO:トーストとか色々出してあげる
        console.log(err.statusCode)
        props.isPosting = false;
      });
    }

    return {
      props,
      postman
    }
  }
})
</script>

<style scoped>

</style>
