import React from 'react';
import { Text, Box } from '@chakra-ui/react';
export default function TaskCard({ task }) {
  return (
    <>
      <Box
        m={5}
        flexDirection="row"
        alignContent="center"
        bgColor="blue.400"
        p={4}
        width={80}
        height={80}
        borderRadius="10"
      >
        <Box mb={4} color="white" as="h1" fontSize="2xl" fontWeight="semibold">
          Tugas
        </Box>
        <Text
          as="span"
          fontsize="md"
          fontWeight="semibold"
          color="white"
          noOfLines={[1, 2, 3]}
        >
          {task}
        </Text>
      </Box>
    </>
  );
}
