import React from 'react'
import { Box, HStack, Image, Spacer, Avatar, InputGroup, InputLeftElement, Input, MenuButton, Menu, Button, MenuItem, MenuList } from '@chakra-ui/react'
import { Link } from 'react-router-dom'
import { ChevronDownIcon, SearchIcon } from '@chakra-ui/icons'
import Logo from '../image/Teenager.png'
export default function Navbar() {
    return (
        <div>
            <Box bg="white" minWidth="100%" height="10vh" boxShadow='lg' px='5' rounded='md' display="flex" justifyContent="flex-start" alignItems="center" position="sticky" top="0">
                {/* Box For Logo Image using Shadow On Bottom*/}
                <Box >
                    <Image src={Logo} alt="logo" width="35px" height="40px" />
                </Box>
                <Spacer />
                <Box>
                    <HStack spacing={4}>
                        {/* Search Bar */}
                        <Box>
                            <InputGroup>
                                <InputLeftElement
                                    pointerEvents='none'
                                    children={<SearchIcon color='gray.300' />}
                                />
                                <Input type='tel' placeholder='Cari' />
                            </InputGroup>
                        </Box>
                        <Box>
                            <Menu>
                                <MenuButton as={Button} p={3} variant='ghost' rightIcon={<ChevronDownIcon />}>
                                    <HStack>
                                        <Avatar name='Irfan Kurniawan' src='https://bit.ly/dan-abramov' mr={2} w={8} h={8} />
                                        <Stack>
                                            <Text as='span' fontSize="md" fontWeight='semibold'>
                                                Irfan Kurniawan
                                            </Text>
                                            <Text as='span' fontSize="sm" align="left" fontWeight='semibold' color="grey">
                                                Siswa
                                            </Text>
                                        </Stack>
                                    </HStack>
                                </MenuButton>
                                <MenuList>
                                    <MenuItem>Profile</MenuItem>
                                    <Link to="/login">
                                        <MenuItem>Log Out</MenuItem>
                                    </Link>
                                </MenuList>
                            </Menu>
                        </Box>
                    </HStack>
                </Box>
            </Box>
        </div>
    )
}
