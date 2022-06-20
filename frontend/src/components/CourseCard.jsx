import { Heading, HStack, Stack, Text, Box, Button, Progress } from '@chakra-ui/react'
import React from 'react'

export default function CourseCard({ courseTitle, courseClass, courseTeacher, courseDescription }) {
    return (
        <Box p={6} shadow='md' borderWidth='1px' w={80} h={64} borderRadius={10}>
            <Stack spacing={3} height={36}>
                <Heading fontSize='xl'>{courseTitle}</Heading>
                <HStack spacing={3} mt={4}>
                    <Text fontWeight="semibold" color="grey">{courseTeacher}</Text>
                    <Text fontWeight="semibold" color="blue.200">{courseClass}</Text>
                </HStack>
                <Progress value={70} />
                <Text mt={4} align="justify">{`${courseDescription.substring(0,60)} ...`}</Text>
            </Stack>
            <Button mt={5} colorScheme="blue" width="full" bottom={2}>
                Lanjutkan
            </Button>
        </Box>
    )
}
