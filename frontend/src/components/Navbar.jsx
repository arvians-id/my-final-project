import React from 'react';
import {
  Box,
  HStack,
  Image,
  Spacer,
  Avatar,
  InputGroup,
  InputLeftElement,
  Input,
  MenuButton,
  Menu,
  Button,
  MenuItem,
  MenuList,
  Stack,
  Text,
} from '@chakra-ui/react';
import { Link, useNavigate } from 'react-router-dom';
import { ChevronDownIcon, SearchIcon } from '@chakra-ui/icons';
import Logo from '../image/Teenager.png';
import useStore from '../provider/zustand/store';
import { localClearToken } from '../utils/token';

export default function Navbar() {
  const user = useStore((state) => state.user);
  const setUser = useStore((state) => state.setUser);
  const navigate = useNavigate();

  const handleLogout = () => {
    setUser(undefined);
    localClearToken();
    navigate('/');
  };

  return (
    <div>
      <Box
        bg="white"
        minWidth="100%"
        height="10vh"
        boxShadow="md"
        px="5"
        rounded="md"
        display="flex"
        justifyContent="flex-start"
        alignItems="center"
        position="fixed"
        top="0"
        zIndex="10"
      >
        {/* Box For Logo Image using Shadow On Bottom*/}
        <Box>
          <Image src={Logo} alt="logo" width="35px" height="40px" />
        </Box>
        <Spacer />
        <Box>
          <HStack spacing={4}>
            {/* Search Bar */}
            <Box>
              <InputGroup>
                <InputLeftElement
                  pointerEvents="none"
                  children={<SearchIcon color="gray.300" />}
                />
                <Input type="tel" placeholder="Cari" />
              </InputGroup>
            </Box>
            <Box>
              <Menu>
                <MenuButton
                  as={Button}
                  p={3}
                  variant="ghost"
                  rightIcon={<ChevronDownIcon />}
                >
                  <HStack>
                    <Avatar
                      name={user?.username}
                      src={user?.profile_image ?? 'https://bit.ly/dan-abramov'}
                      mr={2}
                      w={8}
                      h={8}
                    />
                    <Stack>
                      <Text as="span" fontSize="md" fontWeight="semibold">
                        {user?.username}{' '}
                      </Text>
                      <Text
                        as="span"
                        fontSize="sm"
                        align="left"
                        fontWeight="semibold"
                        color="grey"
                      >
                        {user?.role}
                      </Text>
                    </Stack>
                  </HStack>
                </MenuButton>
                <MenuList>
                  <Link to="/profile">
                    <MenuItem>Profile</MenuItem>
                  </Link>
                  <MenuItem onClick={handleLogout}>Log Out</MenuItem>
                </MenuList>
              </Menu>
            </Box>
          </HStack>
        </Box>
      </Box>
    </div>
  );
}
