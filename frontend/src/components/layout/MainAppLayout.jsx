import { Flex } from '@chakra-ui/layout';
import React from 'react';
import Navbar from '../Navbar';
import Sidebar from '../Sidebar';

export default function MainAppLayout({ children }) {
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
          {children}
        </Flex>
        {/* Main */}
      </Flex>
    </>
  );
}
