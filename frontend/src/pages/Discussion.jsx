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
import DiscussionCard from '../components/DiscussionCard';
import MainAppLayout from '../components/layout/MainAppLayout';
import { API_GET_QUESTION_BY_USER_ID } from '../api/question';
import useStore from '../provider/zustand/store';

let discussionList = [
  {
    id: 1,
    question: 'Apa Saja Struktur Lapisan Bumi',
    module: 'Geografi',
    class: 'X IPS',
    num: 'Jawaban 1',
    answer:
      'Crust (Kerak Bumi), Mantle (Mantel Bumi), Outer Core (Inti Luar), Inner Core / Inti Dalam',
  },
  {
    id: 2,
    question: 'Menentukan Asam Basa',
    module: 'Kimia',
    class: 'X IPA',
  },
  {
    id: 3,
    question: 'Menentukan Asam Basa',
    module: 'Kimia',
    class: 'X IPA',
  },
];

export default function Discussion() {
  const [listDiscusion, setListDiscusion] = useState([]);
  const [loadingGetDiscusion, setLoadingGetDiscusion] = useState(false);
  const user = useStore((state) => state.user);

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
        });
      }
      setListDiscusion(data);
    }
    setLoadingGetDiscusion(false);
  };

  useEffect(() => {
    getListDiscussion();
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
                    />
                  );
                })
              ) : (
                <Text>Belum Ada Discussion Yang Kamu Buat</Text>
              )}
            </VStack>
          </Box>
          </Stack>
        {/* End main */}
      </Box>
    </MainAppLayout>
  );
}
