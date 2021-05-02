<template>
  <div class="profile">
    <b-container>
      <b-row>
        <b-col class="profile-top">このブログの管理人</b-col>
      </b-row>
      <b-row>
        <b-col>
          <b-img fluid src="~/static/vanilla.png"></b-img>
        </b-col>
      </b-row>
      <b-row>
        <b-col class="profile-name">Sho Kitami</b-col>
      </b-row>
      <b-row class="mt-2 pb-2">
        <b-col>Webエンジニアやってます.<br/>嗜む言語はGo, C#, python, Javascript, Typescript, C++. <br/>特にGoとNuxt.js大好きマン</b-col>
      </b-row>
      <b-row class="pb-3">
        <b-col>
          <a href="https://github.com/nossey"><fa :icon="faGithub" class="github"></fa></a>
          <a href="mailto:northernlife.inquiry@gmail.com"><fa :icon="faEnvelope" class="github"></fa></a>
        </b-col>
      </b-row>
    </b-container>
  </div>
</template>

<script lang="ts">
import { defineComponent } from '@nuxtjs/composition-api';
import {faGithub} from "@fortawesome/free-brands-svg-icons";
import {faEnvelope} from "@fortawesome/free-solid-svg-icons";
export default defineComponent({
  setup(){
    return {
      faGithub,
      faEnvelope
    }
  }
})

</script>

<style scoped lang="scss">
@import "~assets/colors";
.profile {
  border-radius: 2px;
  filter: drop-shadow(2px 2px 3px $shadow-color);
  background: $background-white;

  .profile-top {
    font-size: 1.2rem;
    font-weight: bold;
    text-align: center;
  }

  img {
    display: block;
    margin-left: auto;
    margin-right: auto;
    width: 70%;
    border-radius: 50%;
  }

  .profile-name {
    font-size: 1.1rem;
    font-weight: bold;
    text-align: center;
  }

  svg {
    width: 28px;
    height: 28px;
    color: $shadow-color;
    a {
      text-decoration: none;
    }
  }
}

</style>
