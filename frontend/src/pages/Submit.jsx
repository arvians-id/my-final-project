import React from 'react';
import { Box, Flex, Stack, VStack } from '@chakra-ui/react';
import Navbar from '../components/Navbar';
import Sidebar from '../components/Sidebar';
// import CourseCard from '../components/CourseCard'
import ListSubmission from '../components/ListSubmission';
// import DiscussionCard from '../components/DiscussionCard'
import SubmitCard from '../components/SubmitCard';
import TaskCard from '../components/TaskCard';

let submissionList = [
  {
    id: 1,
    name: 'Matematika 2',
    status: false,
  },
];

let taskList = [
  {
    id: 1,
    task: 'Jika x dan y adalah solusi dari sistem persamaan 4x + y = 9 dan x + 4y = 6, maka nilai 2x + 3y',
  },
];

export default function Submit() {
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
                      <ListSubmission
                        key={index}
                        name={submission.name}
                        status={submission.status}
                      />
                    );
                  })}
                </VStack>
              </Box>
              <Box alignContent="left"></Box>
              {/* End Content */}
            </Stack>
          </Box>
          {taskList.map((soal, index) => {
            return <TaskCard key={index} task={soal.task} />;
          })}
          <SubmitCard />
        </Flex>
        {/* End main */}
      </Flex>
    </>
  );
}
