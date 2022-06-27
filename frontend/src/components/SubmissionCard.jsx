import React, { useState } from 'react';
import { Flex, Text, Spacer, Badge, Box, Button } from '@chakra-ui/react';
import SubmitCard from './SubmitCard';

export default function SubmissionCard({
  name,
  status,
  type,
  courseCode,
  submissionId,
  getListSubmission,
}) {
  const [openFormSubmit, setOpenFormSubmit] = useState(false);
  console.log(courseCode, submissionId);
  const onToggleButtonSubmit = () => {
    setOpenFormSubmit(!openFormSubmit);
  };

  const onSuccess = () => {
    setOpenFormSubmit(false);
    getListSubmission();
  };

  return (
    <Box width="full">
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
        <Flex alignItems="center" gap="15px">
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
          {type === 'submit' && status === 'Belum Terkumpul' && (
            <Button colorScheme="green" h="30px" onClick={onToggleButtonSubmit}>
              Kumpul Tugas
            </Button>
          )}
        </Flex>
      </Flex>
      {openFormSubmit && (
        <Box mt="4">
          <SubmitCard
            onSuccess={onSuccess}
            courseCode={courseCode}
            submissionId={submissionId}
          />
        </Box>
      )}
    </Box>
  );
}
