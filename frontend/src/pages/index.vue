<template>
  <div class="container">
    <div>
      <Footer></Footer>
      <img src="/golang.png" />
      <logo />
      <h1 class="title">
        Hello World
      </h1>
      <h2 class="subtitle">
        My first-rate Nuxt.js project
      </h2>
      <div class="links">
        <a href="https://nuxtjs.org/" target="_blank" class="button--green">
          Documentation
        </a>
        <a
          href="https://github.com/nuxt/nuxt.js"
          target="_blank"
          class="button--grey"
        >
          GitHub
        </a>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Logo from "~/components/Logo.vue";
import Footer from "~/components/Footer.vue";
import { Context } from "@nuxt/types";

export default Vue.extend({
  components: {
    Logo,
    Footer
  },
  async asyncData(ctx: Context): Promise<void> {
    //const result = await ctx.app.$axios.$get("/secrets/");
    //console.log(result);
  }
});
</script>

<style>
.container {
  margin: 0 auto;
  min-height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  text-align: center;
}

.title {
  font-family: "Quicksand", "Source Sans Pro", -apple-system, BlinkMacSystemFont,
    "Segoe UI", Roboto, "Helvetica Neue", Arial, sans-serif;
  display: block;
  font-weight: 300;
  font-size: 100px;
  color: #35495e;
  letter-spacing: 1px;
}

.subtitle {
  font-weight: 300;
  font-size: 42px;
  color: #526488;
  word-spacing: 5px;
  padding-bottom: 15px;
}

.links {
  padding-top: 15px;
}
</style>
