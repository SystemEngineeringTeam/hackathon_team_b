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
    setLecture(state, Grade, Department, Semester, Time, Teacher, LectureName) {
      state.Grade = Grade;
      state.Department = Department;
      state.Semester = Semester;
      state.Time = Time;
      state.Teacher = Teacher;
      state.LectureName = LectureName;
    },
  },

  actions: {
    setLecture(context,Grade,Department,Semester,Time,Teacher,LectureName){
      context.commit("setLecture",Grade,Department,Semester,Time,Teacher,LectureName)
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
          console.log(res.data);
        })
        .catch(() => {
          console.log("lectureのgetに失敗しました");
        });
    },
  },
  modules: {},
});
