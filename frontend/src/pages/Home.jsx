import React, { useEffect, useState } from 'react';
import {
  Box,
  Flex,
  Stack,
  HStack,
  Text,
  Spacer,
  Button,
  Spinner,
} from '@chakra-ui/react';
import Navbar from '../components/Navbar';
import Sidebar from '../components/Sidebar';
import CourseCard from '../components/CourseCard';
import SubmissionCard from '../components/SubmissionCard';
import DiscussionCard from '../components/DiscussionCard';
import MainAppLayout from '../components/layout/MainAppLayout';
import { API_GET_COURSE_BY_USER_LOGIN } from '../api/course';
import { Link } from 'react-router-dom';
import { API_GET_SUBMISSION_BY_USER_LOGIN } from '../api/submission';
import { API_GET_QUESTION_BY_USER_ID } from '../api/question';
import useStore from '../provider/zustand/store';
import { getStatusSubmision } from '../utils/submission';

export default function Home() {
  let courseList = [
    {
      courseTitle: 'Geografi',
      courseClass: 'XI IPS',
      courseTeacher: 'Rahmat Pratama, S.Pd',
      courseDescription: 'Mempelajari tentang Struktur Bumi',
    },
    {
      courseTitle: 'Bahasa Indonesia',
      courseClass: 'X Bahasa',
      courseTeacher: 'Isna Rahmawati, S.Pd',
      courseDescription: 'Bahasa Indonesia Pelajaran Mengenai Bahasa Indonesia',
    },
    {
      courseTitle: 'Matematika',
      courseClass: 'XII IPA',
      courseTeacher: 'Suci Rahma, S.Pd',
      courseDescription:
        'Mata Pelajaran yang akan membahas Konversi Biner, Aljabar ',
    },
  ];

  let moduleList = [
    {
      id: 1,
      name: 'Pemrograman Web',
      class: 'XI RPL',
      description: 'tentang Web Programming',
      percent: 70,
    },
    {
      id: 2,
      name: 'Bahasa Indonesia',
      class: 'X TKJ',
      description: 'Bahasa Indonesia Pelajaran Mengenai Bahasa Indonesia',
      percent: 85,
    },
    {
      id: 3,
      name: 'Matematika',
      class: 'XII TKJ',
      description: 'Mata Pelajaran yang akan membahas Konversi Biner, Aljabar ',
      percent: 60,
    },
  ];
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

  let discussionList = [
    {
      id: 1,
      title: 'Apa Saja Struktur Bumi',
      module: 'Geografi',
      class: 'X IPS',
    },
    {
      id: 2,
      title: 'Menentukan Asam Basa',
      module: 'Kimia',
      class: 'X IPA',
    },
  ];
  const [listCourse, setListCourse] = useState([]);
  const [loadingGetCourse, setLoadingGetCourse] = useState(false);
  const [listSubmission, setListSubmission] = useState([]);
  const [loadingGetSubmision, setLoadingGetSubmision] = useState(false);
  const [listDiscusion, setListDiscusion] = useState([]);
  const [loadingGetDiscusion, setLoadingGetDiscusion] = useState(false);
  const user = useStore((state) => state.user);

  const getListCourse = async () => {
    setLoadingGetCourse(true);
    const res = await API_GET_COURSE_BY_USER_LOGIN();
    if (res.status === 200) {
      setListCourse(res.data.data ?? []);
    }
    setLoadingGetCourse(false);
  };

  const getListSubmission = async () => {
    setLoadingGetSubmision(true);
    const res = await API_GET_SUBMISSION_BY_USER_LOGIN('?limit=3');
    if (res.status === 200) {
      setListSubmission(res.data.data ?? []);
    }
    setLoadingGetSubmision(false);
  };

  const getListDiscussion = async () => {
    setLoadingGetDiscusion(true);
    const res = await API_GET_QUESTION_BY_USER_ID(user.id);
    if (res.status === 200) {
      let data = [];
      for (const discussion of res.data.data ?? []) {
        data.push({
          id: discussion.id,
          title: discussion.title,
          module: discussion.course_name,
          class: discussion.course_class,
          description: discussion.description,
          tags: discussion.tags,
        });
      }
      setListDiscusion(data);
    }
    setLoadingGetDiscusion(false);
  };

  useEffect(() => {
    getListCourse();
    getListSubmission();
    getListDiscussion();
  }, []);

  return (
    <MainAppLayout>
      <Box m={5}>
        <Stack spacing={6}>
          {/* Header */}
          <Box>
            <Box as="h1" fontSize="2xl" fontWeight="semibold">
              Selamat Datang Kembali
            </Box>
            <Box as="span" fontSize="l" fontWeight="semibold" color="grey">
              Lanjutkan Pembelajaran Anda
            </Box>
          </Box>
          {/* End Header */}
          <Flex justifyContent="baseline">
            <Box as="h1" fontSize="2xl" fontWeight="semibold">
              Pelajaran Anda
            </Box>
            <Spacer />
            <Link to="/course">
              <Button colorScheme="blue" size="md" variant="ghost">
                Lihat Semua
              </Button>
            </Link>
          </Flex>
          <Box alignContent="flex-start">
            <HStack spacing={8}>
              {loadingGetCourse ? (
                <Spinner />
              ) : listCourse.length > 0 ? (
                listCourse.slice(0, 3).map((course, index) => {
                  return (
                    <CourseCard key={index} course_code={course.course_code} />
                  );
                })
              ) : (
                <Text>Belum Ada Course Yang Kamu Masuki</Text>
              )}
            </HStack>
          </Box>
          <HStack spacing={3}>
            <Box
              p={5}
              width="50%"
              bgColor="white"
              height="60vh"
              boxShadow="lg"
              borderRadius="10"
            >
              <Flex>
                <Text as="span" fontSize="xl" fontWeight="semibold">
                  Tugas
                </Text>
                <Spacer />
                <Link to="/submission">
                  <Button colorScheme="blue" size="md" variant="ghost">
                    Lihat Semua
                  </Button>
                </Link>
              </Flex>
              <Stack mt="4" spacing={3}>
                {loadingGetSubmision ? (
                  <Spinner />
                ) : listSubmission.length > 0 ? (
                  listSubmission.map((submission, index) => {
                    return (
                      <SubmissionCard
                        key={index}
                        name={submission.name_course}
                        status={getStatusSubmision(submission)}
                      />
                    );
                  })
                ) : (
                  <Text>Belum Ada Tugas Untuk Kamu</Text>
                )}
              </Stack>
            </Box>
            <Box
              p={5}
              width="50%"
              bgColor="gray.100"
              height="60vh"
              boxShadow="lg"
              borderRadius="10"
            >
              <Flex>
                <Text as="span" fontSize="xl" fontWeight="semibold">
                  Diskusi
                </Text>
                <Spacer />
                <Link to="/discussion">
                  <Button colorScheme="blue" size="md" variant="ghost">
                    Lihat Semua
                  </Button>
                </Link>
              </Flex>
              <Stack mt={4} spacing={3}>
                {loadingGetDiscusion ? (
                  <Spinner />
                ) : listDiscusion.length > 0 ? (
                  listDiscusion.slice(0, 3).map((discussion, index) => {
                    return (
                      <DiscussionCard
                        key={index}
                        id={discussion.id}
                        title={discussion.title}
                        module={discussion.module}
                        moduleClass={discussion.class}
                        description={discussion.description}
                        tags={discussion.tags}
                      />
                    );
                  })
                ) : (
                  <Text>Belum Ada Discussion Yang Kamu Buat</Text>
                )}
              </Stack>
            </Box>
          </HStack>
        </Stack>
      </Box>
    </MainAppLayout>
  );
}
