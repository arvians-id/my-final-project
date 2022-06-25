import React from 'react';
import {
  Box,
  Flex,
  Stack,
  VStack,
  Text,
  Spacer,
  Button,
} from '@chakra-ui/react';
import Navbar from '../components/Navbar';
import Sidebar from '../components/Sidebar';
// import CourseCard from '../components/CourseCard'
import SubmissionCard from '../components/SubmissionCard';
import MainAppLayout from '../components/layout/MainAppLayout';
// import DiscussionCard from '../components/DiscussionCard'

let submissionList = [
  {
    id: 1,
    name: 'Matematika 1',
    status: true,
  },
  {
    id: 1,
    name: 'Matematika 2',
    status: false,
  },
];

export default function Submission() {
  return (
    <MainAppLayout>
      <Flex
        direction="column"
        width="80%"
        minHeight="90vh"
        bg="white"
        position="sticky"
        left="80"
        marginTop={20}
      >
        <Box m={5}>
          <Stack spacing={6}>
            {/* Header */}
            <Box>
              <Box as="h1" fontSize="2xl" fontWeight="semibold">
                Tugas
              </Box>
              <Box as="span" fontSize="l" fontWeight="semibold" color="grey">
                Kerjakan Tugas anda dan dapatkan Nilai Terbaik dari Guru Anda
              </Box>
            </Box>
            {/* End Header */}
            {/* Content */}
            <Box alignContent="flex-start">
              <VStack spacing={8}>
                {submissionList.map((submission, index) => {
                  return (
                    <SubmissionCard
                      key={index}
                      name={submission.name}
                      status={submission.status}
                    />
                  );
                })}
              </VStack>
            </Box>
            {/* End Content */}
          </Stack>
        </Box>
      </Flex>
    </MainAppLayout>
  );
}
