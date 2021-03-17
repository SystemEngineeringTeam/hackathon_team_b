import Vue from "vue";
import Vuex from "vuex";
import axios from "axios";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    lectures: [
      {
        grade: "",
        semester: "",
        Department: "",
        dayofweek: "",
        time: "",
        teacher: "",
        lectureName: "",
      },
    ],
    getLectureList: [
      {
        Department: "",
        semester: "",
        dayofweek: "",
        time: "",
        teacher: "",
        lectureName: "",
        grade: "",
        reviewStarAverage: "",
        index: "",
      },
    ],
  },

  mutations: {
    setALecture(state, lecture) {
      state.lectures[0] = lecture;
    },
    setGetterLectureList(state, lectureList) {
      state.getLectureList[0] = lectureList;
    },
  },

  actions: {
    setLecture(context, lecture) {
      context.commit("setALecture", lecture);
    },
    async postLecture(context) {
      await axios
        .post(
          process.env.VUE_APP_API_LECTURE,
          JSON.stringify(context.state.lectures)
        )
        .then(() => {})
        .catch(() => {
          console.log("lectureのpostに失敗しました");
        });
    },
    async getLecture(context) {
      await axios
        .get(process.env.VUE_APP_API_LECTURE)
        .then((res) => {
          context.commit("setGetterLectureList", res.data);
          console.log(res.data);
        })
        .catch(() => {
          console.log("lectureのgetに失敗しました");
        });
    },
  },
  modules: {},
});
