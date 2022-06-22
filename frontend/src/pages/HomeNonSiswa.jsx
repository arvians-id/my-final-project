import { Box, Stack } from '@chakra-ui/layout';
import React from 'react';
import MainAppLayout from '../components/layout/MainAppLayout';

export default function HomeNonSiswa() {
  return (
    <MainAppLayout>
      <Box m={5}>
        <Stack spacing={6}>
          {/* Header */}
          <Box>
            <Box as="h1" fontSize="2xl" fontWeight="semibold">
              Selamat Datang
            </Box>
          </Box>
        </Stack>
      </Box>
    </MainAppLayout>
  );
}
