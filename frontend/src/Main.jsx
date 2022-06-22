import { Flex } from '@chakra-ui/layout';
import React from 'react';
import { Route, Routes } from 'react-router-dom';
import Navbar from './components/Navbar';
import Sidebar from './components/Sidebar';
import Home from './pages/Home';
import CoursePage from './pages/Course';
import Submission from './pages/Submission';

export default function Main() {
    return (
        <>
            <Navbar />
            <Flex
                direction="row"
                justifyContent="flex-start"
                alignItems="flex-start"
                top="30"
            >
                {/* Sidebar */}
                <Flex
                    width="20%"
                    minHeight="100vh"
                    bgColor="grey.100"
                    boxShadow="md"
                    position="fixed"
                    left="0"
                    top="20"
                    overflowY="auto"
                >
                    <Sidebar />
                </Flex>
                {/* End Sidebar */}
                {/* Main */}
                <Flex
                    width="80%"
                    minHeight="90vh"
                    bg="white"
                    position="sticky"
                    left="80"
                    marginTop={20}
                >
                    <Routes>
                        {/* <Route path="/course" element={<CoursePage />} /> */}
                        <Route path="submission/*" element={<Submission />} />
                        {/* <Route path="/" element={<Home />} /> */}
                    </Routes>
                </Flex>
                {/* Main */}
            </Flex>
        </>
    );
}
