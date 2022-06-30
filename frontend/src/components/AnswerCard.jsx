import React, { useEffect, useState } from 'react';
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
  VStack,
} from '@chakra-ui/react';
import { API_GET_USER_DETAIL_BY_ID } from '../api/user';

export default function AnswerCard({ answer }) {
  const [siswa, setSiswa] = useState();

  const getDetailSiswa = async () => {
    const res = await API_GET_USER_DETAIL_BY_ID(answer.user_id);
    if (res.status === 200) {
      setSiswa(res.data.data);
    }
  };

  useEffect(() => {
    getDetailSiswa();
  }, []);
  return (
    <>
      <Box
        variant="outline"
        w="full"
        bgColor="white.200"
        p={4}
        borderRadius="10"
      >
        {/* <Flex direction="column" alignContent="center" mt={20} width="full">
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
              value={value}
              onChange={onChange}
            />
            <Spacer />
            <Button onClick={submit}>Insert</Button>
            <Button>Cancel</Button>
          </HStack>
        </Flex> */}
        <Box bgColor="blue.100" p={4} height="auto" borderRadius="10">
          <HStack>
            <Avatar
              name={siswa?.username}
              src="https://bit.ly/dan-abramov"
              mr={4}
              w={10}
              h={10}
            />
            <Stack>
              <Text as="span" fontSize="lg" fontWeight="semibold">
                {siswa?.username}
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
          <VStack mt={4} alignItems="flex-start">
            <Text as="span" fontSize="16px" fontWeight="semibold" color="black">
              Jawab:
            </Text>
            <Text
              mt="0"
              as="span"
              fontSize="xl"
              fontWeight="semibold"
              color="blackAlpha.600"
            >
              {answer.description}
            </Text>
          </VStack>
        </Box>
      </Box>
    </>
  );
}
