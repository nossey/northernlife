<template>
  <div>
    <Header></Header>
    <b-container fluid class="main-content">
      <b-row>
        <b-col>
          <b-container>
            <b-row>
              <b-col>
                <SearchPostForm class="mt-4 mb-4"></SearchPostForm>
              </b-col>
            </b-row>
            <b-row>
             <b-col>
               <nuxt />
             </b-col>
            </b-row>
          </b-container>
        </b-col>
      </b-row>
    </b-container>
  </div>
</template>

<script lang="ts">
import Header from "~/components/molecules/Header.vue";
import SearchPostForm from "~/components/organism/SearchPostForm.vue";

export default {
  components: {Header, SearchPostForm},
}

</script>

<style lang="scss">
@import "~assets/colors.scss";
html {
  font-family: 'Source Sans Pro', -apple-system, BlinkMacSystemFont, 'Segoe UI',
    Roboto, 'Helvetica Neue', Arial, sans-serif;
  font-size: 16px;
  word-spacing: 1px;
  -ms-text-size-adjust: 100%;
  -webkit-text-size-adjust: 100%;
  -moz-osx-font-smoothing: grayscale;
  -webkit-font-smoothing: antialiased;
  box-sizing: border-box;
}

*,
*:before,
*:after {
  box-sizing: border-box;
  margin: 0;
}

.main-content {
  background: $background-color;
}

</style>
