import React from 'react';
import { Flex, Stack, HStack, Text, Avatar, Badge } from '@chakra-ui/react';
import useStore from '../provider/zustand/store';

export default function QuestionCard({
  question,
  module,
  moduleClass,
  description,
  tags,
  userName,
}) {
  const user = useStore((state) => state.user);
  console.log(userName);
  return (
    <>
      <Flex
        bgColor="blue.200"
        p={4}
        // width="700px"
        borderRadius="10"
      >
        <Flex direction="column" mt={2} w="full">
          <HStack>
            <Avatar name={userName} src="/user.png" mr={2} w={14} h={14} />
            <Stack>
              <Text as="span" fontSize="lg" fontWeight="semibold">
                {userName}
              </Text>
              <Text
                as="span"
                fontSize="md"
                align="left"
                fontWeight="semibold"
                color="black"
              >
                Siswa
              </Text>
            </Stack>
          </HStack>
          <Text
            mt={5}
            as="h1"
            fontSize="2xl"
            fontWeight="semibold"
            color="white"
          >
            Judul: {question}
          </Text>
          <Text
            mt={2}
            as="h1"
            fontSize="2xl"
            fontWeight="semibold"
            maxW="450px"
          >
            Pertayaan: {description}
          </Text>
          <HStack alignItems="flex-start">
            {tags.split(',').map((tag, index) => (
              <Badge key={index}>#{tag}</Badge>
            ))}
          </HStack>
          <Stack mt={5} direction="row">
            <Text as="span" fontSize="xl" fontWeight="semibold" color="black">
              {`${module} - ${moduleClass}`}
            </Text>
          </Stack>
        </Flex>
      </Flex>
    </>
  );
}
