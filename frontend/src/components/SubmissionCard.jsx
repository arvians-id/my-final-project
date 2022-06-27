import React from 'react';
import { Flex, Text, Spacer, Badge, Button } from '@chakra-ui/react';
export default function SubmissionCard({ name, status }) {
  return (
    <Flex
      flexDirection="row"
      alignContent="center"
      bgColor="blue.500"
      p={4}
      width="full"
      height={14}
      borderRadius="10"
    >
      <Text as="span" fontsize="md" fontWeight="semibold" color="white">
        {name}
      </Text>
      <Spacer />
      {status ? (
        <Badge colorScheme="green" px={3} py={1} borderRadius={5}>
          Sudah Terkumpul
        </Badge>
      ) : (
        <Badge colorScheme="red" px={3} py={1} borderRadius={5}>
          Belum Terkumpul
        </Badge>
      )}
      <Button ml={4} size="xs" variant="solid">
        Kumpul Disini
      </Button>
    </Flex>
  );
}
