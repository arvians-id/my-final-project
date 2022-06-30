import React from 'react';
import { Flex, Text, Spacer, Box } from '@chakra-ui/react';
import { Link } from 'react-router-dom';
export default function ListModule({ article, courseCode }) {
  return (
    <Link to={`/course/${courseCode}/article/${article.id}`}>
      <Box
        bgColor="blue.500"
        p={4}
        width="full"
        borderRadius="10"
        _hover={{ cursor: 'pointer' }}
      >
        <Text as="span" fontSize="md" fontWeight="semibold" color="white">
          Judul: {article.name}
        </Text>
        <Spacer />
        <Text as="span" fontSize="md" fontWeight="semibold" color="white">
          Deskripsi: {article.content.slice(0, 20)}...
        </Text>
      </Box>
    </Link>
  );
}
