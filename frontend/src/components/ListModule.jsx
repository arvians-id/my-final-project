import React from 'react';
import { Flex, Text, Spacer } from '@chakra-ui/react';
export default function ListModule({ name, list }) {
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
      <Text as="span" fontsize="md" fontWeight="semibold" color="white">
        {list}
      </Text>
    </Flex>
  );
}
