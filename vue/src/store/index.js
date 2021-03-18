import Vue from "vue";
import Vuex from "vuex";
//import axios from "axios";

Vue.use(Vuex);

const postData = {
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
  },
  mutations: {
    setALecture(state, lecture) {
      state.lectures[0] = lecture;
    },
  },
  actions: {
    setLecture(context, lecture) {
      context.commit("setALecture", lecture);
    },
    /*
    async postLecture(context) {
      await axios
        .post(
          process.env.VUE_APP_API_LECTURE,
          JSON.stringify(context.state.lectures)
        )
        .then((res) => {
          context.commit("addGetLectureList", res.data);
        })
        .catch(() => {
          console.log("lectureのpostに失敗しました");
        });
    },
    */
  },
};

const getData = {
  state: {
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
    addGetLectureList(state, lectureList) {
      state.getLectureList.push(lectureList);
    },
    //まだ
    removeLectureList(state) {
      state.getLectureList.splice(0);
    },
  },
};

export default new Vuex.Store({
  modules: {
    p: postData,
    g: getData,
  },
});
