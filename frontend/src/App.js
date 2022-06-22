<<<<<<< HEAD
import Login from "./pages/Login";

function App() {
  return (
    <Login />
  );
=======
import { useEffect, useState } from 'react';
import { Navigate, Route, Routes } from 'react-router-dom';
import { API_CHECK_STATUS } from './api/auth';
import Main from './Main';
import Login from './pages/Login';
import Register from './pages/Register';
import useStore from './provider/zustand/store';
import { adapterUserToFE } from './utils/adapterToFE'
import Submission from './pages/Submission';
import Home from './pages/Home';
import CoursePage from './pages/Course';
import Discussion from './pages/Discussion';
import Profile from './pages/Profile';
import DashboardDataSiswa from './pages/DashboardDataSiswa';
import HomeNonSiswa from './pages/HomeNonSiswa';
import DashboardDataPengguna from './pages/DashboardDataPengguna';

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
        console.log('user', user)
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
                            path="/login"
                            element={<Navigate to="/" replace />}
                        />
                        <Route
                            path="/register"
                            element={<Navigate to="/" replace />}
                        />
                        <Route path="home" element={<Navigate to="/" replace />} />
                        <Route path="course" element={<CoursePage replace />} />
                        <Route path="submission" element={<Submission replace />} />
                        <Route path="discussion" element={<Discussion replace />} />
                        <Route path="profile" element={<Profile replace />} />
                        <Route path="/" element={<Home replace />} />
                    </Routes> : user.role === "Guru" ? <Routes>
                        <Route
                            path="/login"
                            element={<Navigate to="/" replace />}
                        />
                        <Route
                            path="/register"
                            element={<Navigate to="/" replace />}
                        />
                        <Route path="/" element={<HomeNonSiswa replace />} />
                        <Route path="dashboard-siswa" element={<DashboardDataSiswa replace />} />
                        <Route path="profile" element={<Profile replace />} />
                    </Routes> : <Routes>
                        <Route
                            path="/login"
                            element={<Navigate to="/" replace />}
                        />
                        <Route
                            path="/register"
                            element={<Navigate to="/" replace />}
                        />
                        <Route path="/" element={<HomeNonSiswa replace />} />
                        <Route path="dashboard-pengguna" element={<DashboardDataPengguna replace />} />
                        <Route path="profile" element={<Profile replace />} />
                    </Routes>
            ) : (
                <Routes>
                    <Route path="*" element={<Login />} />
                    <Route path="/login" element={<Login />} />
                    <Route path="/register" element={<Register />} />
                    <Route path="/" element={<Login />} />
                </Routes>
            )}
        </>
    );
>>>>>>> 28ee5ed6f3b932b186ee81144b50e15402a23589
}

export default App;
