import React from 'react';
import { Flex, Stack, Text, Spacer, Button } from '@chakra-ui/react';

export default function DiscussionCard({ title, module, moduleClass }) {
  return (
    <Flex bgColor="blue.300" p={4} width="full" height={24} borderRadius="10">
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
      <Button mt={3} colorScheme="gray" size="sm" variant="solid">
        Lihat Semua
      </Button>
    </Flex>
  );
}
