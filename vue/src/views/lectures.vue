<template>
  <div>
    <v-app>
      <ul>
        <li v-for="Lists in $store.state.g.getLectureList" :key="Lists">
          <template>
            <v-card
              :loading="loading"
              class="mx-auto my-12"
              min-width="500"
              max-width="700"
            >
              <v-card-title
                >{{ Lists.teacher }} {{ Lists.lectureName }}</v-card-title
              >
              <v-card-text>
                <v-row align="center" class="mx-0">
                  <v-rating
                    :value="Lists.reviewStarAverage"
                    color="amber"
                    dense
                    half-increments
                    readonly
                    size="14"
                  ></v-rating>

                  <div class="grey--text ml-4">
                    {{ Lists.reviewStarAverage }}
                  </div>
                </v-row>

                <div class="my-4 subtitle-1">ああああ</div>

                <div>ああああ</div>
              </v-card-text>
              <v-divider class="mx-4"></v-divider>
              <v-card-title>ああああ</v-card-title>
              <v-card-actions>
                <v-btn color="deep-purple lighten-2" text @click="reserve">
                  レビューを見る
                </v-btn>
              </v-card-actions>
            </v-card>
          </template>
        </li>
      </ul>
    </v-app>
  </div>
</template>

<script>
export default {};
</script>

<style>
</style>