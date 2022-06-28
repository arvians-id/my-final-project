import { Heading, HStack, Stack, Text, Box } from '@chakra-ui/react';
import React from 'react';

export default function DetailCard({ name, className, description }) {
  return (
    <Box
      m={4}
      p={6}
      shadow="md"
      borderWidth="1px"
      w={950}
      h={350}
      borderRadius={10}
    >
      <Stack spacing={3}>
        <Heading fontSize="6xl">{name}</Heading>
        <HStack spacing={3} mt={4}>
          {/* <Text fontWeight="semibold" color="grey">{courseTeacher}</Text> */}
          <Text fontWeight="semibold" color="blue.200">
            {className}
          </Text>
        </HStack>
        <Text mt={4} align="justify">
          {description}
        </Text>
      </Stack>
    </Box>
  );
}
