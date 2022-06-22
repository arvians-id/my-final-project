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
import DiscussionCard from '../components/DiscussionCard';
import MainAppLayout from '../components/layout/MainAppLayout';

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

export default function Discussion() {
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
              {discussionList.map((discussion, index) => {
                return (
                  <DiscussionCard
                    key={index}
                    title={discussion.title}
                    module={discussion.module}
                    moduleClass={discussion.class}
                  />
                );
              })}
            </VStack>
          </Box>
          {/* End Content */}
        </Stack>
      </Box>
    </MainAppLayout>
  );
}
