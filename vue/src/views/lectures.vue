<template>
  <div>
    <v-btn @click="sorted()">評価が高い順にソート</v-btn>
    <v-app>
      <ul>
        <li v-for="(Lists, i) in $store.state.g.getLectureList" :key="i">
          <template>
            <v-card class="mx-auto my-12" min-width="500" max-width="700">
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
                <div class="my-4 subtitle-1"></div>
                <div>学年 : {{ Lists.grade }}</div>
                <div>曜日 : {{ Lists.dayofweek }}</div>
                <div>時限 : {{ Lists.time }}</div>
                <div>専攻 : {{ Lists.department }}</div>
              </v-card-text>

              <v-divider class="mx-4"></v-divider>
              <v-card-title></v-card-title>
              <v-card-actions>
                <v-btn
                  color="deep-purple lighten-2"
                  text
                  @click="reviewDetail(Lists.indexLectureNumber - 1)"
                >
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
import axios from "axios";
export default {
  methods: {
    reviewDetail: function (index) {
      axios
        .get(
          "http://localhost:3030/review?indexLectureNumber=" +
            this.$store.state.g.getLectureList[index].indexLectureNumber
        )
        .then((res) => {
          console.log(res.data);
          console.log(res.data.length);
          this.$store.commit("removeReviewList");
          this.$store.commit("GetReviewList", res.data);
          this.$router.push("/review");
        })
        .catch(function (error) {
          console.log(error);
          this.$router.push("/review");
        });
    },
    sorted() {
      this.$store.state.g.getLectureList.sort(function (a, b) {
        if (a.reviewStarAverage > b.reviewStarAverage) return -1;
        if (a.reviewStarAverage < b.reviewStarAverage) return 1;
        return 0;
      });
    },
  },
};
</script>

<style>
</style>