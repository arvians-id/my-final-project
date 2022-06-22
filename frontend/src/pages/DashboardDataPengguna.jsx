import {
  Button,
  Table,
  TableCaption,
  TableContainer,
  Tbody,
  Td,
  Th,
  Thead,
  Tr,
} from '@chakra-ui/react';
import React from 'react';
import MainAppLayout from '../components/layout/MainAppLayout';

export default function DashboardDataPengguna() {
  return (
    <MainAppLayout>
      <TableContainer>
        <Table variant="simple">
          <Thead>
            <Tr>
              <Th>Nama</Th>
              <Th>Alamat Tempat Tinggal</Th>
              <Th>Jenis Kelamin</Th>
              <Th>Aksi</Th>
            </Tr>
          </Thead>
          <Tbody>
            <Tr>
              <Td>Budi</Td>
              <Td>Tunanetra</Td>
              <Td>Laki-laki</Td>
              <Td>
                <Button colorScheme="teal" size="sm">
                  Ubah
                </Button>
              </Td>
            </Tr>
          </Tbody>
        </Table>
      </TableContainer>
    </MainAppLayout>
  );
}
