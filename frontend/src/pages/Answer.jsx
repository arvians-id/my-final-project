import React from 'react';
import { Box, Flex, Stack, VStack } from '@chakra-ui/react';
import Navbar from '../components/Navbar';
import Sidebar from '../components/Sidebar';
import QuestionCard from '../components/QuestionCard';
import AnswerCard from '../components/AnswerCard';

let question = [
  {
    id: 1,
    question: 'Apa Saja Struktur Lapisan Bumi',
    module: 'Geografi',
    class: 'X IPS',
  },
];

let answerList = [
  {
    id: 1,
    answer:
      '1. Crust (Kerak Bumi) Crust merupakan bagian terluar dari lapisan Bumi yang lebih tipis dibandingkan dengan lapisan lainnya. 2. Mantle (Mantel Bumi) Lapisan Bumi kedua adalah mantel yang merupakan lapisan paling tebal dengan ketebalan mencapai 2.900 km. lapisan ini juga disebut lapisan astenosfer karena fungsinya yaitu untuk melindungi inti Bumi. 3 .Outer Core (Inti Luar) Lapisan Bumi ini merupakan lapisan cair dengan ketebalan sekitar 2266 km yang terdiri dari besi dan nikel di atas inti dalam dan di bawah mantel. 4. Inner Core / Inti DalamSesuai dengan namanya, inti dalam merupakan lapisan Bumi paling dalam yang berbentuk bola padat berjari-jari sekitar 1.220 km.',
  },
];

export default function Discussion() {
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
                  Jawaban
                </Box>
                <Box as="span" fontSize="l" fontWeight="semibold" color="grey">
                  Tanyakan Dan Temukan Jawaban pertanyaan Anda Bersama Teman
                  yang Lain
                </Box>
              </Box>
              <Box alignContent="flex-start">
                <VStack spacing={8}>
                  {question.map((discussion, index) => {
                    return (
                      <QuestionCard
                        key={index}
                        question={discussion.question}
                        module={discussion.module}
                        moduleClass={discussion.class}
                      />
                    );
                  })}
                </VStack>
              </Box>
              <Box alignContent="flex-start">
                <VStack spacing={8}>
                  {answerList.map((answer, index) => {
                    return <AnswerCard key={index} answer={answer.answer} />;
                  })}
                </VStack>
              </Box>
              {/* End Content */}
            </Stack>
          </Box>
        </Flex>
        {/* End main */}
      </Flex>
    </>
  );
}
