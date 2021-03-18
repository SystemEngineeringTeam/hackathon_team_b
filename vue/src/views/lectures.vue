<template>
  <div>
    <p v-for="lectureList in this.$store.state.g.getLectureList" v-bind:key="lectureList">
      {{ lectureList }}
    </p>
  </div>
</template>

<script>
export default {};
</script>

<style>
</style>