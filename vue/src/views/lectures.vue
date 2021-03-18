<template>
  <div>
    {{ grade }}:{{ Department }}:{{ semester }}
    <h3>hello</h3>
  </div>
</template>

<script>
export default {
  name:"Getlectures",
  props:['grade','Department','semester','dayofweek','time','teacher','lectureName','reviewStarAverage','indexLectureNumber'],
};
</script>

<style>
</style>