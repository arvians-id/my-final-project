import { React } from 'react';
import { Box, Flex } from '@chakra-ui/react';
import Navbar from '../components/Navbar';
import Sidebar from '../components/Sidebar';
import MainAppLayout from '../components/layout/MainAppLayout';
export default function Profile() {
  return (
    <MainAppLayout>
      <Flex width="80%" minHeight="90vh" bg="white">
        <Box m={5}>
          <Box as="h1" fontSize="xl" fontWeight="semibold" mb={2}>
            Profile
          </Box>
        </Box>
      </Flex>
    </MainAppLayout>
  );
}
