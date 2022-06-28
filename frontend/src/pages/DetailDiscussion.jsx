import React, { useEffect, useState } from 'react';
import {
  Box,
  Button,
  Flex,
  FormLabel,
  HStack,
  Spacer,
  Stack,
  Textarea,
  VStack,
} from '@chakra-ui/react';
import QuestionCard from '../components/QuestionCard';
import AnswerCard from '../components/AnswerCard';
import MainAppLayout from '../components/layout/MainAppLayout';
import { axiosWithToken } from '../api/axiosWithToken';
import { BASE_URL } from '../constant/api';
import { useParams } from 'react-router';
import useStore from '../provider/zustand/store';

let question = [
  {
    id: 1,
    user_id: 10,
    course_id: 10,
    title: 'Apa Saja Struktur Lapisan Bumi',
    description: 'description',
    tags: ['tags1'],
    // module: 'Geografi',
  },
];

let answerList = [
  {
    id: 1,
    answer:
      '1. Crust (Kerak Bumi) Crust merupakan bagian terluar dari lapisan Bumi yang lebih tipis dibandingkan dengan lapisan lainnya. 2. Mantle (Mantel Bumi) Lapisan Bumi kedua adalah mantel yang merupakan lapisan paling tebal dengan ketebalan mencapai 2.900 km. lapisan ini juga disebut lapisan astenosfer karena fungsinya yaitu untuk melindungi inti Bumi. 3 .Outer Core (Inti Luar) Lapisan Bumi ini merupakan lapisan cair dengan ketebalan sekitar 2266 km yang terdiri dari besi dan nikel di atas inti dalam dan di bawah mantel. 4. Inner Core / Inti DalamSesuai dengan namanya, inti dalam merupakan lapisan Bumi paling dalam yang berbentuk bola padat berjari-jari sekitar 1.220 km.',
  },
];

export default function DetailDiscussion() {
  let { questionId } = useParams();
  const [questions, setQuestions] = useState();
  const [listAnswer, setListAnswer] = useState([]);
  const user = useStore((state) => state.user);

  const [formAnswer, setFormAnswer] = useState({
    question_id: Number(questionId),
    user_id: user.id,
    description: '',
  });

  const change = (e) => {
    setFormAnswer({
      ...formAnswer,
      description: e.target.value,
    });
  };

  const getQuestionDetail = async (questionId) => {
    axiosWithToken()
      .get(`${BASE_URL}/api/questions/${questionId}`)
      .then((res) => {
        if (res.status === 200) {
          setQuestions(res.data.data);
        }
      });
  };
  const getAnswerBuQuestion = async (questionId) => {
    axiosWithToken()
      .get(`${BASE_URL}/api/answers/${questionId}`)
      .then((res) => {
        if (res.status === 200) {
          setListAnswer(res.data.data ?? []);
        }
      });
  };

  const submit = async (e) => {
    e.preventDefault();
    axiosWithToken()
      .post(`${BASE_URL}/api/answers/create`, formAnswer)
      .then((res) => {
        if (res.status === 200) {
          getAnswerBuQuestion(questionId);
        }
      });
  };

  useEffect(() => {
    if (questionId) {
      getQuestionDetail(questionId);
      getAnswerBuQuestion(questionId);
    }
  }, [questionId]);

  return (
    <MainAppLayout>
      {/* Main */}
      <Flex
        direction="column"
        width="80%"
        minHeight="90vh"
        bg="white"
        position="sticky"
        left="80"
        marginTop={20}
        py="3"
      >
        <Box m={5}>
          <Stack spacing={6}>
            {/* Header */}
            <Box>
              <Box as="h1" fontSize="2xl" fontWeight="semibold">
                Jawaban
              </Box>
              <Box as="span" fontSize="l" fontWeight="semibold" color="grey">
                Tanyakan Dan Temukan Jawaban pertanyaan Anda Bersama Teman yang
                Lain
              </Box>
            </Box>
            <Box alignContent="flex-start">
              {questions && (
                <QuestionCard
                  question={questions.title}
                  module={questions.course_name}
                  moduleClass={questions.course_class}
                />
              )}
            </Box>
            <Box alignContent="flex-start" mb="10">
              <VStack spacing={8}>
                {listAnswer.map((answer, index) => {
                  return (
                    <AnswerCard
                      key={index}
                      answer={answer.answer}
                      value={formAnswer.description}
                      onChange={change}
                      submit={submit}
                    />
                  );
                })}
              </VStack>
            </Box>
            {/* End Content */}
          </Stack>
        </Box>
        {/* <HStack>
          <Textarea
            id="description"
            name="description"
            type="text"
            maxWidth="full"
            height={50}
            placeholder="Description"
            value={formAnswer.description}
            onChange={change}
          />
          <Spacer />
          <Button onClick={submit} mt="10" colorScheme="green">
            Jawab
          </Button>
          <Button>Cancel</Button>
        </HStack> */}
        {/* <Box my="2">
          <FormLabel htmlFor="full-name" fontWeight="bold">
            Description
          </FormLabel>
          <Textarea
            id="description"
            name="description"
            type="text"
            maxWidth="full"
            height={50}
            placeholder="Description"
            value={formAnswer.description}
            onChange={change}
          />
        </Box> */}
      </Flex>
      {/* End main */}
    </MainAppLayout>
  );
}
