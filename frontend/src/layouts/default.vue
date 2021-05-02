<template>
  <div>
    <Header></Header>
    <b-container fluid class="main-content">
      <b-row>
        <b-col>
          <b-container>
            <b-row>
              <b-col class="no-padding-sm head-area mt-4 mb-4">
                <SearchPostForm class="searchForm"></SearchPostForm>
                <a href="https://github.com/nossey"><fa :icon="faGithub" class="github"></fa></a>
              </b-col>
            </b-row>
            <b-row>
             <b-col class="no-padding-sm">
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
import {faGithub} from "@fortawesome/free-brands-svg-icons";

export default {
  components: {Header, SearchPostForm},
  setup(){
    return {
      faGithub
    }
  }
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

.head-area {
  position: relative;
  .searchForm {
    display: inline-block;
  }
  a {
    color: $shadow-color;
    .github {
      display: inline-block;
      position: absolute;
      right: 0;
      width: 28px;
      height: 28px;
    }
    text-decoration: none;
  }
}

@media only screen and (max-device-width : 640px) {
  .no-padding-sm {
    padding-left: 0;
    padding-right: 0;
  }
}
</style>
