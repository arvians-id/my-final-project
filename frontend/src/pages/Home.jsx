import React from 'react';
import { Box, Flex } from '@chakra-ui/react';
import Navbar from '../components/Navbar';
import Sidebar from '../components/Sidebar';
export default function Home() {
    return (
        <div>
            <Navbar />

            <Flex direction="row" justifyContent="center" alignItems="center">
                <Flex
                    width="20%"
                    minHeight="90vh"
                    bgColor="grey.100"
                    boxShadow="dark-lg"
                >
                    <Sidebar />
                </Flex>

                <Flex width="80%" minHeight="90vh" bg="white">
                    <Box m={5}>
                        <Box as="h1" fontSize="xl" fontWeight="semibold" mb={2}>
                            Home
                        </Box>
                    </Box>
                </Flex>
            </Flex>
        </div>
    );
}
