import React from 'react';
import { Flex, Stack, Text, Spacer, Button } from '@chakra-ui/react';
import { Link } from 'react-router-dom';

export default function DiscussionCard({ id, title, module, moduleClass }) {
  return (
    <Flex bgColor="blue.300" p={4} width="full" borderRadius="10">
      <Flex direction="column" alignContent="center">
        <Text as="h1" fontsize="xl" fontWeight="semibold" color="white">
          {title}
        </Text>
        {/* <Stack direction="row"> */}
        <Text as="span" fontsize="md" fontWeight="semibold" color="black">
          {module}
        </Text>
        <Text as="span" fontsize="md" fontWeight="semibold" color="black">
          {moduleClass}
        </Text>
        {/* </Stack> */}
      </Flex>
      <Spacer />
      <Link to={`/discussion/${id}`}>
        <Button mt={3} colorScheme="gray" size="sm" variant="solid">
          Lihat Detail
        </Button>
      </Link>
    </Flex>
  );
}
