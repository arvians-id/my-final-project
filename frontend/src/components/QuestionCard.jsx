import React from 'react';
import { Flex, Stack, HStack, Text, Avatar } from '@chakra-ui/react';
import useStore from '../provider/zustand/store';

export default function QuestionCard({ question, module, moduleClass }) {
  const user = useStore((state) => state.user);

  return (
    <>
      <Flex
        bgColor="blue.200"
        p={4}
        // width="700px"
        borderRadius="10"
      >
        <Flex direction="column" mt={20} w="full">
          <HStack>
            <Avatar name={user.username} src="/user.png" mr={2} w={14} h={14} />
            <Stack>
              <Text as="span" fontSize="lg" fontWeight="semibold">
                {user.username}
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
            {question}
          </Text>
          <Stack mt={5} direction="row">
            <Text as="span" fontSize="xl" fontWeight="semibold" color="black">
              {module}
            </Text>
            <Text as="span" fontSize="xl" fontWeight="semibold" color="black">
              {moduleClass}
            </Text>
          </Stack>
        </Flex>
      </Flex>
    </>
  );
}
