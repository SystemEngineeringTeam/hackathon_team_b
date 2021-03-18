<template>
<div>
  <v-app>
    <ul>
      <li v-for="review,i in $store.state.r.reviewList" :key="i">
        <v-card class="mx-auto m-6" min-width="300" max-width="700">
          <v-card-title>名無しさん</v-card-title>
          <p>{{review.sentence}}</p>
          <v-card-text>
            <v-row align="center" class="mx-0">
              <v-rating
                :value="review.reviewStar"
                color="amber"
                dense
                half-increments
                readonly
                size="14"
              ></v-rating>
              <div class="grey--text ml-4">{{review.reviewStar}}</div>
            </v-row>
          <div class="my-4 subtitle-1"></div>
          <div>
            {{ review.sentence }}
          </div>
        </v-card-text>
        <v-divider class="mx-4"></v-divider>
        </v-card>
      </li>
    </ul>
    <!-- ここからレビュー投稿-->
    <v-row justify="center">
      <v-dialog v-model="dialog" persistent max-width="600px">
        <template v-slot:activator="{ on, attrs }">
          <v-btn color="primary" dark class="m-6" v-bind="attrs" v-on="on">
            この講義へのレビュー
          </v-btn>
        </template>
        <v-card>
          <v-card-title>
            <span class="headline">レビュー投稿</span>
          </v-card-title>
          <v-card-text>
            <v-container>
              <v-row>
                <v-col cols="12">
                  <v-text-field
                    label="TextArea*"
                    v-model="text"
                    required
                  ></v-text-field>
                </v-col>
                <v-col cols="12"> </v-col>
                <v-col cols="12" sm="6">
                  <v-select
                    :items="[1, 2, 3, 4, 5]"
                    label="Star*"
                    required
                    v-model="reviewStar"
                  ></v-select>
                </v-col>
              </v-row>
            </v-container>
            <small>*が入ってる項目は必須入力です</small>
          </v-card-text>
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn color="blue darken-1" text @click="dialog = false">
              cansel
            </v-btn>
            <v-btn
              color="blue darken-1"
              text
              @click="(dialog = false), registerReview(), consoleunko()"
            >
              Save
            </v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>
    </v-row>
    <!-- ここまでレビュー投稿-->
  </v-app>
  </div>
</template>

<script>
export default {
  data() {
    return {
      reviewStar: null,
      text: "",
      dialog: false,
    };
  },
  methods: {
    consoleunko() {
      console.log(this.text);
      console.log(this.reviewStar);
    },
    registerReview() {
      axios
        .post("http://localhost:3030/review", {
          IndexLectureNumber: this.$store.state.r.reviewList[0]
            .indexLectureNumber,
          ReviewStar: this.reviewStar,
          sentence: this.text,
        })
        .then((res) => {
          console.log(res.config.data);
          var test = {
            IndexLectureNumber: this.$store.state.r.reviewList[0].indexLectureNumber,
            reviewStar: this.reviewStar,
            sentence: this.text
          };
          this.$store.state.r.reviewList.push(test)
          console.log(this.$store.state.r.reviewList)
          //console.log("unko");
        })
        .catch(function (error) {
          console.log(error);
        });
    },
  },
};
</script>

<style>
</style>