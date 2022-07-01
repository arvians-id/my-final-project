import React, { useEffect, useState } from 'react';
import {
  Box,
  Flex,
  Stack,
  VStack,
  Text,
  Spacer,
  Button,
  Spinner,
} from '@chakra-ui/react';
import Navbar from '../components/Navbar';
import Sidebar from '../components/Sidebar';
// import CourseCard from '../components/CourseCard'
import SubmissionCard from '../components/SubmissionCard';
import MainAppLayout from '../components/layout/MainAppLayout';
import { API_GET_SUBMISSION_BY_USER_LOGIN } from '../api/submission';
import { getStatusSubmision } from '../utils/submission';
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
  const [listSubmission, setListSubmission] = useState([]);
  const [loadingGetSubmision, setLoadingGetSubmision] = useState(false);

  const getListSubmission = async () => {
    setLoadingGetSubmision(true);
    const res = await API_GET_SUBMISSION_BY_USER_LOGIN('');
    if (res.status === 200) {
      setListSubmission(res.data.data ?? []);
    }
    setLoadingGetSubmision(false);
  };

  useEffect(() => {
    getListSubmission();
  }, []);

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
              {loadingGetSubmision ? (
                <Spinner />
              ) : listSubmission.length > 0 ? (
                <VStack spacing={8}>
                  {listSubmission.map((submission, index) => {
                    return (
                      <SubmissionCard
                        key={index}
                        name={submission.name_course}
                        status={getStatusSubmision(submission)}
                        type="submit"
                        submissionId={submission.id_module_submission}
                        courseCode={submission.course_code}
                        getListSubmission={getListSubmission}
                      />
                    );
                  })}
                </VStack>
              ) : (
                <Text>Belum Ada Tugas Untuk Kamu</Text>
              )}
            </Box>
            {/* End Content */}
          </Stack>
        </Box>
      </Flex>
    </MainAppLayout>
  );
}
