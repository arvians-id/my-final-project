import {
  Heading,
  HStack,
  Stack,
  Text,
  Box,
  Button,
  Progress,
  Spinner,
} from '@chakra-ui/react';
import React, { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';
import { API_GET_COURSE_BY_CODE } from '../api/course';

export default function CourseCard({ course_code }) {
  const [loadingGetDetailCourse, setLoadingGetDetailCourse] = useState(false);
  const [detailCourse, setDetailCourse] = useState();

  const getCourseDetail = async (code) => {
    setLoadingGetDetailCourse(true);
    const res = await API_GET_COURSE_BY_CODE(code);
    if (res.status === 200) {
      setDetailCourse(res.data.data);
    }
    setLoadingGetDetailCourse(false);
  };
  useEffect(() => {
    getCourseDetail(course_code);
  }, []);

  return (
    <Box p={6} shadow="md" borderWidth="1px" w={80} h={64} borderRadius={10}>
      {loadingGetDetailCourse ? (
        <Spinner />
      ) : detailCourse ? (
        <>
          <Stack spacing={3} height={36}>
            <Heading fontSize="xl">{detailCourse.name}</Heading>
            <HStack spacing={3} mt={4}>
              {/* <Text fontWeight="semibold" color="grey">{courseTeacher}</Text> */}
              <Text fontWeight="semibold" color="blue.200">
                {detailCourse.class}
              </Text>
            </HStack>
            <Progress value={80} />
            <Text mt={4} align="justify">{`${detailCourse.description.substring(
              0,
              60
            )} ...`}</Text>
          </Stack>
          <Link to={`/course/${course_code}`}>
            <Button mt={5} colorScheme="blue" width="full" bottom={2}>
              Lanjutkan
            </Button>
          </Link>
        </>
      ) : (
        <Text>Course tidak ditemukan</Text>
      )}
    </Box>
  );
}
