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
  Select,
  FormLabel,
  Input,
  createStandaloneToast,
} from '@chakra-ui/react';
import Navbar from '../components/Navbar';
import Sidebar from '../components/Sidebar';
import DiscussionCard from '../components/DiscussionCard';
import MainAppLayout from '../components/layout/MainAppLayout';
import {
  API_GET_ALL_QUESTION,
  API_GET_QUESTION_BY_USER_ID,
} from '../api/question';
import useStore from '../provider/zustand/store';
import {
  API_GET_ALL_COURSE,
  API_GET_COURSE_BY_USER_LOGIN,
} from '../api/course';
import { axiosWithToken } from '../api/axiosWithToken';
import { BASE_URL } from '../constant/api';

// Id          int       `json:"id"`
// UserId      int       `json:"user_id"`
// CourseId    int       `json:"course_id"`
// Title       string    `json:"title"`
// Tags        string    `json:"tags"`
// Description string    `json:"description"`
// CreatedAt   time.Time `json:"created_at"`
// UpdatedAt   time.Time `json:"updated_at"`

let discussionList = [
  {
    id: 1,
    user_id: 10,
    course_id: 10,
    title: 'Apa Saja Struktur Lapisan Bumi',
    description: 'description',
    tags: ['tags1'],
    module: 'Geografi',
  },
  // {
  //   id: 2,
  //   question: 'Menentukan Asam Basa',
  //   module: 'Kimia',
  //   class: 'X IPA',
  // },
  // {
  //   id: 3,
  //   question: 'Menentukan Asam Basa',
  //   module: 'Kimia',
  //   class: 'X IPA',
  // },
];

export default function Discussion() {
  const [listDiscusion, setListDiscusion] = useState([]);
  const [listMyDiscusion, setListMyDiscusion] = useState([]);
  const [loadingGetDiscusion, setLoadingGetDiscusion] = useState(false);
  const [loadingGetMyDiscusion, setLoadingGetMyDiscusion] = useState(false);
  const user = useStore((state) => state.user);
  const { toast } = createStandaloneToast();
  const [formDiscussion, setFormDiscussion] = useState({
    title: '',
    tags: [],
    user_id: user.id,
    course_id: 0,
    description: '',
  });
  const [listCourse, setListCourse] = useState([]);
  const [selectedCodeCourse, setSelectedCodeCourse] = useState();

  const getListDiscussion = async () => {
    setLoadingGetDiscusion(true);
    // const res = await API_GET_QUESTION_BY_USER_ID(user.id);
    const res = await API_GET_ALL_QUESTION();
    if (res.status === 200) {
      let data = [];
      for (const discussion of res.data.data ?? []) {
        data.push({
          id: discussion.id,
          title: discussion.title,
          module: discussion.course_name,
          class: discussion.course_class,
          tags: discussion.tags ?? '',
          description: discussion.description ?? '',
        });
      }
      setListDiscusion(data);
    }
    setLoadingGetDiscusion(false);
  };

  const getMyQuestion = async () => {
    setLoadingGetMyDiscusion(true);
    const res = await API_GET_QUESTION_BY_USER_ID(user.id);
    if (res.status === 200) {
      let data = [];
      for (const discussion of res.data.data ?? []) {
        data.push({
          id: discussion.id,
          title: discussion.title,
          module: discussion.course_name,
          class: discussion.course_class,
          tags: discussion.tags ?? '',
          description: discussion.description ?? '',
        });
        console.log(discussion);
      }
      setListMyDiscusion(data);
    }
    setLoadingGetMyDiscusion(false);
  };

  const getListCourse = async () => {
    const res = await API_GET_COURSE_BY_USER_LOGIN();
    if (res.status === 200) {
      const data = res.data.data ?? [];
      let result = [];
      for (const course of data) {
        console.log('course', course);
        result.push({
          label: `${course.course_name} - ${course.course_class}`,
          value: course.id_course,
        });
      }
      setListCourse(result);
    }
  };

  const onChangeCourse = (e) => {
    setSelectedCodeCourse(e.target.value);
    setFormDiscussion({
      ...formDiscussion,
      course_id: Number(e.target.value),
    });
  };

  const onChange = (e) => {
    setFormDiscussion({
      ...formDiscussion,
      [e.target.name]: e.target.value,
    });
  };

  const submit = async (e) => {
    e.preventDefault();
    axiosWithToken()
      .post(`${BASE_URL}/api/questions/create`, formDiscussion)
      .then((res) => {
        if (res.status === 200) {
          toast({
            status: 'success',
            title: 'Berhasil',
            description: 'Berhasil buat question',
          });
          getListDiscussion();
          getMyQuestion();
        } else {
          toast({
            status: 'error',
            title: 'Gagal',
            description: 'Gagal buat question',
          });
        }
      });
  };

  useEffect(() => {
    getListCourse();
    getListDiscussion();
    getMyQuestion();
  }, []);

  return (
    <MainAppLayout>
      <Box m={5}>
        <Stack spacing={6}>
          {/* Header */}
          <Box>
            <Box as="h1" fontSize="2xl" fontWeight="semibold">
              Diskusi
            </Box>
            <Box as="span" fontSize="l" fontWeight="semibold" color="grey">
              Tanyakan Dan Temukan Jawaban pertanyaan Anda Bersama Teman yang
              Lain
            </Box>
          </Box>
          {/* End Header */}
          {/* Content */}
          <Box alignContent="flex-start">
            <VStack spacing={8}>
              {loadingGetDiscusion ? (
                <Spinner />
              ) : listDiscusion.length > 0 ? (
                listDiscusion.map((discussion, index) => {
                  return (
                    <DiscussionCard
                      key={index}
                      id={discussion.id}
                      title={discussion.title}
                      module={discussion.module}
                      moduleClass={discussion.class}
                      tags={discussion.tags}
                      description={discussion.description}
                    />
                  );
                })
              ) : (
                <Text>Belum Ada Discussion</Text>
              )}
            </VStack>
          </Box>
        </Stack>
        <Stack spacing={6} mt="6">
          {/* Header */}
          <Box>
            <Box as="h1" fontSize="2xl" fontWeight="semibold">
              Pertanyaan Saya
            </Box>
          </Box>
          {/* End Header */}
          {/* Content */}
          <Box alignContent="flex-start">
            <VStack spacing={8}>
              {loadingGetMyDiscusion ? (
                <Spinner />
              ) : listMyDiscusion.length > 0 ? (
                listMyDiscusion.map((discussion, index) => {
                  return (
                    <DiscussionCard
                      key={index}
                      id={discussion.id}
                      title={discussion.title}
                      module={discussion.module}
                      moduleClass={discussion.class}
                      tags={discussion.tags}
                      description={discussion.description}
                    />
                  );
                })
              ) : (
                <Text>Belum Ada Discussion Yang Kamu Buat</Text>
              )}
            </VStack>
          </Box>
        </Stack>
        <Box mt="8">
          <Text fontSize="xx-large">Buat Diskusi</Text>
          <Box mt="6">
            <Text>Pilih Course</Text>
            <Select
              id="course"
              name="course"
              placeholder="Select Course"
              value={selectedCodeCourse}
              onChange={onChangeCourse}
              required
            >
              {listCourse.map((course, index) => (
                <option value={course.value} key={index}>
                  {course.label}
                </option>
              ))}
            </Select>
            <Box my="2">
              <FormLabel htmlFor="full-name" fontWeight="bold">
                Title
              </FormLabel>
              <Input
                id="title"
                name="title"
                type="text"
                maxWidth="full"
                height={50}
                placeholder="Title"
                value={formDiscussion.title}
                onChange={onChange}
              />
            </Box>
            <Box my="2">
              <FormLabel htmlFor="full-name" fontWeight="bold">
                Description
              </FormLabel>
              <Input
                id="description"
                name="description"
                type="text"
                maxWidth="full"
                height={50}
                placeholder="Description"
                value={formDiscussion.description}
                onChange={onChange}
              />
            </Box>
            <Box my="2">
              <FormLabel htmlFor="full-name" fontWeight="bold">
                Tags (pisah dengan koma)
              </FormLabel>
              <Input
                id="tags"
                name="tags"
                type="text"
                maxWidth="full"
                height={50}
                placeholder="tags"
                value={formDiscussion.tags}
                onChange={onChange}
              />
            </Box>
            <Button onClick={submit} mt="10" colorScheme="green">
              Buat Pertanyaan
            </Button>
          </Box>
        </Box>
        {/* End main */}
      </Box>
    </MainAppLayout>
  );
}
