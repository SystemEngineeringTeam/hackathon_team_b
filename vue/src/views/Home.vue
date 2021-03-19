<template>
  <div id="app">
    <v-app>
      <v-container class="mt-6">
        <!-- 最大画面の表示は正常だが、中画面〜スマホサイズのレスポンシブは最悪
      最大画面の表示で正常な処理ができれば、後からでも中画面のレスポンシブは変えれるので
      今は処理が出来ることを目指す。
      -->
        <v-row class="lighten-3" style="height: 200px">
          <v-col cols="12" sm="8" md="6" lg="4" xl="3">
            <b-card title="">
              <p class="text-3xl">学期</p>
              <v-chip-group v-model="semester" column>
                <v-chip filter outlined>前期</v-chip>
                <v-chip filter outlined>後期</v-chip>
              </v-chip-group>
              <div class="mt-2">
                Selected: <strong>{{ semester }}</strong>
              </div>
            </b-card>
          </v-col>
          <v-col cols="12" sm="8" md="6" lg="4" xl="3">
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
          <v-col cols="12" sm="8" md="6" lg="4" xl="3">
            <b-card title="">
              <p class="text-3xl">時限</p>
              <b-form-select
                v-model="timeSelected"
                :options="timeOptions"
              ></b-form-select>
              <div class="mt-3">
                Selected: <strong>{{ timeSelected }}</strong>
              </div>
            </b-card>
          </v-col>
          <v-col cols="12" sm="8" md="6" lg="4" xl="3">
            <b-card title="">
              <p class="text-3xl">曜日</p>
              <b-form-select
                v-model="dayofweekSelected"
                :options="dayofweekOptions"
              ></b-form-select>
              <div class="mt-3">
                Selected: <strong>{{ dayofweekSelected }}</strong>
              </div>
            </b-card>
          </v-col>
          <v-col cols="12" sm="8" md="6" lg="4" xl="3">
            <b-card title="">
              <p class="text-3xl">専攻</p>
              <b-form-select
                v-model="DepartmentSelected"
                :options="DepartmentOptions"
              ></b-form-select>
              <div class="mt-3">
                Selected: <strong>{{ DepartmentSelected }}</strong>
              </div>
            </b-card>
          </v-col>
          <v-col cols="12" sm="8" md="6" lg="4" xl="3">
            <b-card title="">
              <p class="text-3xl">講師</p>
              <div>
                <b-form-textarea
                  id="textarea"
                  v-model="teacherText"
                  placeholder="Enter something..."
                  rows="1"
                  max-rows="1"
                ></b-form-textarea>
              </div>
              <div class="mt-3">
                Selected: <strong>{{ teacherText }}</strong>
              </div>
            </b-card>
          </v-col>
          <v-col cols="12" sm="8" md="6" lg="4" xl="3">
            <b-card title="">
              <p class="text-3xl">科目名</p>
              <div>
                <b-form-textarea
                  id="lectureName"
                  v-model="lectureName"
                  placeholder="Enter something..."
                  rows="1"
                  max-rows="1"
                ></b-form-textarea>
              </div>
              <div class="mt-3">
                Selected: <strong>{{ lectureName }}</strong>
              </div>
            </b-card>
          </v-col>
        </v-row>
      </v-container>
      <div class="absolute right-2 bottom-2">
        <v-btn color="primary" @click="Clicksubmit(), lectureGetParams()" fab x-large dark>
          <v-icon>mdi-layers-search</v-icon>
        </v-btn>
        <!-- <button @click="urlpush()">push</button> -->
      </div>
    </v-app>
  </div>
</template>

<script>
import axios from "axios";
export default {
  data() {
    return {
      semester: null,
      gradeSelected: null,
      gradeOptions: [
        { value: null, text: "Please select an option" },
        { value: "1", text: "1年" },
        { value: "2", text: "2年" },
        { value: "3", text: "3年" },
        { value: "4", text: "4年" },
      ],
      timeSelected: null,
      timeOptions: [
        { value: null, text: "Please select an option" },
        { value: "1限", text: "1限" },
        { value: "2限", text: "2限" },
        { value: "3限", text: "3限" },
        { value: "4限", text: "4限" },
        { value: "5限", text: "5限" },
        { value: "6限", text: "6限" },
      ],
      dayofweekSelected: null,
      dayofweekOptions: [
        { value: null, text: "Please select an option" },
        { value: "月曜日", text: "月曜日" },
        { value: "火曜日", text: "火曜日" },
        { value: "水曜日", text: "水曜日" },
        { value: "木曜日", text: "木曜日" },
        { value: "金曜日", text: "金曜日" },
        { value: "土曜日", text: "土曜日" },
      ],
      DepartmentSelected: null,
      DepartmentOptions: [
        { value: null, text: "Please select an option" },
        { value: "ee", text: "電気工学専攻" },
        { value: "ev", text: "電子情報工学専攻" },
        { value: "cc", text: "応用化学専攻" },
        { value: "cb", text: "バイオ環境化学専攻" },
        { value: "mm", text: "機械工学専攻" },
        { value: "mp", text: "機械創造工学専攻" },
        { value: "dd", text: "土木工学専攻" },
        { value: "ds", text: "防災土木工学専攻" },
        { value: "fa", text: "建築学専攻" },
        { value: "fl", text: "住居デザイン専攻" },
        { value: "ht", text: "経営情報システム専攻" },
        { value: "hh", text: "スポーツマネジメント専攻" },
        { value: "kk", text: "コンピュータシステム専攻" },
        { value: "kx", text: "メディア情報専攻" },
      ],
      teacherText: null,
      lectureName: null,
      loader: null,
      loading: false,
    };
  },
  watch: {
    loader() {
      const l = this.loader;
      this[l] = !this[l];
      setTimeout(() => (this[l] = false), 1500);
      this.loader = null;
    },
  },
  methods: {
    Clicksubmit: function () {
      this.$store.dispatch("setLecture", {
        grade: this.gradeSelected,
        semester: this.semester,
        Department: this.DepartmentSelected,
        dayofweek: this.dayofweekSelected,
        time: this.timeSelected,
        teacher: this.teacherText,
        lectureName: this.lectureName,
      });
    },
    lectureGetParams: function () {
      axios
        .get(
          "http://localhost:3030/lecture" +
            "?grade=" +
            this.$store.state.p.lectures[0].grade +
            "&semester=" +
            this.$store.state.p.lectures[0].semester +
            "&Department=" +
            this.$store.state.p.lectures[0].Department +
            "&dayofweek=" +
            this.$store.state.p.lectures[0].dayofweek +
            "&time=" +
            this.$store.state.p.lectures[0].time +
            "&teacher=" +
            this.$store.state.p.lectures[0].teacher +
            "&lectureName=" +
            this.$store.state.p.lectures[0].lectureName
        )
        .then((res) => {
          console.log(res);
          // console.log(res.data.length);
          this.$store.commit("removeLectureList");

          for (var i = 0; i < res.data.length; i++) {
            this.$store.commit("addGetLectureList", res.data[i]);
          }
          this.$router.push("/lectures");
        })
        .catch(function (error) {
          console.log(error);
          this.$router.push("/lectures");
        });
    },
  },
};
</script>

<style>
</style>
