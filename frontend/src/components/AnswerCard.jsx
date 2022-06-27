import React from 'react';
import {
  Box,
  Flex,
  Stack,
  HStack,
  Text,
  Spacer,
  Button,
  Avatar,
  Input,
} from '@chakra-ui/react';

export default function AnswerCard({ answer }) {
  return (
    <>
      <Box
        variant="outline"
        bgColor="white.200"
        p={4}
        width="1500px"
        height="md"
        borderRadius="10"
      >
        <Flex direction="column" alignContent="center" mt={20} width="full">
          <HStack>
            <Avatar
              name="Irfan Kurniawan"
              src="https://bit.ly/dan-abramov"
              mr={4}
              w={10}
              h={10}
            />
            <Input
              size="lg"
              maxWidth="auto"
              variant="flushed"
              placeholder="Jawab"
            />
            <Spacer />
            <Button>Insert</Button>
            <Button>Cancel</Button>
          </HStack>
        </Flex>
        <Box
          mt={7}
          bgColor="blue.100"
          p={4}
          width="1500px"
          height="auto"
          borderRadius="10"
        >
          <HStack>
            <Avatar
              name="Irfan Kurniawan"
              src="https://bit.ly/dan-abramov"
              mr={4}
              w={10}
              h={10}
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
          <Stack mt={5} direction="row">
            <Text
              as="span"
              fontSize="xl"
              fontWeight="semibold"
              color="blackAlpha.600"
            >
              {answer}
            </Text>
          </Stack>
        </Box>
      </Box>
    </>
  );
}
