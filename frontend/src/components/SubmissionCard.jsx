import React from 'react';
import { Flex, Text, Spacer, Badge } from '@chakra-ui/react';
export default function SubmissionCard({ name, status }) {
  return (
    <Flex
      flexDirection="row"
      alignContent="center"
      bgColor="blue.500"
      p={4}
      width="full"
      borderRadius="10"
    >
      <Text as="span" fontsize="md" fontWeight="semibold" color="white">
        {name}
      </Text>
      <Spacer />
      <Badge
        colorScheme={status ? 'green' : 'red'}
        px={3}
        display="flex"
        justifyContent="center"
        alignItems="center"
        py={1}
        h="30px"
        borderRadius={5}
      >
        {status ? 'Sudah Terkumpul' : 'Belum Terkumpul'}
      </Badge>
    </Flex>
  );
}
