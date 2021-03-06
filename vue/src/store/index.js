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
        department: "",
        semester: "",
        dayofweek: "",
        time: "",
        teacher: "",
        lectureName: "",
        grade: "",
        reviewStarAverage: 1,
        indexLectureNumber: "",
      },
    ],
  },
  mutations: {
    addGetLectureList(state, lectureList) {
      state.getLectureList.push(lectureList);
    },
    removeLectureList(state) {
      state.getLectureList.splice(0);
    },
  },
};

const reviewData = {
  state: {
    reviewList: [
      {
        indexLectureNumber: "",
        reviewStar: "",
        sentence: "",
      },
    ],
  },
  mutations: {
    //reviewsが配列
    GetReviewList(state, reviews) {
      for (var i = 0;i < reviews.length; i++){
        state.reviewList.push(reviews[i])
      }
      // state.reviewList=review;
    },
    removeReviewList(state) {
      state.reviewList.splice(0);
    },
  },
};

export default new Vuex.Store({
  modules: {
    p: postData,
    g: getData,
    r: reviewData,
  },
});
