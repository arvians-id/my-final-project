import React from 'react';
import { Box, Flex, Stack, Text, Spacer } from '@chakra-ui/react';

export default function DiscussionCard({
  question,
  module,
  moduleClass,
  answerNum,
  answer,
}) {
  return (
    <Box width="full">
      <Flex bgColor="blue.300" p={4} width="full" height={24} borderRadius="10">
        <Flex direction="column" alignContent="center">
          <Text as="h1" fontsize="4xl" fontWeight="semibold" color="white">
            {question}
          </Text>
          <Stack direction="row">
            <Text as="span" fontsize="md" fontWeight="semibold" color="black">
              {module}
            </Text>
            <Text as="span" fontsize="md" fontWeight="semibold" color="black">
              {moduleClass}
            </Text>
          </Stack>
        </Flex>
        <Spacer />
      </Flex>
    </Box>
  );
}
