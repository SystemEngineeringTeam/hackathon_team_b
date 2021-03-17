<template>
  <div>
    <v-col
      cols="12"
      sm="8"
      md="6"
      lg="4"
      xl="3"
      style="background-color: #f8bbd0"
    >
      <b-card title="">
        <p class="text-3xl">学年</p>
        <b-form-select
          v-model="gradeSelected"
          :options="gradeOptions"
        ></b-form-select>
        <div class="mt-3">
          Selected: <strong>{{ gradeSelected }}</strong>
        </div>
      </b-card>
    </v-col>
  </div>
</template>

<script>
export default {};
</script>

<style>
</style>