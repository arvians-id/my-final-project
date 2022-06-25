import React from 'react';
import { Flex, Stack, HStack, Text, Avatar } from '@chakra-ui/react';

export default function QuestionCard({ question, module, moduleClass }) {
  return (
    <>
      <Flex
        bgColor="blue.200"
        p={4}
        width="1500px"
        height="md"
        borderRadius="10"
      >
        <Flex direction="column" alignContent="center" mt={20}>
          <HStack>
            <Avatar
              name="Irfan Kurniawan"
              src="https://bit.ly/dan-abramov"
              mr={2}
              w={14}
              h={14}
            />
            <Stack>
              <Text as="span" fontSize="lg" fontWeight="semibold">
                Irfan Kurniawan
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
            fontSize="5xl"
            fontWeight="semibold"
            color="white"
          >
            {question}
          </Text>
          <Stack mt={5} direction="row">
            <Text as="span" fontSize="3xl" fontWeight="semibold" color="black">
              {module}
            </Text>
            <Text as="span" fontSize="3xl" fontWeight="semibold" color="black">
              {moduleClass}
            </Text>
          </Stack>
        </Flex>
      </Flex>
    </>
  );
}
