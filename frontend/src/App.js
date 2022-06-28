import { useEffect, useState } from 'react';
import { Navigate, Route, Routes } from 'react-router-dom';
import { API_CHECK_STATUS } from './api/auth';
import Main from './Main';
import Login from './pages/Login';
import Register from './pages/Register';
import useStore from './provider/zustand/store';
import { adapterUserToFE } from './utils/adapterToFE';
import Submission from './pages/Submission';
import Home from './pages/Home';
import CoursePage from './pages/Course';
import Discussion from './pages/Discussion';
import Profile from './pages/Profile';
import DashboardDataSiswa from './pages/DashboardDataSiswa';
import HomeNonSiswa from './pages/HomeNonSiswa';
import DashboardDataPengguna from './pages/DashboardDataPengguna';
import EditProfile from './pages/EditProfile';
import CourseDetail from './pages/CourseDetail';
// import Answer from './pages/Answer';
import AdminUserList from './pages/Admin/AdminUserList';
import AdminAddCourseModule from './pages/Admin/AdminAddCourseModule';
import AdminAddCourseStudent from './pages/Admin/AdminAddCourseStudent';
import AdminCourseDetail from './pages/Admin/AdminCourseDetail';
import AdminCourseList from './pages/Admin/AdminCourseList';
import AdminCourseSubmissionList from './pages/Admin/AdminCourseSubmissionList';
import DetailDiscussion from './pages/DetailDiscussion';

function App() {
  const user = useStore((state) => state.user);
  const setUser = useStore((state) => state.setUser);
  const [isReady, setIsReady] = useState(false);

  const checkLogin = async () => {
    const res = await API_CHECK_STATUS();
    // kalo oke, berarti set user di zustand
    if (res.status === 200) {
      setUser(adapterUserToFE(res.data.data));
    }
    setIsReady(true);
  };
  // API_CHECK_STATUS
  useEffect(() => {
    checkLogin();
  }, []);

  if (!isReady) return null;

  return (
    <>
      {user !== undefined ? (
        user.role === "Siswa" ?
          <Routes>
            <Route
              path="login"
              element={<Navigate to="/" replace />}
            />
            <Route
              path="register"
              element={<Navigate to="/" replace />}
            />
            <Route path="home" element={<Navigate to="/" replace />} />
            <Route path="course" element={<CoursePage replace />} />
            <Route path="submission" element={<Submission replace />} />
            <Route path="discussion" element={<Discussion replace />} />
            <Route path="profile" element={<Profile replace />} />
            <Route path="edit-profile" element={<EditProfile replace />} />
            <Route path="course/:courseCode" element={<CourseDetail replace />} />
            {/* <Route path="answer" element={<Answer replace />} /> */}
            <Route path="discussion/:questionId" element={<DetailDiscussion replace />} />

            <Route path="/" element={<Home replace />} />
          </Routes> : user.role === "Guru" ? <Routes>
            <Route
              path="login"
              element={<Navigate to="/" replace />}
            />
            <Route
              path="register"
              element={<Navigate to="/" replace />}
            />
            <Route path="/" element={<HomeNonSiswa replace />} />
            <Route path="dashboard-siswa" element={<DashboardDataSiswa replace />} />
            <Route path="add-course" element={<AdminAddCourseModule replace />} />
            <Route path="add-student-to-course" element={<AdminAddCourseStudent replace />} />
            <Route path="admin-course-detail" element={<AdminCourseDetail replace />} />
            <Route path="dashbord-course" element={<AdminCourseList replace />} />
            <Route path="dashboard-submission" element={<AdminCourseSubmissionList replace />} />
            <Route path="dashboard-user" element={<AdminUserList replace />} />
            <Route path="profile" element={<Profile replace />} />
            <Route path="edit-profile" element={<EditProfile replace />} />
          </Routes> : <Routes>
            <Route
              path="login"
              element={<Navigate to="/" replace />}
            />
            <Route
              path="/register"
              element={<Navigate to="/" replace />}
            />
            <Route path="/" element={<HomeNonSiswa replace />} />
            <Route path="add-course" element={<AdminAddCourseModule replace />} />
            <Route path="add-student-to-course" element={<AdminAddCourseStudent replace />} />
            <Route path="admin-course-detail" element={<AdminCourseDetail replace />} />
            <Route path="dashbord-course" element={<AdminCourseList replace />} />
            <Route path="dashboard-submission" element={<AdminCourseSubmissionList replace />} />
            <Route path="profile" element={<Profile replace />} />
            <Route path="edit-profile" element={<EditProfile replace />} />
          </Routes>
      ) : (
        <Routes>
          <Route path="*" element={<Login />} />
          <Route path="login" element={<Login />} />
          <Route path="register" element={<Register />} />
          {/* <Route path="/" element={<Login />} /> */}
        </Routes>
      )}
    </>
  );
}

export default App;
