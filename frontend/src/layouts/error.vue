<template>
  <div>{{error.statusCode}}</div>
</template>

<script>
export default {
  name: "error",
  props: {
    error: {type: Object, default() {return {}}}
  }
}
</script>

<style scoped>

</style>
