import React from 'react';
import {
    Button,
    InputGroup,
    InputLeftElement,
    Input,
    FormLabel,
    MenuButton,
    MenuList,
    Menu,
    MenuItem,
    Spacer,
    VStack,
    Box,
    Flex,
    HStack,
    SimpleGrid,
} from '@chakra-ui/react';
import { SearchIcon, ChevronDownIcon } from '@chakra-ui/icons';
import Navbar from '../components/Navbar';
import Sidebar from '../components/Sidebar';
export default function Course() {
    return (
        <div>
            <Navbar />

            <Flex direction="row" justifyContent="center" alignItems="center">
                <Flex
                    width="20%"
                    minHeight="200vh"
                    bgColor="grey.100"
                    boxShadow="dark-lg"
                >
                    <Sidebar />
                </Flex>
                <Box width="80%" minHeight="200vh" bg="white">
                    <Box m={5}>
                        <Box as="h1" fontSize="xl" fontWeight="semibold" mb={2}>
                            Course
                        </Box>
                    </Box>
                    <Spacer />
                    <Box mt="80px" ml="1150px">
                        <Menu>
                            <MenuButton
                                as={Button}
                                rightIcon={<ChevronDownIcon />}
                            >
                                Kelas
                            </MenuButton>
                            <MenuList>
                                <MenuItem>1 SMA</MenuItem>
                                <MenuItem>2 SMA</MenuItem>
                                <MenuItem>3 SMA</MenuItem>
                            </MenuList>
                        </Menu>
                    </Box>
                    <Box>
                        <SimpleGrid
                            bg="white"
                            columns={{ sm: 4, md: 4 }}
                            spacing="8"
                            p="10"
                            textAlign="center"
                            rounded="lg"
                            color="gray.400"
                        >
                            <Box
                                boxSize="250px"
                                color="black"
                                as="button"
                                boxShadow="md"
                                p="6"
                                rounded="md"
                                bg="gray.100"
                            >
                                Bahasa Indonesia Kelas 3 SMA
                            </Box>
                            <Box
                                boxSize="250px"
                                color="black"
                                as="button"
                                boxShadow="md"
                                p="6"
                                rounded="md"
                                bg="gray.100"
                            >
                                Bahasa Indonesia Kelas 3 SMA
                            </Box>
                            <Box
                                boxSize="250px"
                                color="black"
                                as="button"
                                boxShadow="md"
                                p="6"
                                rounded="lg"
                                bg="gray.100"
                            >
                                Bahasa Indonesia Kelas 3 SMA
                            </Box>
                            <Box
                                boxSize="250px"
                                color="black"
                                as="button"
                                boxShadow="md"
                                p="6"
                                rounded="lg"
                                bg="gray.100"
                            >
                                Bahasa Indonesia Kelas 3 SMA
                            </Box>
                            <Box
                                boxSize="250px"
                                color="black"
                                as="button"
                                boxShadow="md"
                                p="6"
                                rounded="lg"
                                bg="gray.100"
                            >
                                Bahasa Indonesia Kelas 3 SMA
                            </Box>
                            <Box
                                boxSize="250px"
                                color="black"
                                as="button"
                                boxShadow="md"
                                p="6"
                                rounded="lg"
                                bg="gray.100"
                            >
                                Bahasa Indonesia Kelas 3 SMA
                            </Box>
                            <Box
                                boxSize="250px"
                                color="black"
                                as="button"
                                boxShadow="md"
                                p="6"
                                rounded="lg"
                                bg="gray.100"
                            >
                                Bahasa Indonesia Kelas 3 SMA
                            </Box>
                            <Box
                                boxSize="250px"
                                color="black"
                                as="button"
                                boxShadow="md"
                                p="6"
                                rounded="lg"
                                bg="gray.100"
                            >
                                Bahasa Indonesia Kelas 3 SMA
                            </Box>
                            <Box
                                boxSize="250px"
                                color="black"
                                as="button"
                                boxShadow="md"
                                p="6"
                                rounded="lg"
                                bg="gray.100"
                            >
                                Bahasa Indonesia Kelas 3 SMA
                            </Box>
                            <Box
                                boxSize="250px"
                                color="black"
                                as="button"
                                boxShadow="md"
                                p="6"
                                rounded="lg"
                                bg="gray.100"
                            >
                                Bahasa Indonesia Kelas 3 SMA
                            </Box>
                            <Box
                                boxSize="250px"
                                color="black"
                                as="button"
                                boxShadow="md"
                                p="6"
                                rounded="lg"
                                bg="gray.100"
                            >
                                Bahasa Indonesia Kelas 3 SMA
                            </Box>
                            <Box
                                boxSize="250px"
                                color="black"
                                as="button"
                                boxShadow="md"
                                p="6"
                                rounded="lg"
                                bg="gray.100"
                            >
                                Bahasa Indonesia Kelas 3 SMA
                            </Box>
                        </SimpleGrid>
                    </Box>
                </Box>
            </Flex>
        </div>
    );
}
