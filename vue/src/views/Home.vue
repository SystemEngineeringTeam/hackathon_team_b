<template>
  <div id="app">
    <v-container class="mt-6">
      <!-- 最大画面の表示は正常だが、中画面〜スマホサイズのレスポンシブは最悪
      側だけと最大画面の表示で正常な処理ができれば、後からでもレスポンシブは変えれるので
      今は処理が出来ることを目指す。
      -->
      <v-row class="grey lighten-3" style="height: 200px">
        <v-col
          cols="12"
          sm="8"
          md="6"
          lg="4"
          xl="3"
          style="background-color: #ffcdd2"
        >
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
        <v-col
          cols="12"
          sm="8"
          md="6"
          lg="4"
          xl="3"
          style="background-color: #e1bee7"
        >
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
      </v-row>
      <v-row class="grey lighten-3" style="height: 200px">
        <v-col
          cols="12"
          sm="8"
          md="6"
          lg="4"
          xl="3"
          style="background-color: #ffcdd2"
        >
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
        <v-col
          cols="12"
          sm="8"
          md="6"
          lg="4"
          xl="3"
          style="background-color: #f8bbd0"
        >
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
        <v-col
          cols="12"
          sm="8"
          md="6"
          lg="4"
          xl="3"
          style="background-color: #e1bee7"
        >
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
      </v-row>
      <v-row class="grey lighten-3" style="height: 200px">
        <v-col
          cols="12"
          sm="8"
          md="6"
          lg="4"
          xl="3"
          style="background-color: #e1bee7"
        >
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
    <div class="float-right m-10">
      <v-btn
        class="ma-2"
        :loading="loading"
        :disabled="loading"
        @click="loader = 'loading'"
      >
        Accept Terms
      </v-btn>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      semester: null,
      gradeSelected: null,
      gradeOptions: [
        { value: null, text: "Please select an option" },
        { value: "1年", text: "1年" },
        { value: "2年", text: "2年" },
        { value: "3年", text: "3年" },
        { value: "4年", text: "4年" },
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
        { value: "電気工学専攻", text: "電気工学専攻" },
        { value: "電子情報工学専攻", text: "電子情報工学専攻" },
        { value: "応用化学専攻", text: "応用化学専攻" },
        { value: "バイオ環境化学専攻", text: "バイオ環境化学専攻" },
        { value: "機械工学専攻", text: "機械工学専攻" },
        { value: "機械創造工学専攻", text: "機械創造工学専攻" },
        { value: "土木工学専攻", text: "土木工学専攻" },
        { value: "防災土木工学専攻", text: "防災土木工学専攻" },
        { value: "建築学専攻", text: "建築学専攻" },
        { value: "住居デザイン専攻", text: "住居デザイン専攻" },
        { value: "経営情報システム専攻", text: "経営情報システム専攻" },
        { value: "スポーツマネジメント専攻", text: "スポーツマネジメント専攻" },
        { value: "コンピュータシステム専攻", text: "コンピュータシステム専攻" },
        { value: "メディア情報専攻", text: "メディア情報専攻" },
      ],
      teacherText: "",
      lectureName: "",
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
};
</script>

<style>
</style>