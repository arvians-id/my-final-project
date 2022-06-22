import React from 'react';
import {
  Accordion,
  AccordionButton,
  AccordionIcon,
  AccordionItem,
  AccordionPanel,
  Box,
  Button,
  HStack,
  VStack,
} from '@chakra-ui/react';
import { AiFillBook, AiFillHome } from 'react-icons/ai';
import { MdAccountBox } from 'react-icons/md';
import { useNavigate } from 'react-router';
import useStore from '../provider/zustand/store';

export default function Sidebar() {
  const navigate = useNavigate();
  const user = useStore((state) => state.user);
  const gotoPage = (route) => {
    navigate(route);
  };

  return (
    <Box m={4} width="100%">
      <VStack spacing={4}>
        <Button
          variant="ghost"
          colorScheme="blue"
          width="full"
          justifyContent="start"
          ml={8}
          onClick={() => gotoPage('/')}
        >
          <HStack spacing={3}>
            <AiFillHome />
            <Box as="span" fontWeight="semibold">
              Home
            </Box>
          </HStack>
        </Button>
        <Accordion defaultIndex={[0]} allowMultiple width="full">
          {/* AccordionItem */}
          {user.role === 'Siswa' && (
            <AccordionItem>
              {/* Accordion Button */}
              <AccordionButton>
                <Button
                  variant="ghost"
                  colorScheme="blue"
                  width="full"
                  justifyContent="start"
                >
                  <HStack spacing={3}>
                    <AiFillBook />
                    <Box as="span" fontWeight="semibold">
                      Pelajaran
                    </Box>
                  </HStack>
                </Button>
                <AccordionIcon />
              </AccordionButton>
              {/* Acccordion Panel */}
              <AccordionPanel ml={5}>
                <Button
                  variant="ghost"
                  onClick={() => gotoPage('/course')}
                  colorScheme="blue"
                  width="full"
                  justifyContent="start"
                >
                  <HStack spacing={3}>
                    <Box as="span" fontWeight="semibold">
                      Pelajaran Anda
                    </Box>
                  </HStack>
                </Button>
                <Button
                  variant="ghost"
                  colorScheme="blue"
                  width="full"
                  justifyContent="start"
                  onClick={() => gotoPage('/submission')}
                >
                  <HStack spacing={3}>
                    <Box as="span" fontWeight="semibold">
                      Tugas
                    </Box>
                  </HStack>
                </Button>
                <Button
                  variant="ghost"
                  colorScheme="blue"
                  width="full"
                  justifyContent="start"
                  onClick={() => gotoPage('/discussion')}
                >
                  <HStack spacing={3}>
                    <Box as="span" fontWeight="semibold">
                      Diskusi
                    </Box>
                  </HStack>
                </Button>
              </AccordionPanel>
            </AccordionItem>
          )}
          {user.role === 'Guru' && (
            <AccordionItem>
              {/* Accordion Button */}
              <AccordionButton>
                <Button
                  variant="ghost"
                  colorScheme="blue"
                  width="full"
                  justifyContent="start"
                >
                  <HStack spacing={3}>
                    <AiFillBook />
                    <Box as="span" fontWeight="semibold">
                      Manajemen Siswa
                    </Box>
                  </HStack>
                </Button>
                <AccordionIcon />
              </AccordionButton>
              {/* Acccordion Panel */}
              <AccordionPanel ml={5}>
                <Button
                  variant="ghost"
                  colorScheme="blue"
                  width="full"
                  justifyContent="start"
                  onClick={() => gotoPage('/dashboard-siswa')}
                >
                  <HStack spacing={3}>
                    <Box as="span" fontWeight="semibold">
                      Data Siswa
                    </Box>
                  </HStack>
                </Button>
              </AccordionPanel>
            </AccordionItem>
          )}
          {user.role === 'Admin' && (
            <AccordionItem>
              {/* Accordion Button */}
              <AccordionButton>
                <Button
                  variant="ghost"
                  colorScheme="blue"
                  width="full"
                  justifyContent="start"
                >
                  <HStack spacing={3}>
                    <AiFillBook />
                    <Box as="span" fontWeight="semibold">
                      Manajemen Pengguna
                    </Box>
                  </HStack>
                </Button>
                <AccordionIcon />
              </AccordionButton>
              {/* Acccordion Panel */}
              <AccordionPanel ml={5}>
                <Button
                  variant="ghost"
                  colorScheme="blue"
                  width="full"
                  justifyContent="start"
                  onClick={() => gotoPage('/dashboard-pengguna')}
                >
                  <HStack spacing={3}>
                    <Box as="span" fontWeight="semibold">
                      Data Pengguna
                    </Box>
                  </HStack>
                </Button>
              </AccordionPanel>
            </AccordionItem>
          )}
          {/* Accordion Item */}
          <AccordionItem>
            {/* Accordion Button */}
            <AccordionButton>
              <Button
                variant="ghost"
                colorScheme="blue"
                width="full"
                justifyContent="start"
              >
                <HStack spacing={3}>
                  <MdAccountBox />
                  <Box as="span" fontWeight="semibold">
                    Profile
                  </Box>
                </HStack>
              </Button>
              <AccordionIcon />
            </AccordionButton>
            {/* Acccordion Panel */}
            <AccordionPanel ml={5}>
              <Button
                variant="ghost"
                colorScheme="blue"
                width="full"
                justifyContent="start"
                onClick={() => gotoPage('/profile')}
              >
                <HStack spacing={3}>
                  <Box as="span" fontWeight="semibold">
                    Profile Anda
                  </Box>
                </HStack>
              </Button>
              <Button
                variant="ghost"
                colorScheme="blue"
                width="full"
                justifyContent="start"
              >
                <HStack spacing={3}>
                  <Box as="span" fontWeight="semibold">
                    Edit Profile
                  </Box>
                </HStack>
              </Button>
              <Button
                variant="ghost"
                colorScheme="blue"
                width="full"
                justifyContent="start"
              >
                <HStack spacing={3}>
                  <Box as="span" fontWeight="semibold">
                    Log Out
                  </Box>
                </HStack>
              </Button>
            </AccordionPanel>
          </AccordionItem>
        </Accordion>
      </VStack>
    </Box>
  );
}
