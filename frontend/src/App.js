import { useEffect, useState } from 'react';
import { Navigate, Route, Routes } from 'react-router-dom';
import { API_CHECK_LOGIN } from './api/auth';
import Main from './Main';
import Home from './pages/Home';
import Login from './pages/Login';
import Register from './pages/Register';
import useStore from './provider/zustand/store';
import { localClearToken, localLoadToken } from './utils/token';

function App() {
    const user = useStore((state) => state.user);
    const setUser = useStore((state) => state.setUser);
    const [isReady, setIsReady] = useState(false);

    const checkLogin = async () => {
        const token = localLoadToken();
        if (token) {
            const res = await API_CHECK_LOGIN(token);
            // kalo oke, berarti set user di zustand
            if (res.status === 200) {
                setUser(res.data.data);
            } else {
                // buang token kalo token user invalid
                localClearToken();
            }
        }

        setIsReady(true);
    };
    // API_CHECK_LOGIN
    useEffect(() => {
        checkLogin();
    }, []);

    if (!isReady) return null;

    return (
        <>
            {user !== undefined ? (
                <Routes>
                    <Route
                        path="/login"
                        element={<Navigate to="/" replace />}
                    />
                    <Route
                        path="/register"
                        element={<Navigate to="/" replace />}
                    />
                    <Route path="/home" element={<Navigate to="/" replace />} />
                    <Route path="/" element={<Main />} />
                </Routes>
            ) : (
                <Routes>
                    <Route path="/login" element={<Login />} />
                    <Route path="/register" element={<Register />} />
                    <Route path="/home" element={<Home />} />
                    <Route path="/" element={<Login />} />
                </Routes>
            )}
        </>
    );
}

export default App;
