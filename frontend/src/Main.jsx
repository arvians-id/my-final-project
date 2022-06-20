import React from 'react'
import { Route, Routes } from 'react-router-dom'
import Home from './pages/Home'
import Login from './pages/Login'
import Course from './pages/Course'

export default function Main() {
    return (
        <Routes>
            <Route path="/login" element={<Login />} />
            <Route path="/" element={<Home />} />
            {/* <Route path="/course" element={<Course/>} /> */}
        </Routes>
    )
}
