import React from 'react';
import {
  Flex,
  Stack,
  Text,
  Spacer,
  Button,
  Badge,
  VStack,
  HStack,
} from '@chakra-ui/react';
import { Link } from 'react-router-dom';

export default function DiscussionCard({
  id,
  title,
  description,
  module,
  moduleClass,
  tags,
}) {
  return (
    <Flex bgColor="blue.300" p={4} width="full" borderRadius="10">
      <Flex direction="column" alignContent="center">
        <Text as="h1" fontsize="xl" fontWeight="semibold" color="white">
          Judul: {title}
        </Text>
        {/* <Stack direction="row"> */}
        <Text
          as="span"
          fontsize="md"
          maxW="450px"
          fontWeight="semibold"
          color="black"
        >
          Isi Pertayaan: {description}
        </Text>
        <HStack alignItems="flex-start">
          {tags?.split(',').map((tag, index) => (
            <Badge key={index}>#{tag}</Badge>
          ))}
        </HStack>
        <Text
          mt="4"
          as="span"
          fontsize="md"
          fontWeight="semibold"
          color="black"
        >
          Kelas: {`${module} - ${moduleClass}`}
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
