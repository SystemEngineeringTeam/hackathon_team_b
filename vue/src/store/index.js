import Vue from "vue";
import Vuex from "vuex";
import axios from "axios";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    lectures: [
      {
        Grade: "",
        Department: "",
        Semester: "",
        Time: "",
        Teacher: "",
        LectureName: "",
      },
    ],
  },
  mutations: {
    setGrade(state, grade) {
      state.lectures.Grade = grade;
    },
  },
  actions: {
    async postLecture(context) {
      await axios
        .post(
          process.env.VUE_APP_API_LECTURE,
          JSON.stringify(context.state.lectures)
        )
        .then(() => {})
        .catch(() => {
          console.log("lectureのpostに失敗しました🥺");
        });
    },
    async getLecture(context) {
      await axios
        .get(process.env.VUE_APP_API_LECTURE)
        .then((res) => {})
        .catch(() => {
          console.log("lectureのgetに失敗しました🥺");
        });
    },
  },
  modules: {},
});
